package middleware

import (
	"QuisAPI/lib/database"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Basic base64string-blablala=
		authorizationFromHeader := c.Request().Header.Get("authorization")

		// Replace Basic, Get base64string-blablala=
		base64data := strings.ReplaceAll(authorizationFromHeader, "Basic ", "")

		// Get data from base64string-blablala=
		data, err := base64.StdEncoding.DecodeString(base64data)
		if err != nil {
			return c.String(http.StatusForbidden, "email atau password salah")
		}

		// Check is data valid
		if database.IsValidBasic(string(data)) {
			return next(c)
		}

		return c.String(http.StatusForbidden, "email atau password salah")
	}
}

func JWTAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bearer token-blablala=
		authorizationFromHeader := c.Request().Header.Get("authorization")

		// Replace Bearer, Get token-blablala=
		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("asdasdasd"), nil
		})
		if err != nil {
			return c.String(http.StatusForbidden, "token salah")
		}

		c.Set("email", claims["userId"])
		return next(c)
	}
}
