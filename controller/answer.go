package controller

import (
	"QuisAPI/lib/database"
	"QuisAPI/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllAnswersController(c echo.Context) error {
	answers := database.GetAnswer()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAllAnswersController",
		"data":    answers,
	})
}

func CreateAnswerController(c echo.Context) error {
	var newAnswer model.Answer
	if err := c.Bind(&newAnswer); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateAnswerController",
			"error":   err.Error(),
		})
	}

	newAnswer = database.CreateAnswer(newAnswer)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "CreateAnswerController",
		"data":    newAnswer,
	})
}

func GetAnswerByIDController(c echo.Context) error {
	id := c.Param("id")
	answer := database.GetAnswerByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAnswerByIDController",
		"data":    answer,
	})
}

func DeleteAnswerByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteAnswerByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "DeleteQuestionByIDController",
	})
}

func UpdateAnswerByIDController(c echo.Context) error {
	id := c.Param("id")

	var answer model.Answer
	if err := c.Bind(&answer); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "CreateAnswerController",
			"error":   err.Error(),
		})
	}
	database.UpdateAnswerByID(id, answer)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "GetAnswerByIDController",
		"data":    answer,
	})
}
