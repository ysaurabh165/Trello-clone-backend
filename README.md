# Overview
This is a POC for Trello Clone Backend.
Few of the major APIs implemented in this POC are as follows:
## 1. Login and Registration (login.go):
   This API handles the authentication of the users already present as well as creating a new user, based on two separate end points.
   In memory storage with maps is used as DB.
   
## 2. Board Contents (boardDetails.go):
   This API provides the content to populate the board page UI and the pop up card.
   In memory storage with maps is used as DB.
   
## 3. Update Project Details (updateProject.go):
   This API facilitates the following:
   i. Moving tasks from one list to another via drag and drop.
   ii. Adding SubTasks to a Project.
   iii. Adding Comments to a Project.
