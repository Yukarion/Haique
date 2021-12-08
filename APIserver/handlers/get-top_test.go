package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTop(t *testing.T) {

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0)  //テスト用ユーザーID設定
	c.RedisClient.Set(ctxBG, "global:next_haiku_id", 0, 0) //テスト用haiku_id設定
	for i := 0; i < TOP_HAIKUS_NUM; i++ {
		//テスト用のtop_haiku_listの初期化
		c.RedisClient.LPush(context.Background(), "global:top_haiku_id_list", DUMMY_HAIKU_ID).Result()
		c.RedisClient.RPop(context.Background(), "global:top_haiku_id_list")
	}
	session_id_cnt = 0 //テスト用session_id設定
	users = []models.InlineObject{{Name: "get-top_first", Pw: "test"}, {Name: "get-top_second", Pw: "test"}}
	signupUsersForTest(users) //テスト用ユーザーの登録
	haiku_list = []models.InlineObject2{
		{
			SessionId: "1",
			Content: models.ApiPostHaikuContent{
				First:  "a",
				Second: "b",
				Third:  "c",
			},
		},
		{
			SessionId: "2",
			Content: models.ApiPostHaikuContent{
				First:  "d",
				Second: "e",
				Third:  "f",
			},
		},
	}
	postHaikusForTest(haiku_list)

	tests := []struct {
		title             string
		expected_code     int
		expected_response []models.Haiku
	}{
		{
			title:         "Get top haikus",
			expected_code: http.StatusOK,
			expected_response: []models.Haiku{
				//後にPOSTされた方が上位に出ることに注意
				{
					HaikuId:  2,
					AuthorId: 2,
					Content: models.HaikuContent{
						First:  "d",
						Second: "e",
						Third:  "f",
					},
					Likes: 0,
					//timestampまわりのテストはコスパが悪すぎるので省略
				},
				{
					HaikuId:  1,
					AuthorId: 1,
					Content: models.HaikuContent{
						First:  "a",
						Second: "b",
						Third:  "c",
					},
					Likes: 0,
					//timestampまわりのテストはコスパが悪すぎるので省略
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/get-top")

			if assert.NoError(t, c.GetTop(ctx)) {
				var actual []models.Haiku
				json.Unmarshal(rec.Body.Bytes(), &actual)
				assert.Equal(t, test.expected_code, rec.Code)
				assert.Equal(t, len(test.expected_response), len(actual))
				assert.NotEqual(t, test.expected_response[0].AuthorId, actual[1].AuthorId)
				assert.Equal(t, test.expected_response[1].Content.Second, actual[1].Content.Second)
			}
		})
	}
}
