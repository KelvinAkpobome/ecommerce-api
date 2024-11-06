package utils

import (
    "time"
    "errors"
    "github.com/golang-jwt/jwt/v4"
)


var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
    UserID  uint `json:"user_id"`
    IsAdmin bool `json:"is_admin"`
    jwt.RegisteredClaims
}

var jwtKey = []byte("secret_key")

func GenerateJWT(userID uint, isAdmin bool) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "is_admin": isAdmin,
        "exp":      time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, ErrInvalidToken
        }
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, ErrInvalidToken
    }

    return claims, nil
}