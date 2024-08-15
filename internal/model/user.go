package model

import (
	"database/sql"
	"time"
)

type Role int

const (
	UserRole  Role = 0
	AdminRole Role = 1
)

type CreateUser struct {
	Info     UserInfo `json:"info"`
	Password string   `json:"password"`
}

type UpdateUser struct {
	ID   int64
	Info UserInfo
}

type User struct {
	ID        int64
	Info      UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
