package dbRoutes

import (
	"app/models"
	"app/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ChangePasswordDb(newPassword models.NewPassword) (userAcc models.UserAccount, idf bool) {

	db, err := sql.Open("postgres", utils.Db_Config())

	if err != nil {

		log.Println(" Error: Connecting databse")
		panic(err)
	}

	defer db.Close()

	idf = true

	fmt.Println(newPassword)
	fmt.Println(newPassword.NewPassword)
	fmt.Println(newPassword.UserName)
	fmt.Println(newPassword.CurrentPassword)

	err = db.QueryRow(
		utils.Update,
		newPassword.NewPassword,
		newPassword.UserName,
		newPassword.CurrentPassword,
	).
		Scan(
			&userAcc.UserName,
			&userAcc.Email,
		)

	if len(userAcc.UserName) == 0 || len(userAcc.Email) == 0 {
		//panic(err)
		idf = false
	}

	if err != nil {
		idf = false
		//panic(err)
	}

	fmt.Println(userAcc)

	return userAcc, idf
}
