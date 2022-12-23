package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ysaurabh165/Trello-clone-backend/pkg/config"
	"github.com/ysaurabh165/Trello-clone-backend/internal/api/models"
)

func UpdateProject(c *gin.Context) {
	var updatedContent config.UpdateProjectRequest
	var code int
	var response config.AllProjectResponse
	c.BindJSON(&updatedContent)
	fmt.Println(updatedContent)
	// fmt.Println(login.CreateUser)
	response, code = models.UpdateProjectContents(&updatedContent)
	fmt.Println(response)
	c.JSON(code, gin.H{"response": response})
}