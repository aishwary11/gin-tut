package controller

import (
	"net/http"

	"gin.com/aishwary11/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user utils.User
	if err := c.BindJSON(&user); err != nil {
		utils.ResponseHelper(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	if user.Email == "aish@gmail.com" && user.Name == "aish" {
		token, err := utils.GenerateToken(user)
		if err != nil {
			utils.ResponseHelper(c, http.StatusInternalServerError, "Error generating token", nil)
			return
		}
		utils.ResponseHelper(c, http.StatusOK, "Login successful", token)
	} else {
		utils.ResponseHelper(c, http.StatusUnauthorized, "Invalid credentials", nil)
	}
}
