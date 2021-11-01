// sana: 2021 11 1
// dasturchi: Jasurbek Shamsiddinov
// maqsad: gin web fremwork ini o'rganish

package main

import (

	"log"
	"github.com/gin-gonic/gin"
)

func main() {

    router := gin.Default()

    router.GET("/albums", GetAlbums)
    router.GET("/albums/:id", GetAlbumByID)
    router.POST("albums", PostAlbums)

    if err := router.Run("localhost:8080"); err != nil {

    	log.Fatalf("server connection error: %v", err)
    }
}

