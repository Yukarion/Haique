package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetTimeline(t *testing.T) {

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0)  //テスト用ユーザーID設定
	c.RedisClient.Set(ctxBG, "global:next_haiku_id", 0, 0) //テスト用haiku_id設定
	session_id_cnt = 0                                     //テスト用session_id設定
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
	log.Println(c.RedisClient.SMembers(ctxBG, "user_id:1:subscription").Result())
	delete_subscribe_pair := []pair{
		{
			subscriber_session_id: models.SessionId{SessionId: "2"},
			receiver_user_id:      1,
		},
	}
	deleteSubscribeUserForTest(delete_subscribe_pair)

	tests := []struct {
		title             string
		input             models.InlineObject5
		expected_code     int
		expected_response []models.Haiku
	}{
		{
			title:         "Get First User's timeline (Following Second User)",
			input:         models.InlineObject5{SessionId: "1"},
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
		{
			title:         "Get Second User's timeline (Following Nobody)",
			input:         models.InlineObject5{SessionId: "2"},
			expected_code: http.StatusOK,
			expected_response: []models.Haiku{
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
			},
		},
		{
			title:             "Get timeline with invalid session_id",
			input:             models.InlineObject5{SessionId: "Invalid"},
			expected_code:     http.StatusBadRequest,
			expected_response: []models.Haiku{},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			inputJson, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(inputJson)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/get-timeline")

			if assert.NoError(t, c.GetTimeline(ctx)) {
				var actual []models.Haiku
				json.Unmarshal(rec.Body.Bytes(), &actual)
				assert.Equal(t, test.expected_code, rec.Code)
				assert.Equal(t, len(test.expected_response), len(actual))
			}
		})
	}
}
