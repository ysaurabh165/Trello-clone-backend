package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysaurabh165/Trello-clone-backend/pkg/config"
	"github.com/ysaurabh165/Trello-clone-backend/internal/api/models"
)

func GetLogin(c *gin.Context) {
	var login config.LoginRequest
	var code int
	var response config.LoginResponse
	c.BindJSON(&login)
	fmt.Println(login.Username)
	fmt.Println(login.CreateUser)
	if login.CreateUser == 1 {
		response, code = models.CreateUser(&login)
	} else {
		response, code = models.VerifyUser(&login)
	}
	fmt.Println(response)
	c.JSON(code, gin.H{"response": response})
}