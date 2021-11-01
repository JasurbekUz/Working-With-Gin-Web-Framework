package dbRoutes

import (
	"app/models"
	"app/utils"
	"database/sql"
	_ "github.com/lib/pq"
)

func SignUpDb(body models.UserAccount, password string) (models.SignedUser, bool) {

	db, err := sql.Open("postgres", utils.Db_Config())

	defer db.Close()

	if err != nil {
		panic(err)
	}

	idf := true

	signedUser := models.SignedUser{}

	err = db.QueryRow(
		utils.Insert,
		body.UserName,
		body.Email,
		password,
	).Scan(
		&signedUser.UserId,
		&signedUser.UserName,
		&signedUser.Email,
		&signedUser.Password,
	)

	if err != nil {
		idf = false
		//panic(err)
	}

	return signedUser, idf
}
