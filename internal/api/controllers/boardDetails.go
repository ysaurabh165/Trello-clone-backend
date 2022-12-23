package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "github.com/ysaurabh165/Trello-clone-backend/pkg/config"
	"github.com/ysaurabh165/Trello-clone-backend/internal/api/models"
)

func GetBoardContent(c *gin.Context) {
	response, code := models.GetAllProjects()
	fmt.Println(response)
	c.JSON(code, gin.H{"response": response})
}