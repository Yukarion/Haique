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

func genUUIDForTest() (string, error) {
	return "test", nil
}

var namepw = models.InlineObject{Name: "tatara", Pw: "fuga"}

func TestSignup(t *testing.T) {
	tests := []struct {
		input    models.InlineObject
		code     int
		expected models.SessionId
	}{
		{
			input:    models.InlineObject{Name: "hoge", Pw: "fuga"},
			code:     http.StatusCreated,
			expected: models.SessionId{SessionId: "test"},
		},
	}

	e := echo.New()
	c, _ := NewContainerForTest(genUUIDForTest)
	for _, test := range tests {
		inputJson, _ := json.Marshal(test.input)
		req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(string(inputJson)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		if assert.NoError(t, c.PostSignup(ctx)) {
			var actual models.SessionId
			json.Unmarshal(rec.Body.Bytes(), &actual)
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, test.expected.SessionId, actual.SessionId)
		}
	}

}
