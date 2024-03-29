package schemas

import "github.com/golang-jwt/jwt"

//Token struct declaration
type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}
