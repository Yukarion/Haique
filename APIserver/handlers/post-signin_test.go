package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func signupUsersForTest(users []models.InlineObject) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	for _, user := range users {
		inputJson, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(string(inputJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		c.PostSignup(ctx)
	}
}

var users []models.InlineObject

func TestSignin(t *testing.T) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0) //テスト用ユーザーID設定
	session_id_cnt = 0                                    //テスト用session_id設定
	users = []models.InlineObject{{Name: "signin_first", Pw: "test"}, {Name: "signin_second", Pw: "test"}}
	signupUsersForTest(users) //テスト用ユーザーの登録
	tests := []struct {
		title             string
		input             models.InlineObject
		expected_code     int
		expected_response models.SessionId
	}{
		{
			title:             "Signin First User",
			input:             models.InlineObject{Name: "signin_first", Pw: "test"},
			expected_code:     http.StatusOK,
			expected_response: models.SessionId{SessionId: "3"},
		},
		{
			title:             "Signin with wrong name",
			input:             models.InlineObject{Name: "UNKNOWN", Pw: "UNKNOWN"},
			expected_code:     http.StatusBadRequest,
			expected_response: models.SessionId{SessionId: ""},
		},
		{
			title:             "Signin with wrong password",
			input:             models.InlineObject{Name: "signin_first", Pw: "WRONG_PASSWORD"},
			expected_code:     http.StatusBadRequest,
			expected_response: models.SessionId{SessionId: ""},
		},
		{
			title:             "Signin Second User",
			input:             models.InlineObject{Name: "signin_second", Pw: "test"},
			expected_code:     http.StatusOK,
			expected_response: models.SessionId{SessionId: "4"},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			inputJson, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(inputJson)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/signin")

			if assert.NoError(t, c.PostSignin(ctx)) {
				var actual models.SessionId
				json.Unmarshal(rec.Body.Bytes(), &actual)
				assert.Equal(t, test.expected_code, rec.Code)
				assert.Equal(t, test.expected_response.SessionId, actual.SessionId)
			}
		})
	}
}
