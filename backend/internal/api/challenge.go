package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"
	"swjtu-ctf-oj/internal/service"

	"github.com/gin-gonic/gin"
)

type ChallengeAPI struct {
	db      *repository.Database
	service *service.ChallengeService
}

type FlagSubmitRequest struct {
	Flag string `json:"flag" binding:"required"`
}

type AdminCreateChallengeRequest struct {
	Title                    string `json:"title" binding:"required"`
	Description              string `json:"description" binding:"required"`
	Category                 string `json:"category" binding:"required"`
	Points                   int    `json:"points"`
	Difficulty               string `json:"difficulty" binding:"required"`
	Type                     string `json:"type"`
	Image                    string `json:"image" binding:"required"`
	InternalPort             int    `json:"internal_port" binding:"required"`
	ExpectedFlag             string `json:"expected_flag" binding:"required"`
	ContainerDurationSeconds int    `json:"container_duration_seconds"`
	IsVisible                *bool  `json:"is_visible"`
}

type AdminUpdateChallengeRequest struct {
	Title                    string `json:"title" binding:"required"`
	Description              string `json:"description" binding:"required"`
	Category                 string `json:"category" binding:"required"`
	Points                   int    `json:"points"`
	Difficulty               string `json:"difficulty" binding:"required"`
	Type                     string `json:"type"`
	Image                    string `json:"image" binding:"required"`
	InternalPort             int    `json:"internal_port" binding:"required"`
	ExpectedFlag             string `json:"expected_flag" binding:"required"`
	ContainerDurationSeconds int    `json:"container_duration_seconds"`
	IsVisible                *bool  `json:"is_visible"`
}

type AdminVisibilityRequest struct {
	IsVisible bool `json:"is_visible"`
}

func NewChallengeAPI(db *repository.Database, svc *service.ChallengeService) *ChallengeAPI {
	return &ChallengeAPI{db: db, service: svc}
}

func (a *ChallengeAPI) GetChallenges(c *gin.Context) {
	category := c.Query("category")
	difficulty := c.Query("difficulty")
	search := c.Query("search")

	query := a.db.DB.Model(&model.Challenge{}).Where("deleted_at IS NULL AND is_visible = ?", true)
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if difficulty != "" {
		query = query.Where("difficulty = ?", difficulty)
	}
	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	query.Count(&total)

	var challenges []model.Challenge
	if err := query.Order("category ASC, points ASC").Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch challenges: "+err.Error()))
		return
	}

	items := make([]gin.H, 0, len(challenges))
	for i := range challenges {
		items = append(items, buildSafeChallenge(&challenges[i], false))
	}

	c.JSON(http.StatusOK, model.PaginatedResponse(items, model.Pagination{
		Page:     1,
		PageSize: len(items),
		Total:    total,
	}))
}

func (a *ChallengeAPI) GetChallenge(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var challenge model.Challenge
	if err := a.db.DB.Where("is_visible = ?", true).First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Challenge not found"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(buildSafeChallenge(&challenge, false)))
}

func (a *ChallengeAPI) GetChallengesByCategory(c *gin.Context) {
	type CategoryStats struct {
		Category string `json:"category"`
		Total    int    `json:"total"`
		Solved   int    `json:"solved"`
	}

	var stats []CategoryStats
	query := `
		SELECT category, COUNT(*) as total, COALESCE(SUM(solved_count), 0) as solved
		FROM challenges
		WHERE deleted_at IS NULL AND is_visible = true
		GROUP BY category
		ORDER BY category
	`
	if err := a.db.DB.Raw(query).Scan(&stats).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch category stats: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(stats))
}

func (a *ChallengeAPI) SearchChallenges(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))
	if query == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Search query is required"))
		return
	}

	var challenges []model.Challenge
	if err := a.db.DB.
		Where("is_visible = ?", true).
		Where("title ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").
		Limit(10).
		Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Search failed: "+err.Error()))
		return
	}

	items := make([]gin.H, 0, len(challenges))
	for i := range challenges {
		items = append(items, buildSafeChallenge(&challenges[i], false))
	}
	c.JSON(http.StatusOK, model.SuccessResponse(items))
}

func (a *ChallengeAPI) StartContainer(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		return
	}

	record, err := a.service.StartContainer(userID, challengeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Failed to start challenge container: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(a.service.BuildContainerResponse(record)))
}

