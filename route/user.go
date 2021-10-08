package route

import (
	"QuisAPI/controller"

	"github.com/labstack/echo"
)

func NewUser(e *echo.Echo) {

	e.POST("/users/login", controller.LoginUserController)
	e.GET("/generate-hash/:password", controller.GenerateHashPassword)
	e.GET("/users", controller.GetAllUsersController)
	e.POST("/users", controller.CreateUserController)
	// //e.GET("/users/:id", controller.GetUserByIDController, middleware.AuthJWT)
	// e.DELETE("/users/:id", controller.DeleteUserByIDController)
	// e.PUT("/users/:id", controller.UpdateUserByIDController)
}
