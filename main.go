package main

import (
	"QuisAPI/config"
	"QuisAPI/lib/database"
	"QuisAPI/route"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func main() {
	config.InitDB()

	e := echo.New()
	route.NewUser(e)
	route.NewQuestion(e)
	route.NewAnswer(e)

	// e.GET(
	// 	"/",
	// 	func(c echo.Context) error {
	// 		email := c.Get("email")
	// 		return c.String(http.StatusOK, fmt.Sprintf("selamat datang %s", email))
	// 	},
	// 	middleware.JWTAuthMiddleware,
	// )

	e.POST("/ambiltoken", func(c echo.Context) error {
		user := struct {
			Email    string
			Password string
		}{}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		isValid := database.IsValid(user.Email, user.Password)
		if !isValid {
			return c.String(http.StatusBadRequest, "email atau password salah")
		}

		claims := jwt.MapClaims{}
		claims["userId"] = user.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("asdasdasd"))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, tokenString)
	})

	e.Start(":8080")
}
