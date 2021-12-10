package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetHaiku(t *testing.T) {

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	c.RedisClient.Set(ctxBG, "global:next_user_id", 0, 0) //テスト用ユーザーID設定
	users = []models.InlineObject{{Name: "get-haiku_first", Pw: "test"}, {Name: "get-haiku_second", Pw: "test"}}
	signupUsersForTest(users) //テスト用ユーザーの登録

	subscribe_pair := []pair{
		{
			subscriber_session_id: models.InlineObject3{SessionId: "1"},
			receiver_user_id:      2,
		},
		{
			subscriber_session_id: models.InlineObject3{SessionId: "2"},
			receiver_user_id:      1,
		},
	}
	subscribeUserForTest(subscribe_pair)
	//テスト用ユーザーのsubscribe

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
		path_param        string
		expected_code     int
		expected_response models.InlineResponse2001
	}{
		{
			title:         "Get Second Haiku",
			path_param:    "2",
			expected_code: http.StatusOK,
			expected_response: models.InlineResponse2001{
				Author: models.User{
					UserId:            2,
					Name:              "get-haiku_second",
					Subscription:      []int64{1},
					SubscribedBy:      []int64{1},
					AuthorHaikuIdList: []int64{2},
				},
				Haiku: models.Haiku{
					HaikuId:  2,
					AuthorId: 2,
					Content: models.HaikuContent{
						First:  "d",
						Second: "e",
						Third:  "f",
					},
					Likes: 0,
					//timestampのテストはめんどいので省略
				},
			},
		},
		{
			title:         "Get first Haiku",
			path_param:    "1",
			expected_code: http.StatusOK,
			expected_response: models.InlineResponse2001{
				Author: models.User{
					UserId:            1,
					Name:              "get-haiku_first",
					Subscription:      []int64{2},
					SubscribedBy:      []int64{2},
					AuthorHaikuIdList: []int64{1},
				},
				Haiku: models.Haiku{
					HaikuId:  1,
					AuthorId: 1,
					Content: models.HaikuContent{
						First:  "a",
						Second: "b",
						Third:  "c",
					},
					Likes: 0,
					//timestampのテストはめんどいので省略
				},
			},
		},
		{
			title:             "Get Wrong user_id",
			path_param:        "WRONG",
			expected_code:     http.StatusBadRequest,
			expected_response: models.InlineResponse2001{},
		},
		{
			title:             "Get Unregistered user_id",
			path_param:        "100000",
			expected_code:     http.StatusBadRequest,
			expected_response: models.InlineResponse2001{},
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/api/haiku/:haiku_id")
			ctx.SetParamNames("haiku_id")
			ctx.SetParamValues(test.path_param)

			if assert.NoError(t, c.GetHaiku(ctx)) {
				var actual models.InlineResponse200
				json.Unmarshal(rec.Body.Bytes(), &actual)
				assert.Equal(t, test.expected_code, rec.Code)
			}
		})
	}
}
