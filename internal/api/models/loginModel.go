package models

import (
	_ "fmt"
	"github.com/ysaurabh165/Trello-clone-backend/pkg/config"
)

var response config.LoginResponse

func VerifyUser(login *config.LoginRequest) (config.LoginResponse, int) {
	if val, ok := config.Credentials[login.Username]; ok {
		if val==login.Password {
			response.Code = 200
			response.Status = "Success"
			response.Msg = "User Found"
		} else {
			response.Code = 403
			response.Status = "Failure"
			response.Msg = "Invalid Username/Password"
		}	
	} else {
		response.Code = 403
		response.Status = "Failure"
		response.Msg = "Invalid Username/Password"
	}
	return response, response.Code
}

func CreateUser(login *config.LoginRequest) (config.LoginResponse, int) {
	if _, ok := config.Credentials[login.Username]; ok {
		response.Code = 400
		response.Status = "Failure"
		response.Msg = "Username Already exists"
	} else {
		config.Credentials[login.Username] = login.Password
		response.Code = 200
		response.Status = "Success"
		response.Msg = "User Created"
	}
	return response, response.Code
}