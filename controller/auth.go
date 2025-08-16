package controller

import (
	"net/http"
	"time"

	"zip_archive/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// User credentials for demo (in real app — use DB)

// LoginRequest represents login payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login generates JWT token
func (c *Controller) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Username != config.GetAdminUsername() || req.Password != config.GetAdminPassword() {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT claims
	claims := &jwt.RegisteredClaims{
		Subject:   req.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.TokenExpiry)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
signedToken, err := token.SignedString([]byte(config.GetJWTSecret())) // Берём из .env
if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
    return
}

	ctx.JSON(http.StatusOK, gin.H{
		"token":     signedToken,
		"expiresIn": config.TokenExpiry.Seconds(),
	})
}