func (a *ChallengeAPI) GetContainer(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		return
	}

	record, err := a.service.GetContainer(userID, challengeID)
	if err != nil {
		if errors.Is(err, service.ErrActiveContainerNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Active container not found"))
			return
		}
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Failed to fetch container: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(a.service.BuildContainerResponse(record)))
}

func (a *ChallengeAPI) ExtendContainer(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		return
	}

	record, err := a.service.ExtendContainer(userID, challengeID)
	if err != nil {
		if errors.Is(err, service.ErrActiveContainerNotFound) {
			c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Active container not found"))
			return
		}
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Failed to extend container: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(a.service.BuildContainerResponse(record)))
}

func (a *ChallengeAPI) DestroyContainer(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		return
	}

	if err := a.service.DestroyContainer(userID, challengeID); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Failed to destroy container: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{"status": "stopped"}))
}

func (a *ChallengeAPI) SubmitFlag(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var req FlagSubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid request: "+err.Error()))
		return
	}

	result, err := a.service.SubmitFlag(userID, challengeID, req.Flag)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Failed to submit flag: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(result))
}

func (a *ChallengeAPI) AdminListChallenges(c *gin.Context) {
	var challenges []model.Challenge
	if err := a.db.DB.Order("created_at DESC").Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch admin challenges: "+err.Error()))
		return
	}

	items := make([]gin.H, 0, len(challenges))
	for i := range challenges {
		items = append(items, buildSafeChallenge(&challenges[i], true))
	}
	c.JSON(http.StatusOK, model.SuccessResponse(items))
}

func (a *ChallengeAPI) AdminListDeletedChallenges(c *gin.Context) {
	var challenges []model.Challenge
	if err := a.db.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		Order("deleted_at DESC").
		Find(&challenges).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to fetch deleted challenges: "+err.Error()))
		return
	}

	items := make([]gin.H, 0, len(challenges))
	for i := range challenges {
		items = append(items, buildSafeChallenge(&challenges[i], true))
	}
	c.JSON(http.StatusOK, model.SuccessResponse(items))
}

func (a *ChallengeAPI) AdminCreateChallenge(c *gin.Context) {
	var req AdminCreateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid request: "+err.Error()))
		return
	}

	if req.Points <= 0 {
		req.Points = 100
	}
	if req.ContainerDurationSeconds <= 0 {
		req.ContainerDurationSeconds = 3600
	}

	isVisible := true
	if req.IsVisible != nil {
		isVisible = *req.IsVisible
	}

	challengeType := req.Type
	if challengeType == "" {
		challengeType = "web_static_flag"
	}

	containerConfig, _ := json.Marshal(model.ContainerConfig{
		Image:       req.Image,
		ExposedPort: req.InternalPort,
	})

	challenge := model.Challenge{
		Title:                    req.Title,
		Description:              req.Description,
		Category:                 req.Category,
		Points:                   req.Points,
		Difficulty:               req.Difficulty,
		Type:                     challengeType,
		Image:                    req.Image,
		InternalPort:             req.InternalPort,
		ExpectedFlag:             req.ExpectedFlag,
		Flag:                     req.ExpectedFlag,
		ContainerDurationSeconds: req.ContainerDurationSeconds,
		IsVisible:                isVisible,
		ContainerConfig:          containerConfig,
	}

	if err := a.db.DB.Create(&challenge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to create challenge: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(buildSafeChallenge(&challenge, true)))
}

func (a *ChallengeAPI) AdminUpdateChallenge(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req AdminUpdateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid request: "+err.Error()))
		return
	}

	var challenge model.Challenge
	if err := a.db.DB.First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Challenge not found"))
		return
	}

	if req.Points <= 0 {
		req.Points = 100
	}
	if req.ContainerDurationSeconds <= 0 {
		req.ContainerDurationSeconds = 3600
	}

	challengeType := req.Type
	if challengeType == "" {
		challengeType = challenge.Type
	}
	if challengeType == "" {
		challengeType = "web_static_flag"
	}

	containerConfig, _ := json.Marshal(model.ContainerConfig{
		Image:       req.Image,
		ExposedPort: req.InternalPort,
	})

	isVisible := challenge.IsVisible
	if req.IsVisible != nil {
		isVisible = *req.IsVisible
	}

	updates := map[string]any{
		"title":                      req.Title,
		"description":                req.Description,
		"category":                   req.Category,
		"points":                     req.Points,
		"difficulty":                 req.Difficulty,
		"type":                       challengeType,
		"image":                      req.Image,
		"internal_port":              req.InternalPort,
		"expected_flag":              req.ExpectedFlag,
		"flag":                       req.ExpectedFlag,
		"container_duration_seconds": req.ContainerDurationSeconds,
		"is_visible":                 isVisible,
		"container_config":           containerConfig,
	}

	if err := a.db.DB.Model(&challenge).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to update challenge: "+err.Error()))
		return
	}

	if err := a.db.DB.First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to reload challenge"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(buildSafeChallenge(&challenge, true)))
}

