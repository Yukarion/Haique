package handlers

import (
	"strings"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func genUUID() (string, error) {
	uuid, err := uuid.NewRandom()
	return uuid.String(), err

}

func isValidSessionId(session_id models.SessionId) bool { //一旦全通し
	//これいらないかも。
	/*
		author_id, err := c.RedisClient.Get(ctxBG, session_id.Id.String()+":linked_user_id").Result()
		if err != nil {
			return ctx.HTML(http.StatusBadRequest, "invalid session id")
		}
		で完全に代替可能
	*/
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

func eraseSpaceInContent(content models.ApiPostHaikuContent) models.ApiPostHaikuContent {
	content.First = strings.ReplaceAll(content.First, " ", "")
	content.Second = strings.ReplaceAll(content.Second, " ", "")
	content.Third = strings.ReplaceAll(content.Third, " ", "")

	return content
}
