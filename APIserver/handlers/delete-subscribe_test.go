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

func TestDeleteSubscribe(t *testing.T) {

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0) //テスト用ユーザーID設定
	session_id_cnt = 0                                    //テスト用session_id設定
	users = []models.InlineObject{{Name: "subscribe_first", Pw: "test"}, {Name: "subscribe_second", Pw: "test"}}
	signupUsersForTest(users) //テスト用ユーザーの登録
	subscribe_pair := []pair{
		{
			subscriber_session_id: models.SessionId{SessionId: "1"},
			receiver_user_id:      2,
		},
		{
			subscriber_session_id: models.SessionId{SessionId: "2"},
			receiver_user_id:      1,
		},
	}
	subscribeUserForTest(subscribe_pair)

	tests := []struct {
		title         string
		input         models.SessionId
		path_param    string
		expected_code int
	}{
		{
			title:         "First User Remove Second User",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "2",
			expected_code: http.StatusOK,
		},
		{
			title:         "First User Remove Second User Twice",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "2",
			expected_code: http.StatusOK, //DELETEの冪等性に注意
		},
		{
			title:         "Remove myself",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "1",
			expected_code: http.StatusBadRequest,
		},
		{
			title:         "Remove Unregistered User",
			input:         models.SessionId{SessionId: "1"},
			path_param:    "100000",
			expected_code: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			inputJson, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(string(inputJson)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/subscribe/:user_id")
			ctx.SetParamNames("user_id")
			ctx.SetParamValues(test.path_param)

			if assert.NoError(t, c.DeleteSubscribe(ctx)) {
				assert.Equal(t, test.expected_code, rec.Code)
			}
		})
	}
}
