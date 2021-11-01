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

func GetClassifiedsCtrl (ctx *gin.Context) {

	ctx.IndentedJSON(http.StatusOK, DBGetClassifiedsCtrl())
	/*if err := ctx.IndentedJSON(http.StatusOK, DBGetClassifiedsCtrl()); err != nil {

		return
	}*/
}

// GET DATA FROM DATABASE
func DBGetClassifiedsCtrl () []m.GetClassified {

	db, err := sql.Open("postgres", dc.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	rows, err := db.Query(m.SELECT_ALL)

	defer rows.Close()

	if err != nil {
		log.Fatalf("db query error:\n %v", err)
	}

	var classifieds []m.GetClassified

	for rows.Next() {

		var classified m.GetClassified

		rows.Scan(
			&classified.ClassifiedId,
			&classified.Title,
			&classified.Price,
			&classified.CreatedAt,
			&classified.Category.CategoryId,
			&classified.Category.Name,
		)

		classifieds = append(classifieds, classified)
	}

	return classifieds

}