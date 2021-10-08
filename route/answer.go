package route

import (
	"QuisAPI/controller"
	"QuisAPI/middleware"

	"github.com/labstack/echo"
)

func NewAnswer(e *echo.Echo) {
	e.POST("/answer", controller.CreateAnswerController)
	e.GET("/answer/:id", controller.GetAnswerByIDController, middleware.JWTAuthMiddleware)
	e.GET("/answer", controller.GetAllAnswersController, middleware.JWTAuthMiddleware)
	e.DELETE("/answer/:id", controller.DeleteAnswerByIDController)
	e.PUT("/answer/:id", controller.UpdateAnswerByIDController)

}
