package models

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (user *User) SetPassword(password string) {
	hashedPassword := sha256.Sum256([]byte(password))
	user.Password = hex.EncodeToString(hashedPassword[:])
}
