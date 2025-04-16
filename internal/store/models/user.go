package models

type User struct {
	ID       uint   `json:"-" gorm:"primaryKey" binding:"required"`
	Username string `json:"username" gorm:"uniqueIndex" binding:"required"`
	Password string `json:"-"`
}
