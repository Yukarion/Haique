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

func TestPostSubscribe(t *testing.T) {

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0) //テスト用ユーザーID設定
	session_id_cnt = 0                                    //テスト用session_id設定
	users = []models.InlineObject{{Name: "post-subscribe_first", Pw: "test"}, {Name: "post-subscribe_second", Pw: "test"}}
	signupUsersForTest(users) //テスト用ユーザーの登録
	tests := []struct {
		title         string
		input         models.SessionId
		path_param    string
		expected_code int
	}{
		{
			title:         "First User Subscribes Second User",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "2",
			expected_code: http.StatusOK,
		},
		{
			title:         "Subscribe with wrong session_id",
			input:         models.SessionId{SessionId: "WRONG"},
			path_param:    "2",
			expected_code: http.StatusBadRequest,
		},
		{
			title:         "Subscribe myself",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "1",
			expected_code: http.StatusBadRequest,
		},
		{
			title:         "Subscribe Unregistered User",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "100000",
			expected_code: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			inputJson, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(inputJson)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/subscribe/:user_id")
			ctx.SetParamNames("user_id")
			ctx.SetParamValues(test.path_param)

			if assert.NoError(t, c.PostSubscribe(ctx)) {
				assert.Equal(t, test.expected_code, rec.Code)
			}
		})
	}
}
