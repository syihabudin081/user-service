package models

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
}
