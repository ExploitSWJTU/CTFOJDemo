package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"

	"github.com/gin-gonic/gin"
)

// ChallengeAPI handles challenge-related requests
type ChallengeAPI struct {
	db *repository.Database
}

// NewChallengeAPI creates a new ChallengeAPI instance
func NewChallengeAPI(db *repository.Database) *ChallengeAPI {
	return &ChallengeAPI{db: db}
}

// GetChallenges returns a list of challenges with optional filters
// @Summary List challenges
// @Tags Challenges
// @Produce json
// @Param category query string false "Filter by category"
// @Param difficulty query string false "Filter by difficulty"
// @Param search query string false "Search in title and description"
// @Success 200 {object} model.APIResponse
// @Router /api/challenges [get]
func (a *ChallengeAPI) GetChallenges(c *gin.Context) {
	category := c.Query("category")
	difficulty := c.Query("difficulty")
	search := c.Query("search")

	query := a.db.DB.Model(&model.Challenge{}).Where("deleted_at IS NULL")

	// Apply filters
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}
	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	var total int64
	query.Count(&total)

	// Get challenges
	var challenges []model.Challenge
	if err := query.Order("category ASC, points ASC").Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch challenges: "+err.Error()))
		return
	}

	// Transform challenges (exclude sensitive fields)
	type SafeChallenge struct {
		ID              uint                  `json:"id"`
		Title           string                `json:"title"`
		Description     string                `json:"description"`
		Category        string                `json:"category"`
		Points          int                   `json:"points"`
		Difficulty      string                `json:"difficulty"`
		Type            string                `json:"type"`
		SolvedCount     int                   `json:"solved_count"`
		ContainerConfig model.ContainerConfig `json:"container_config,omitempty"`
		Attachments     []model.Attachment    `json:"attachments,omitempty"`
	}

	safeChallenges := make([]SafeChallenge, len(challenges))
	for i, ch := range challenges {
		var containerConfig model.ContainerConfig
		if ch.ContainerConfig != nil {
			if err := json.Unmarshal(ch.ContainerConfig, &containerConfig); err == nil {
				safeChallenges[i].ContainerConfig = containerConfig
			}
		}

		var attachments []model.Attachment
		if ch.Attachments != nil {
			if err := json.Unmarshal(ch.Attachments, &attachments); err == nil {
				safeChallenges[i].Attachments = attachments
			}
		}

		safeChallenges[i].ID = ch.ID
		safeChallenges[i].Title = ch.Title
		safeChallenges[i].Description = ch.Description
		safeChallenges[i].Category = ch.Category
		safeChallenges[i].Points = ch.Points
		safeChallenges[i].Difficulty = ch.Difficulty
		safeChallenges[i].Type = ch.Type
		safeChallenges[i].SolvedCount = ch.SolvedCount
	}

	c.JSON(http.StatusOK, model.PaginatedResponse(safeChallenges, model.Pagination{
		Page:     1,
		PageSize: len(safeChallenges),
		Total:    total,
	}))
}

// GetChallenge returns a single challenge by ID
// @Summary Get challenge detail
// @Tags Challenges
// @Produce json
// @Param id path int true "Challenge ID"
// @Success 200 {object} model.APIResponse
// @Router /api/challenges/:id [get]
func (a *ChallengeAPI) GetChallenge(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid challenge ID"))
		return
	}

	var challenge model.Challenge
	if err := a.db.DB.First(&challenge, id).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Challenge not found"))
		return
	}

	// Transform challenge (exclude sensitive fields like Flag)
	var containerConfig model.ContainerConfig
	if challenge.ContainerConfig != nil {
		json.Unmarshal(challenge.ContainerConfig, &containerConfig)
	}

	var attachments []model.Attachment
	if challenge.Attachments != nil {
		json.Unmarshal(challenge.Attachments, &attachments)
	}

	safeChallenge := gin.H{
		"id":               challenge.ID,
		"title":            challenge.Title,
		"description":      challenge.Description,
		"category":         challenge.Category,
		"points":           challenge.Points,
		"difficulty":       challenge.Difficulty,
		"type":             challenge.Type,
		"solved_count":     challenge.SolvedCount,
		"container_config": containerConfig,
		"attachments":      attachments,
		"created_at":       challenge.CreatedAt,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(safeChallenge))
}

// GetChallengesByCategory returns challenges grouped by category
// @Summary Get challenges by category
// @Tags Challenges
// @Produce json
// @Success 200 {object} model.APIResponse
// @Router /api/challenges/categories [get]
func (a *ChallengeAPI) GetChallengesByCategory(c *gin.Context) {
	type CategoryStats struct {
		Category   string `json:"category"`
		Total      int    `json:"total"`
		Solved     int    `json:"solved"`
		UserSolved int    `json:"user_solved"`
	}

	// Get all categories with counts
	var stats []CategoryStats
	query := `
		SELECT 
			category,
			COUNT(*) as total,
			SUM(solved_count) as solved
		FROM challenges 
		WHERE deleted_at IS NULL 
		GROUP BY category 
		ORDER BY category
	`

	if err := a.db.DB.Raw(query).Scan(&stats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch category stats: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(stats))
}

// SearchChallenges searches challenges by keyword
// @Summary Search challenges
// @Tags Challenges
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {object} model.APIResponse
// @Router /api/challenges/search [get]
func (a *ChallengeAPI) SearchChallenges(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Search query is required"))
		return
	}

	var challenges []model.Challenge
	searchPattern := "%" + strings.TrimSpace(query) + "%"
	if err := a.db.DB.Where("title ILIKE ? OR description ILIKE ?", searchPattern, searchPattern).
		Limit(10).Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Search failed: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(challenges))
}
