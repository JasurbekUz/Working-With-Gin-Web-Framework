package apis

import (

	"app/crapp/database"
	"app/crapp/models"
)

var users []models.User
var user models.User

func PostUsersDb (newUser models.PostUser) models.User {

	db := database.DB()

	var p_user = map[string]interface{}{

		"full_name":newUser.FullName,
		"user_name":newUser.UserName,
		"user_role":newUser.UserRole,
	}

	db.Model(&user).Create(p_user)
	db.Last(&user)

	return user
}

func GetUsersDb () []models.User {

	db := database.DB()

	db.Model(&users).Find(&users)

	return users
}