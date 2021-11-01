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

func PostClassifiedCtrl (ctx *gin.Context) {

	var newClassified m.PostClassifed

	err := ctx.BindJSON(&newClassified)

       if err != nil {
            return
       }

       ctx.IndentedJSON(http.StatusOK, DBPostClassifiedCtrl(newClassified))
	/*if err := ctx.IndentedJSON(http.StatusOK, DBPostClassifiedCtrl(newClassified)); err != nil {

		return
	}*/
}

func DBPostClassifiedCtrl (newClassified m.PostClassifed) m.GetClassified {

	db, err := sql.Open("postgres", dc.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified m.GetClassified

	fmt.Println(newClassified)

	err = db.QueryRow(
				m.INSERT,
				newClassified.Title,
				newClassified.Price,
				newClassified.CategoryId,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				//&classified.Category.Name,
		)

	if err != nil {
		//log.Fatalf("db query error:\n %v", err)
	}

	return classified

}