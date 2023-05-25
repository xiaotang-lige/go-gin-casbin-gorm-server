package model

import "github.com/golang-jwt/jwt"

type Token struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	StatusId int    `json:"statusId"`
	jwt.StandardClaims
}
