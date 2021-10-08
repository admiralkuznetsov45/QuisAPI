package controller

import (
	"QuisAPI/lib/database"
	"QuisAPI/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllQuestionsController(c echo.Context) error {
	questions := database.GetQuestion()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllQuestionsController",
		"data":    questions,
	})
}

func CreateQuestionController(c echo.Context) error {
	var newQuestion model.Question
	if err := c.Bind(&newQuestion); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateQuestionController",
			"error":   err.Error(),
		})
	}

	newQuestion = database.CreateQuestion(newQuestion)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateQuestionController",
		"data":    newQuestion,
	})
}

func GetQuestionByIDController(c echo.Context) error {
	id := c.Param("id")
	question := database.GetQuestionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetQuestionByIDController",
		"data":    question,
	})
}

func DeleteQuestionByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteQuestionByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteQuestionByIDController",
	})
}

func UpdateQuestionByIDController(c echo.Context) error {
	id := c.Param("id")

	var question model.Question
	if err := c.Bind(&question); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateQuestionController",
			"error":   err.Error(),
		})
	}
	database.UpdateQuestionByID(id, question)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetQuestionByIDController",
		"data":    question,
	})
}
