package main

import (
	//"fmt"
	ctrl "app/controllers"
	ws "app/websocket"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {

		log.Fatal(err)
	}

	port := os.Getenv("APP_HTTP_PORT")

	httpRouter := gin.Default()

	httpRouter.POST("/signup", ctrl.SignUp)
	httpRouter.POST("/login", ctrl.LogIn)
	httpRouter.POST("/change_password", ctrl.ChangePassword)

	httpRouter.GET("/echo", ws.Websocket)

	httpRouter.Run(port)

}
