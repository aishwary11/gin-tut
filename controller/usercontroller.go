package controller

import (
	"context"
	"net/http"
	"time"

	"gin.com/aishwary11/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login(c *gin.Context) {
	var user utils.User
	if err := c.BindJSON(&user); err != nil {
		utils.ResponseHelper(c, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	collection := utils.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var foundUser utils.User
	err := collection.FindOne(ctx, bson.M{"email": user.Email, "name": user.Name}).Decode(&foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.ResponseHelper(c, http.StatusUnauthorized, "Invalid credentials", nil)
		} else {
			utils.ResponseHelper(c, http.StatusInternalServerError, "Error querying database", nil)
		}
		return
	}
	token, err := utils.GenerateToken(foundUser)
	if err != nil {
		utils.ResponseHelper(c, http.StatusInternalServerError, "Error generating token", nil)
		return
	}
	utils.ResponseHelper(c, http.StatusOK, "Login successful", map[string]string{"token": token})
}
