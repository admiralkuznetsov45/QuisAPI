package model

type Question struct {
	ID uint `gorm:"primarykey"`
	//User          *User
	QuestionType  string `json:"QuestionType"`
	QuestionPoint uint   `json:"QuestionPoint"`
	Text          string `json:"Text"`
	//Answer        []*Answer `gorm:"many2many:answers" json:"answers,omitempty"`
}
