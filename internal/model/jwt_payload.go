package model

import "github.com/dgrijalva/jwt-go"

const (
	ExamplePath = "user_v1.UserV1/GetById"
)

type UserClaims struct {
	jwt.StandardClaims
	Name string `json:"name"`
	Role Role   `json:"role"`
}
