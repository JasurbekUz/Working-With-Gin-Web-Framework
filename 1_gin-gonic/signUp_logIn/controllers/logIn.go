package controllers

import (
	"app/dbRoutes"
	"app/models"
	"app/utils"
	"github.com/gin-gonic/gin"
)

func LogIn(ctx *gin.Context) {

	user := models.PostUserPassword{}

	ctx.BindJSON(&user)

	userIn, ok := dbRoutes.LogIndb(user)

	if ok {

		tkn := utils.JwtTokenMaker(userIn.UserName)

		token := models.Token{

			tkn,
			models.UserAccount{
				userIn.UserName,
				userIn.Email,
			},
		}

		ctx.JSON(200, gin.H{

			"access-token": token.Token,
			"user":         token.UserAccount,
		})

	} else {

		ctx.JSON(403, gin.H{

			"message": "password or username is ERROR!!!",
		})
	}

}
