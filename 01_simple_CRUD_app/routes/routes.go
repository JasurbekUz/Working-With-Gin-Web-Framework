package routes

import (
	"net/http"
	ctrl "gin_crud_app/controllers"
	"github.com/gin-gonic/gin"
)

func StartService () {

	router := gin.Default()

	api := router.Group("/api")

	{
		api.GET("/classifieds", ctrl.GetClassifiedsCtrl)
		api.GET("/classified/:id", ctrl.GetClassifiedCtrl)		
		api.POST("/classifieds", ctrl.PostClassifiedCtrl)
		api.PUT("/classifieds", ctrl.PutClassifiedCtrl)
		api.DELETE("/classifieds", ctrl.DelClassifiedCtrl)
	}

	router.NoRoute ( func (c *gin.Context) {

        c.AbortWithStatus(http.StatusNotFound)
    })

    router.Run(":8080")
}