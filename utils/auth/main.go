package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func IsValidPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func IssueToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	oneWeek := 7
	claims["user_id"] = userID
	claims["exp"] = time.Now().AddDate(oneWeek, 0, 0).Unix()

	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
