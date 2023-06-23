package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)
var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
// Tạo JWT với claims đã cho và secret key
func GenerateJWT(claims jwt.Claims) (string, error) {
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// Xác thực JWT và trả về các claims nếu hợp lệ
func VerifyJWT(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, jwt.ErrInvalidKey
    }
    return token, nil
}