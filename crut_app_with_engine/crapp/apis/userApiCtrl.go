package apis

import (

	"net/http"

	"github.com/gin-gonic/gin"

	"app/crapp/models"
)

var newUser models.PostUser 

func PostUsers ( ctx *gin.Context) {

	ctx.BindJSON(&newUser)

	ctx.JSON(http.StatusOK, PostUsersDb(newUser))
}

func GetUsers ( ctx *gin.Context) {

	ctx.JSON(http.StatusOK, GetUsersDb())
}