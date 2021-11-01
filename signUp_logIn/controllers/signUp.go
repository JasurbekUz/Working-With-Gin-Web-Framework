package controllers

import (
	"app/dbRoutes"
	"app/models"
	"app/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {

	newUser := models.UserAccount{}

	ctx.BindJSON(&newUser)

	password := utils.CodeGenerator()

	utils.SendCode(newUser.Email, password)

	addedUser, ok := dbRoutes.SignUpDb(newUser, password)

	if ok {

		ctx.JSON(200, gin.H{

			"userId":   addedUser.UserId,
			"userName": addedUser.UserName,
			"email":    addedUser.Email,
			"password": addedUser.Password,
		})

	} else {

		ctx.JSON(403, gin.H{

			"message": "userName or email is already exists",
		})
	}

}
