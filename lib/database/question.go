package database

import (
	"QuisAPI/config"
	"QuisAPI/model"

	"gorm.io/gorm/clause"
)

func CreateQuestion(question model.Question) model.Question {
	config.DB.Create(&question)
	return question
}

func GetQuestionByID(id string) model.Question {
	var question model.Question
	config.DB.Where("id = ?", id).Find(&question)
	return question
}

func GetQuestion() []model.Question {
	var questions []model.Question
	config.DB.Preload(clause.Associations).Find(&questions)
	return questions
}

func DeleteQuestionByID(id string) {
	var question model.Question
	config.DB.Where("id = ?", id).Delete(&question)
}

func UpdateQuestionByID(id string, question model.Question) {
	config.DB.Where("id = ?", id).Updates(&question)
}
