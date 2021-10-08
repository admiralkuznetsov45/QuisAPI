package database

import (
	"QuisAPI/config"
	"QuisAPI/model"

	"gorm.io/gorm/clause"
)

func CreateAnswer(answer model.Answer) model.Answer {
	config.DB.Create(&answer)
	return answer
}

func GetAnswerByID(id string) model.Answer {
	var answer model.Answer
	config.DB.Where("id = ?", id).Find(&answer)
	return answer
}

func GetAnswer() []model.Answer {
	var answers []model.Answer
	config.DB.Preload(clause.Associations).Find(&answers)
	return answers
}

func DeleteAnswerByID(id string) {
	var answer model.Answer
	config.DB.Where("id = ?", id).Delete(&answer)
}

func UpdateAnswerByID(id string, answer model.Answer) {
	config.DB.Where("id = ?", id).Updates(&answer)
}
