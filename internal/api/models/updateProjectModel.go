package models

import (
	"fmt"
	"github.com/ysaurabh165/Trello-clone-backend/pkg/config"
)

func UpdateProjectContents(updatedContent *config.UpdateProjectRequest) (config.AllProjectResponse, int) {
	var response config.AllProjectResponse
	allProjects := config.AllProjects
	for key,val := range allProjects {
		if val.Title == updatedContent.Title {
			field := updatedContent.KeyUpdated
			if field == "SubTasks" {
				if x, ok := allProjects[key]; ok {
					x.SubTasks = append(x.SubTasks, updatedContent.Data)
					allProjects[key] = x
					fmt.Println(x.SubTasks)
				}
			} else if field == "Comments" {
				if x, ok := allProjects[key]; ok {
					x.Comments = append(x.Comments, updatedContent.Data)
					allProjects[key] = x
					fmt.Println(x.Comments)
				}
			} else if field == "Status" {
				if x, ok := allProjects[key]; ok {
					val.Status = updatedContent.Data
					allProjects[key] = x
					fmt.Println(x.Status)
				}
			}
			response.Code = 200
			response.Status = "Success"
			response.Msg = "Data Updated"
			response.Data = config.AllProjects
			return response, 200	
		}
	}
	response.Code = 404
	response.Status = "Failure"
	response.Msg = "Project Title Not Found"
	return response, 404
}