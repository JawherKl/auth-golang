 // PATH: auth-golang/models/Claims.go

 package models

 import "github.com/dgrijalva/jwt-go"

 type Claims struct {
     Role string `json:"role"`
     jwt.StandardClaims
 }
