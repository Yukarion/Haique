package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/Mackyson/Haique/APIserver/models"
	"github.com/labstack/echo/v4"
)

//utils.goのテストではなくテスト用のUtilityをここに置く。
//末尾が_testだとビルド時には含まれないらしいからこういう名前にした
var (
	haiku_list []models.InlineObject2
	users      []models.InlineObject
)

func signupUsersForTest(users []models.InlineObject) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	for _, user := range users {
		inputJson, _ := json.Marshal(user)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(inputJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/signup")
		c.PostSignup(ctx)
	}
}

func postHaikusForTest(haiku_list []models.InlineObject2) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	for _, haiku := range haiku_list {
		inputJson, _ := json.Marshal(haiku)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(inputJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/post-haiku")
		time.Sleep(100 * time.Millisecond)
		c.PostHaiku(ctx)
	}
}

type pair struct {
	subscriber_session_id models.InlineObject3
	receiver_user_id      int64
}

func subscribeUserForTest(pair_list []pair) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	for _, pair := range pair_list {
		inputJson, _ := json.Marshal(pair.subscriber_session_id)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(inputJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/subscribe/:user_id")
		ctx.SetParamNames("user_id")
		ctx.SetParamValues(strconv.Itoa(int(pair.receiver_user_id)))
		c.PostSubscribe(ctx)
	}
}
func deleteSubscribeUserForTest(pair_list []pair) {
	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	for _, pair := range pair_list {
		inputJson, _ := json.Marshal(pair.subscriber_session_id)
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(string(inputJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)
		ctx.SetPath("/api/subscribe/:user_id")
		ctx.SetParamNames("user_id")
		ctx.SetParamValues(strconv.Itoa(int(pair.receiver_user_id)))
		c.DeleteSubscribe(ctx)
	}
}
