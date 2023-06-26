package utils

import (
	"github.com/heriant0/purplestore/internal/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

var cfg config.Config

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), cfg.HashCost)
	return string(hashedPassword)
}

func VerifiedPassword(password string, hanshedPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hanshedPassword),
		[]byte(password),
	)

	if err != nil {
		return false
	}

	return true
}
