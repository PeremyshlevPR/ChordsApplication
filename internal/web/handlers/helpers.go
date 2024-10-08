package handlers

import (
	"chords_app/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(c *gin.Context, requestSchema interface{}, validate *validator.Validate) bool {
	if err := c.BindJSON(requestSchema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	if err := validate.Struct(requestSchema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}

	return true
}

func parseUintParam(c *gin.Context, param string) (uint, error) {
	idStr := c.Param(param)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func parseUintQueryParam(c *gin.Context, param string, defaultValue uint) (uint, error) {
	ParamStr := c.Query(param)
	if ParamStr == "" {
		return defaultValue, nil
	}

	value, err := strconv.ParseUint(ParamStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

func GetUserModel(c *gin.Context) (*models.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "access token not passed ar invalid"})
		return &models.User{}, exists
	}
	userModel := user.(*models.User)
	return userModel, exists
}
