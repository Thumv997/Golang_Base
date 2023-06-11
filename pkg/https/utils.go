package https

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// respondWithError is a utility function to respond with an error message and status code
func respondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

// respondWithSuccess is a utility function to respond with a success message and status code
func respondWithSuccess(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}

// setJSONResponseHeaders is a utility function to set the response headers for JSON
func setJSONResponseHeaders(c *gin.Context) {
	c.Header("Content-Type", "application/json")
}

// setCorsHeaders is a utility function to set the CORS headers
func setCorsHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
}

// handleOptionsRequest is a utility function to handle the OPTIONS request for CORS preflight
func handleOptionsRequest(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

// SuccessMessage sends a JSON response with HTTP 200 status code and a success message
func SuccessMessage(c *gin.Context, message string) {
	respondWithSuccess(c, http.StatusOK, message)
}

// BadRequest sends a JSON response with HTTP 400 status code and an error message
func BadRequest(c *gin.Context, message string) {
	respondWithError(c, http.StatusBadRequest, message)
}

// NotFound sends a JSON response with HTTP 404 status code and an error message
func NotFound(c *gin.Context, message string) {
	respondWithError(c, http.StatusNotFound, message)
}

// InternalServerError sends a JSON response with HTTP 500 status code and an error message
func InternalServerError(c *gin.Context, message string) {
	respondWithError(c, http.StatusInternalServerError, message)
}

// Unauthorized sends a JSON response with HTTP 401 status code and an error message
func Unauthorized(c *gin.Context, message string) {
	respondWithError(c, http.StatusUnauthorized, message)
}
// Conflict sends a JSON response with HTTP 409 status code and an error message
func Conflict(c *gin.Context, message string) {
	respondWithError(c, http.StatusConflict, message)
}