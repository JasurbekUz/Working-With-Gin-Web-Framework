package controllers

import (
	"log"
	"net/http"
	"strconv"
	"database/sql"
	_"github.com/lib/pq"
	"github.com/gin-gonic/gin"
	dc "gin_crud_app/dbConfig"
	m "gin_crud_app/models"	
)

func GetClassifiedCtrl (ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	ctx.IndentedJSON(http.StatusOK, DBGetClassifiedCtrl(id))
	/*if err := ctx.IndentedJSON(http.StatusOK, DBGetClassifiedCtrl(id)); err != nil {

		return
	}*/
}

func DBGetClassifiedCtrl (id int) m.GetClassified {

	db, err := sql.Open("postgres", dc.DB_CONFIG)

	defer db.Close()

	if err != nil {
		log.Fatalf("db connection error:\n %v", err)
	}

	var classified m.GetClassified

	err = db.QueryRow(
				m.SELECT_ONCE,
				id,
			).Scan(
				&classified.ClassifiedId,
				&classified.Title,
				&classified.Price,
				&classified.CreatedAt,
				&classified.Category.CategoryId,
				&classified.Category.Name,
		)

	if err != nil {
		//log.Fatalf("db query error:\n %v", err)
	}

	return classified

}