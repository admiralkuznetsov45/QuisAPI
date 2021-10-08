package route

import (
	"QuisAPI/controller"
	"QuisAPI/middleware"

	"github.com/labstack/echo"
)

func NewQuestion(e *echo.Echo) {
	e.POST("/question", controller.CreateQuestionController)
	e.GET("/question/:id", controller.GetQuestionByIDController, middleware.JWTAuthMiddleware)
	e.GET("/question", controller.GetAllQuestionsController, middleware.JWTAuthMiddleware)
	e.DELETE("/question/:id", controller.DeleteQuestionByIDController)
	e.PUT("/question/:id", controller.UpdateQuestionByIDController)

}
