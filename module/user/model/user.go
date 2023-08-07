package usermodel

import (
	"errors"
	"ocean-app-be/common"
	"time"
)

const EntityName = "User"

type User struct {
	UserId         string     `json:"-" gorm:"column:user_id"`
	Username       string     `json:"username" gorm:"column:username" `
	Email          string     `json:"email" gorm:"column:email;"`
	SaltedPassword string     `json:"-" gorm:"column:salted_password" `
	FirstName      string     `json:"first_name" gorm:"column:first_name" `
	Role           string     `json:"role" gorm:"column:role" `
	LastName       string     `json:"last_name" gorm:"column:last_name" `
	CreatedAt      *time.Time `json:"date_created,omitempty" gorm:"date_created"`
	UpdatedAt      *time.Time `json:"date_updated,omitempty" gorm:"date_updated"`
}

func (u *User) GetUserID() string {
	return u.UserId

}
func (u *User) GetEmail() string {
	return u.Email

}
func (u *User) GetRole() string {
	return u.Role

}
func (u User) TableName() string {
	return "users"
}

type UserCreate struct {
	Username       string `json:"username" gorm:"column:username" `
	Email          string `json:"email" gorm:"column:email;"`
	SaltedPassword string `json:"-" gorm:"column:salted_password" `
	Role           string `json:"role" gorm:"column:role" `
	Status         int    `json:"status" gorm:"column:status;default:1;"`
	Salt           string `json:"-" gorm:"column:salt;"`
	LastName       string `json:"last_name" gorm:"column:last_name" `
	FirstName      string `json:"first_name" gorm:"column:first_name" `
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
