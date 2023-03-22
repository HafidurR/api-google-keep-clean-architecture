package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
	Notes []Note
}

type Note struct{
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content`
	User_id int 	`json:"user_id"`
	User User
}

type UserCreate struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}