package model

type User struct {
	// gorm.Model
	ID uint `gorm:"primarykey"`
	//Question []*Question `json:"question"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}
