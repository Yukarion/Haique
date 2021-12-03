package handlers

import (
	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func genUUID() (uuid.UUID, error) {
	return uuid.NewRandom()
}

func isValidSessionId(session_id models.SessionId) bool { //一旦全通し
	return true
}

func hashPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hash), err
}

func isValidPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
