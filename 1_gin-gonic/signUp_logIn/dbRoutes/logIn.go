package dbRoutes

import (
	"app/models"
	"app/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func LogIndb(user models.PostUserPassword) (models.UserAccount, bool) {

	db, err := sql.Open("postgres", utils.Db_Config())

	if err != nil {

		log.Println(" Error: Connecting databse")
		panic(err)
	}

	defer db.Close()

	idf := true

	userAcc := models.UserAccount{}

	err = db.QueryRow(
		utils.Select,
		user.UserName,
		user.Password,
	).
		Scan(
			&userAcc.UserName,
			&userAcc.Email,
		)

	if len(userAcc.UserName) == 0 || len(userAcc.Email) == 0 {

		idf = false
	}

	if err != nil {
		idf = false
		//panic(err)
	}

	fmt.Println(userAcc)

	return userAcc, idf
}
