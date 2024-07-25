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
	Info     UserInfo
	Password string
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
	Name  string
	Email string
	Role  Role
}
