package model

import (
	"time"
)

const (
	GeneralUser = 1
	AdminUser   = 99
)

type User struct {
	ID                     string `gorm:"primaryKey"`
	Email                  string
	ChangeEmail            *string
	Type                   int
	EmailVerifiedAt        *time.Time
	EmailVerificationToken string
	Password               string
	RememberToken          *string
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
}

func (u *User) TypeAsString() string {
	switch u.Type {
	case GeneralUser:
		return "一般"
	case AdminUser:
		return "管理者"
	default:
		return "未定義"
	}
}

func (u *User) IsGeneral() bool {
	return u.Type == GeneralUser
}

func (u *User) IsAdmin() bool {
	return u.Type == AdminUser
}

func (u *User) IsEmailVerified() bool {
	return u.EmailVerifiedAt != nil
}
