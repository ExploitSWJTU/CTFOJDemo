package api

import (
	"net/http"
	"time"

	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"

	"github.com/gin-gonic/gin"
)

type SubmissionAPI struct {
	db *repository.Database
}

type submissionRecordRow struct {
	ID             uint
	UserID         uint
	Username       string
	ChallengeID    uint
	ChallengeTitle string
	Category       string
	Difficulty     string
	Points         int
	SubmittedFlag  string
	IsCorrect      bool
	IsSolve        bool
	CreatedAt      time.Time
}

func NewSubmissionAPI(db *repository.Database) *SubmissionAPI {
	return &SubmissionAPI{db: db}
}

func (a *SubmissionAPI) GetMySubmissions(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	items, err := a.loadSubmissionRecords(&userID, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch submission history: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(items))
}

func (a *SubmissionAPI) GetMySolveRecords(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	items, err := a.loadSubmissionRecords(&userID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch solve records: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(items))
}

func (a *SubmissionAPI) AdminListSubmissions(c *gin.Context) {
	items, err := a.loadSubmissionRecords(nil, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch submission records: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(items))
}

func (a *SubmissionAPI) loadSubmissionRecords(userID *uint, solveOnly bool) ([]gin.H, error) {
	var rows []submissionRecordRow

	query := a.db.DB.Table("submissions").
		Select(`
			submissions.id,
			submissions.user_id,
			COALESCE(users.username, '') AS username,
			submissions.challenge_id,
			COALESCE(challenges.title, '') AS challenge_title,
			COALESCE(challenges.category, '') AS category,
			COALESCE(challenges.difficulty, '') AS difficulty,
			COALESCE(challenges.points, 0) AS points,
			submissions.submitted_flag,
			submissions.is_correct,
			CASE
				WHEN submissions.is_solve = true THEN true
				WHEN submissions.is_correct = true AND first_solve_events.first_submission_id = submissions.id THEN true
				ELSE false
			END AS is_solve,
			submissions.created_at
		`).
		Joins("LEFT JOIN users ON users.id = submissions.user_id").
		Joins("LEFT JOIN challenges ON challenges.id = submissions.challenge_id").
		Joins(`
			LEFT JOIN (
				SELECT MIN(id) AS first_submission_id, user_id, challenge_id
				FROM submissions
				WHERE deleted_at IS NULL AND is_correct = true
				GROUP BY user_id, challenge_id
			) AS first_solve_events
			ON first_solve_events.first_submission_id = submissions.id
		`).
		Where("submissions.deleted_at IS NULL").
		Order("submissions.created_at DESC")

	if userID != nil {
		query = query.Where("submissions.user_id = ?", *userID)
	}

	if solveOnly {
		query = query.Where(`
			submissions.is_solve = true
			OR (submissions.is_correct = true AND first_solve_events.first_submission_id = submissions.id)
		`)
	}

	if err := query.Find(&rows).Error; err != nil {
		return nil, err
	}

	items := make([]gin.H, 0, len(rows))
	for i := range rows {
		items = append(items, gin.H{
			"id":              rows[i].ID,
			"user_id":         rows[i].UserID,
			"username":        rows[i].Username,
			"challenge_id":    rows[i].ChallengeID,
			"challenge_title": rows[i].ChallengeTitle,
			"category":        rows[i].Category,
			"difficulty":      rows[i].Difficulty,
			"points":          rows[i].Points,
			"submitted_flag":  rows[i].SubmittedFlag,
			"is_correct":      rows[i].IsCorrect,
			"is_solve":        rows[i].IsSolve,
			"created_at":      rows[i].CreatedAt,
		})
	}

	return items, nil
}
