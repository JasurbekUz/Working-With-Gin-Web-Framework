package router

import (

	"github.com/gin-gonic/gin"
	"app/crapp/apis"
)

func InitRouter () *gin.Engine {

	router := gin.Default()

	router.GET("/users", apis.GetUsers)
	router.POST("/create_user", apis.PostUsers)
	router.PUT("/user/:id", apis.Update)
	/*router.DELETE("/user/:id", api.RemoveUser)*/

	return router
}