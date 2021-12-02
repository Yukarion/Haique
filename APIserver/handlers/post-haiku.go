package handlers

import (
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/models"
	"github.com/labstack/echo/v4"
)

// PostHaiku -
func (c *Container) postHaiku(ctx echo.Context) error {
	var payload models.ApiPostHaikuContent
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &payload); err != nil {
		return err
	}

	return ctx.HTML(http.StatusOK, ""+payload.First)

	//return ctx.NoContent(http.StatusOK)
}
