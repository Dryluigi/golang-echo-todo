package models

import "github.com/golang-jwt/jwt/v5"

type AuthJwtClaims struct {
	jwt.RegisteredClaims
	UserId     int      `json:"user_id"`
	UserName   string   `json:"user_name"`
	UserEmail  string   `json:"user_email"`
	UserScopes []string `json:"user_scopes"`
}
