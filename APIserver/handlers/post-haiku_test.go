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

//将来的にhaikuに対してあらたなバリデーションルールを追加する可能性が高く、このテストは更新される見込み
func TestPostHaiku(t *testing.T) {

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0) //テスト用ユーザーID設定
	session_id_cnt = 0                                    //テスト用session_id設定
	users = []models.InlineObject{{Name: "post-haiku_first", Pw: "test"}}
	signupUsersForTest(users) //テスト用ユーザーの登録
	tests := []struct {
		title         string
		input         models.InlineObject2
		expected_code int
	}{
		{
			title: "Post legal haiku",
			input: models.InlineObject2{
				SessionId: "1",
				Content: models.ApiPostHaikuContent{
					First:  "a",
					Second: "b",
					Third:  "c",
				},
			},
			expected_code: http.StatusCreated,
		},
		{
			title: "Post haiku with invalid session_id",
			input: models.InlineObject2{
				SessionId: "INVALID",
				Content: models.ApiPostHaikuContent{
					First:  "a",
					Second: "b",
					Third:  "c",
				},
			},
			expected_code: http.StatusBadRequest,
		},
		{
			title: "Post haiku with empty clause",
			input: models.InlineObject2{
				SessionId: "1",
				Content: models.ApiPostHaikuContent{
					First:  "a",
					Second: "",
					Third:  "c",
				},
			},
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
			ctx.SetPath("/api/post-haiku")

			if assert.NoError(t, c.PostHaiku(ctx)) {
				assert.Equal(t, test.expected_code, rec.Code)
			}
		})
	}
}
