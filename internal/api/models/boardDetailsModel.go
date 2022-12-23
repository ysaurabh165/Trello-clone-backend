package models

import (
	_ "fmt"
	"github.com/ysaurabh165/Trello-clone-backend/pkg/config"
)

func GetAllProjects() (config.AllProjectResponse, int) {
	var response config.AllProjectResponse
	response.Code = 200
	response.Status = "Success"
	response.Msg = "Data Found"
	response.Data = config.AllProjects
	return response, 200
}