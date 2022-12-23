package config

// import (
// 	"fmt"
// )

type LoginRequest struct {
	Username		string `json:"username"`
	Password		string `json:"password"`
	CreateUser		int `json:"createUser"`
}

type LoginResponse struct {
	Code      		int	`json:"code,omitempty"`
	Status   	 	string `json:"status,omitempty"`
	Msg       		string `json:"msg,omitempty"`
}

var Credentials = map[string]string {
	"user1": "password1",
	"user2": "password2",
}

type Projects struct {
	Title			string `json:"title"`
	Status			string `json:"status"`
	Category		string `json:"category"`
	StartDate		string `json:"startDate"`
	SubTasks		[]string `json::"subTasks"`
	Likes			int `json:"likes"`
	Comments		[]string `json:"comments"`
	
}

type UpdateProjectRequest struct {
	Title			string `json:"title"`
	KeyUpdated		string `json:"keyUpdated"`
	Data			string `json:"data"`
}

var AllProjects = map[string]Projects {
	"1":{Title:"Performance Improvements", Status:"Backlog", Category:"Engineering", StartDate:"", SubTasks:[]string{}, Likes:3, Comments: []string{"I will look into it.","Thanks"}},
	"2":{Title:"Desing Prototype", Status:"Ready", Category:"Design", StartDate:"Jan 16", SubTasks:[]string{"Design review", "Testing plan"}, Likes	:1, Comments: []string{"Great work.","Good.","Looking forward to it."}},
	"3":{Title:"Redesign Overview", Status:"In Progress", Category:"Engineering", StartDate:"Nov 15", SubTasks:[]string{}, Likes:6, Comments: []string{"Looking forward to it."}},
}

type AllProjectResponse struct {
	Code      		int	`json:"code,omitempty"`
	Status   	 	string `json:"status,omitempty"`
	Msg       		string `json:"msg,omitempty"`
	Data			map[string]Projects `json:"data,omitempty"`
}
