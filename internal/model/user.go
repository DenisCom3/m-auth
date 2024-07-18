package model

import (
	"database/sql"
	"time"
)

type Role string

const (
	user  Role = "user"
	admin Role = "admin"
)

// User нужны ли геттеры?
type User struct {
	ID        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
