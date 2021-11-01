package controllers

import (
	"app/dbRoutes"
	"app/models"
	"app/utils"
	"github.com/gin-gonic/gin"
)

func ChangePassword(ctx *gin.Context) {

	newPassword := models.NewPassword{}

	ctx.BindJSON(&newPassword)

	userAcc, ok := dbRoutes.ChangePasswordDb(newPassword)

	tkn := utils.JwtTokenMaker(userAcc.UserName)

	if ok {

		token := models.Token{

			tkn,
			models.UserAccount{
				userAcc.UserName,
				userAcc.Email,
			},
		}

		ctx.JSON(200, gin.H{

			"new_access_token": token.Token,
			"user":             token.UserAccount,
		})

	} else {

		ctx.JSON(403, gin.H{

			"message": "password or username ERROR!!!",
		})
	}

}
