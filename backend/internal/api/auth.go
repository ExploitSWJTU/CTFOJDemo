package api

import (
	"net/http"
	"regexp"

	"swjtu-ctf-oj/internal/middleware"
	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"

	"github.com/gin-gonic/gin"
)

// AuthAPI handles authentication requests
type AuthAPI struct {
	db *repository.Database
}

// NewAuthAPI creates a new AuthAPI instance
func NewAuthAPI(db *repository.Database) *AuthAPI {
	return &AuthAPI{db: db}
}

// LoginRequest represents login request body
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents register request body
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// Login handles user login
// @Summary User login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} model.APIResponse
// @Router /api/auth/login [post]
func (a *AuthAPI) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid request: "+err.Error()))
		return
	}

	// Find user
	var user model.User
	result := a.db.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Invalid username or password"))
		return
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(401, "Invalid username or password"))
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to generate token"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"score":      user.Score,
			"created_at": user.CreatedAt,
		},
	}))
}

// Register handles user registration
// @Summary Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Registration details"
// @Success 200 {object} model.APIResponse
// @Router /api/auth/register [post]
func (a *AuthAPI) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Invalid request: "+err.Error()))
		return
	}

	// Validate username format (alphanumeric + underscore)
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(req.Username)
	if !validUsername {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(400, "Username can only contain letters, numbers, and underscores"))
		return
	}

	// Check if username exists
	var existingUser model.User
	if a.db.DB.Where("username = ?", req.Username).First(&existingUser).Error == nil {
		c.JSON(http.StatusConflict, model.ErrorResponse(409, "Username already exists"))
		return
	}

	// Check if email exists
	if a.db.DB.Where("email = ?", req.Email).First(&existingUser).Error == nil {
		c.JSON(http.StatusConflict, model.ErrorResponse(409, "Email already registered"))
		return
	}

	// Create new user
	user := model.User{
		Username: req.Username,
		Password: req.Password, // Will be hashed by BeforeCreate hook
		Email:    req.Email,
		Role:     "user",
		Score:    0,
	}

	if err := a.db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to create user: "+err.Error()))
		return
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(500, "Failed to generate token"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{
		"token": token,
		"user": gin.H{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"score":      user.Score,
			"created_at": user.CreatedAt,
		},
	}))
}

// GetCurrentUser returns current user info
// @Summary Get current user info
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.APIResponse
// @Router /api/auth/me [get]
func (a *AuthAPI) GetCurrentUser(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user model.User
	if err := a.db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse(404, "User not found"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role,
		"score":      user.Score,
		"created_at": user.CreatedAt,
	}))
}
