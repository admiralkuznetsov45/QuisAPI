package database

import (
	"QuisAPI/config"
	"QuisAPI/model"
	"database/sql"
	"errors"
	"fmt"
)

//check login
func CheckLogin(Email, password string) (model.User, error) {
	var obj model.User

	//Ngambil Data username dari DB (DB not Found)
	con := config.DB //Error

	err := con.Where("Email = ?", Email).Find(&obj).Error

	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return obj, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return obj, err
	}

	//Salah Password
	if obj.Password != password {
		return obj, errors.New("salah gan")
	}
	return obj, nil

}

func IsValidBasic(u string) bool {
	return true
}

func IsValid(u string, p string) bool {
	var user model.User
	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
		return false
	}

	return p == user.Password
}
