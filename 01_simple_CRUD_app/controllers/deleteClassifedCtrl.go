package controllers

import (
	"log"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	dc "gin_crud_app/dbConfig"
	m "gin_crud_app/models"

)


func DelClassifiedCtrl(ctx *gin.Context) {

	type DelClsfdId struct {
		Id uint16	`json:"id"`
	}

	var id DelClsfdId	

	ctx.BindJSON(&id)

	ctx.IndentedJSON(http.StatusOK, DBDeleteClassifiedCtrl(id.Id))
	/*if err := ctx.IndentedJSON(http.StatusOK, DBDeleteClassifiedCtrl(id.Id)); err != nil {

		return
	}*/
}

func DBDeleteClassifiedCtrl (id uint16) m.GetClassified {

	db, err := sql.Open("postgres", dc.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified m.GetClassified

	err = db.QueryRow(
				m.DELETE,
				id,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				//&classified.Category.Name,
		)

	if err != nil {
		return classified
		//log.Fatalf("db query error:\n %v", err)
	}

	return classified

}