package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/habit-tracker-api/config"
	_exceptionType "github.com/habit-tracker-api/enums/exception"
	_exception "github.com/habit-tracker-api/exception"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func HashToken(info interface{}, exp int) (string, time.Time, error) {

	expDate := time.Now().Add(time.Hour * time.Duration(exp)) // exp in hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"info": info,
		"exp":  expDate.Unix(), // exp in hours
	})

	secret := config.ConfigGetting().Security.JwtSecret

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", time.Now(), _exception.Handle("token invalid", _exceptionType.Info)
	}

	return tokenString, expDate, err

}
