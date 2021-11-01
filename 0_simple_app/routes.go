package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAlbums (ctx *gin.Context) {

    ctx.IndentedJSON(http.StatusOK, Albums)
}

func PostAlbums (ctx *gin.Context) {

	var newAlbum Album

	if err := ctx.BindJSON(&newAlbum); err != nil {

	}

	fmt.Println(newAlbum)

	newAlbum.Id = GetId()

	Albums = append(Albums, newAlbum)

	ctx.IndentedJSON(http.StatusCreated, newAlbum)

}

func GetAlbumByID(ctx *gin.Context) {

    id := ctx.Param("id")


    for _, album := range Albums {

    	uId := fmt.Sprintf("%s", album.Id)

        if uId == id {

            ctx.IndentedJSON(http.StatusOK, album)

            return
        }
    }

    ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
