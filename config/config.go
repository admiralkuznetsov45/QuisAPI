package config

import (
	"QuisAPI/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//var DBLog *mongo.Client

func InitDB() {
	var err error
	//                              username:password@tcp(host:port)/nama_database
	DB, err = gorm.Open(mysql.Open("admin:ayamjantan@tcp(quis-api.cu1pwj6ylp2t.us-east-2.rds.amazonaws.com:3306)/quisapi?parseTime=True"))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB connected")
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(
		&model.Answer{},
		&model.Question{},
		&model.User{},
	)

}
