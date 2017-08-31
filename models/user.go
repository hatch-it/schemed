package models

// User contains all user-related data.
type User struct {
	Model
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
