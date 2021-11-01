package controllers

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	dc "gin_crud_app/dbConfig"
	m "gin_crud_app/models"	
	
)

func PutClassifiedCtrl (ctx *gin.Context) {

	var changedClassified m.PutClassified
	
	ctx.BindJSON(&changedClassified)

	ctx.IndentedJSON(http.StatusOK, DBPutClassifiedCtrl(changedClassified))
	/*if err := ctx.IndentedJSON(http.StatusOK, DBPutClassifiedCtrl(changedClassified)); err != nil {

		return
	}*/
}

func DBPutClassifiedCtrl (chClassified m.PutClassified) m.GetClassified {

	db, err := sql.Open("postgres", dc.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified m.GetClassified

	err = db.QueryRow(
				m.UPDATE,
				chClassified.ClassifiedId,
				chClassified.Title,
				chClassified.Price,
				chClassified.CategoryId,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				//&classified.Category.Name,
		)

	fmt.Println(classified)

	if err != nil {
		return classified
		log.Fatalf("db query error:\n %v", err)
	}

	return classified

}