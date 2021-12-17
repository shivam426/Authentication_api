package main

import "github.com/dgrijalva/jwt-go"

var SECRETKEY = []byte("yhqVH5UwedSwt3b9zqXWxAOIdWSfHlXyLcTXoBGbClKOom6FgVggsfCRF0FxVNK")

type Costumer struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"pwd"`

}


type Claims struct {
	Email string
	
	jwt.StandardClaims
}
