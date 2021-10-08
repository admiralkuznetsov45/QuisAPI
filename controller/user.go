package controller

import (
	"QuisAPI/config"
	"QuisAPI/helpers"
	"QuisAPI/lib/database"
	"QuisAPI/model"

	"net/http"

	"github.com/labstack/echo"
)

//response
type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func GetAllUsersController(c echo.Context) error {
	//membuat model tampung dalam bentuk array
	var newUser []model.User

	//Mencari User
	result := config.DB.Find(&newUser)

	//Apabila Error
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}

	return c.JSON(http.StatusOK, Response{Status: http.StatusOK, Success: true, Message: "OK", Content: newUser})
}

func CreateUserController(c echo.Context) error {
	//membuat model user
	var newUser model.User

	//ambil dari nilai form postman
	newUser.Email = c.FormValue("email")
	newUser.Password = c.FormValue("password")

	//bikinDB
	result := config.DB.Create(&newUser)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, Response{Status: http.StatusOK, Success: true, Message: "OK", Content: newUser})
}

func LoginUserController(c echo.Context) error {

	var loginUser model.User

	loginUser.Email = c.FormValue("email")
	loginUser.Password = c.FormValue("password")

	//ambil data dari database kalau sukses
	user, err := database.CheckLogin(loginUser.Email, loginUser.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, Response{Status: http.StatusOK, Success: true, Message: "OK", Content: user})
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
