package handles

import (
	"log"
	"lore_project/internal/models"
	"lore_project/internal/services"
	"lore_project/pkg/auth"
	http_utils "lore_project/pkg/https"
	validation "lore_project/pkg/validatation"
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	// Dependencies or services required for authentication
	// such as UserService, EncryptionService, etc.
	UserService *services.UserService
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(service *services.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: service,
	}
}

// LoginRequest represents the data structure for the login request
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the data structure for the registration request
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// Login handles the login request
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		http_utils.BadRequest(c, "Invalid request body")
		return
	}

	// Validate the username and password (e.g., by calling UserService)
	// and generate an authentication token if the credentials are valid
	// (e.g., by calling TokenService)
	// ...
	if !validation.ValidateEmail(req.Email) {
		http_utils.BadRequest(c, "Invalid email")
		return
	}

	user, err := h.UserService.GetUserByEmail(req.Email)

	if err != nil {
		// Xử lý khi người dùng không tồn tại
		http_utils.Unauthorized(c, "Invalid credentials")
		return
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	http_utils.Unauthorized(c, "Invalid password")
	// 	return
	// }
	claims := jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	tokenString, err := auth.GenerateJWT(claims)
	if err != nil {
		log.Fatalf("Failed to generate JWT: %v", err)
	}

	// Return the authentication token in the response
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Register handles the user registration request
func (h *AuthHandler) Register(c *gin.Context) {

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		http_utils.BadRequest(c, "Invalid request body")
		return
	}

	// Validate the user registration data (e.g., by calling UserService)
	// and create a new user account if the data is valid
	// ...
	if !validation.ValidateEmail(req.Email) {
		http_utils.BadRequest(c, "Invalid email")
		return
	}
	if !validation.ValidatePassword(req.Password) {
		http_utils.BadRequest(c, "Invalid password")
		return
	}
	count, err := h.UserService.CheckEmailNotExist(req.Email)
	if err != nil {
		http_utils.InternalServerError(c, err.Error())
		return
	}
	if count > 0 {
		http_utils.Conflict(c, "Email already registered")
		return
	}
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http_utils.InternalServerError(c, "Failed to hash password")
		return
	}

	err = h.UserService.CreateUser(&models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	})
	if err != nil {
		http_utils.InternalServerError(c, "Failed create data user")
		return
	}

	// Return a success message in the response
	http_utils.SuccessMessage(c, "Registration successful")
}
