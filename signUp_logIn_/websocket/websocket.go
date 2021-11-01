package websocket

import (
	"app/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
)

//tuggallanmagan

func Websocket(ctx *gin.Context) {

	ctx.String(200, "Hello! It's Websocket")

	m := melody.New()

	header := models.Header{}

	if err := ctx.ShouldBindHeader(&header); err != nil {
		ctx.JSON(200, err)
	}

	cookie, err := ctx.Cookie(header.Name)

	if err != nil {

		if err == http.ErrNoCookie {
			ctx.JSON(http.StatusUnauthorized, "")
			return
		}

		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	claims := &models.Payload{}

	tkn, err := jwt.ParseWithClaims(cookie, claims,
		func(*jwt.Token) (interface{}, error) {
			return []byte(header.Name), nil
		})

	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			ctx.JSON(http.StatusUnauthorized, "")
			return
		}

		ctx.JSON(http.StatusBadRequest, "")
		return
	}

	if !tkn.Valid {
		ctx.JSON(http.StatusUnauthorized, "")
		return
	}

	m.HandleRequest(ctx.Writer, ctx.Request)

	ctx.String(200, "Hello %s", header.Name)
}