func (a *ChallengeAPI) AdminSetChallengeVisibility(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var req AdminVisibilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid request: "+err.Error()))
		return
	}

	var challenge model.Challenge
	if err := a.db.DB.First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Challenge not found"))
		return
	}

	if err := a.db.DB.Model(&challenge).Update("is_visible", req.IsVisible).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to update challenge visibility: "+err.Error()))
		return
	}

	if err := a.db.DB.First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to reload challenge"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(buildSafeChallenge(&challenge, true)))
}

func (a *ChallengeAPI) AdminDeleteChallenge(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var challenge model.Challenge
	if err := a.db.DB.First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Challenge not found"))
		return
	}

	destroyedContainers, err := a.service.DestroyChallengeContainers(challengeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to stop active challenge containers: "+err.Error()))
		return
	}

	if err := a.db.DB.Delete(&challenge).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to delete challenge: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(service.AdminDeleteResult{
		Deleted:             true,
		DestroyedContainers: destroyedContainers,
	}))
}

func (a *ChallengeAPI) AdminRestoreChallenge(c *gin.Context) {
	challengeID, ok := parseUintParam(c, "id")
	if !ok {
		return
	}

	var challenge model.Challenge
	if err := a.db.DB.Unscoped().
		Where("deleted_at IS NOT NULL").
		First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "Deleted challenge not found"))
		return
	}

	if err := a.db.DB.Unscoped().Model(&challenge).Update("deleted_at", nil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to restore challenge: "+err.Error()))
		return
	}

	if err := a.db.DB.First(&challenge, challengeID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to reload restored challenge"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(buildSafeChallenge(&challenge, true)))
}

func buildSafeChallenge(challenge *model.Challenge, includeSecret bool) gin.H {
	var containerConfig model.ContainerConfig
	if len(challenge.ContainerConfig) > 0 {
		_ = json.Unmarshal(challenge.ContainerConfig, &containerConfig)
	}

	var attachments []model.Attachment
	if len(challenge.Attachments) > 0 {
		_ = json.Unmarshal(challenge.Attachments, &attachments)
	}

	payload := gin.H{
		"id":                         challenge.ID,
		"title":                      challenge.Title,
		"description":                challenge.Description,
		"category":                   challenge.Category,
		"points":                     challenge.Points,
		"difficulty":                 challenge.Difficulty,
		"type":                       challenge.Type,
		"image":                      challenge.Image,
		"internal_port":              challenge.InternalPort,
		"container_duration_seconds": challenge.ContainerDurationSeconds,
		"is_visible":                 challenge.IsVisible,
		"solved_count":               challenge.SolvedCount,
		"container_config":           containerConfig,
		"attachments":                attachments,
		"created_at":                 challenge.CreatedAt,
		"updated_at":                 challenge.UpdatedAt,
	}

	if challenge.DeletedAt.Valid {
		payload["deleted_at"] = challenge.DeletedAt.Time
	}

	if includeSecret {
		expectedFlag := challenge.ExpectedFlag
		if expectedFlag == "" {
			expectedFlag = challenge.Flag
		}
		payload["expected_flag"] = expectedFlag
	}

	return payload
}

func parseUintParam(c *gin.Context, key string) (uint, bool) {
	value := c.Param(key)
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid "+key))
		return 0, false
	}
	return uint(parsed), true
}

func getUserID(c *gin.Context) (uint, bool) {
	value, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "User not authenticated"))
		return 0, false
	}

	userID, ok := value.(uint)
	if ok {
		return userID, true
	}

	c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Invalid user context"))
	return 0, false
}
