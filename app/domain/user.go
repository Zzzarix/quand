package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

type User struct {
	// username is an unique
	UserName string `json:"userName"`
	Password string `json:"password"`

	// Main user identifier is an email
	Email      string   `json:"email"`
	Id         uint64   `json:"id"`
	History    History  `json:"history"`
	Keys       []string `json:"keys"`
	StaffLevel int      `json:"staffLevel"`
}

func HashPwd(pwd string) string {
	h := sha256.New()
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

func CheckPwd(pwd string, user User) bool {
	return HashPwd(pwd) == user.Password
}
