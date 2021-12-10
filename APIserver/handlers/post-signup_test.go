package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var session_id_cnt int = 0

func genUUIDForTest() (string, error) { //テスト時、session_idは1から順番に振る
	session_id_cnt++
	return strconv.Itoa(session_id_cnt), nil
}

func TestSignup(t *testing.T) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0) //テスト用ユーザーID設定
	session_id_cnt = 0                                    //テスト用session_id設定
	tests := []struct {
		title             string
		input             models.InlineObject
		expected_code     int
		expected_response models.InlineObject3
	}{
		{
			title:             "Create First User",
			input:             models.InlineObject{Name: "signup_first", Pw: "test"},
			expected_code:     http.StatusCreated,
			expected_response: models.InlineObject3{SessionId: "1"},
		},
		{
			title:             "Username Conflict",
			input:             models.InlineObject{Name: "signup_first", Pw: "test"},
			expected_code:     http.StatusConflict,
			expected_response: models.InlineObject3{SessionId: ""},
		},
		{
			title:             "Create Second User",
			input:             models.InlineObject{Name: "signup_second", Pw: "test"},
			expected_code:     http.StatusCreated,
			expected_response: models.InlineObject3{SessionId: "2"},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			inputJson, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(inputJson)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/signup")

			if assert.NoError(t, c.PostSignup(ctx)) {
				var actual models.InlineObject3
				json.Unmarshal(rec.Body.Bytes(), &actual)
				assert.Equal(t, test.expected_code, rec.Code)
				assert.Equal(t, test.expected_response.SessionId, actual.SessionId)
			}
		})
	}
}
