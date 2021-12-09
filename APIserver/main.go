package main

import (
	"github.com/Mackyson/Haique/APIserver/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// DeleteApiHaikuId -
	e.DELETE("/api/:haiku_id", c.DeleteHaiku)

	// DeleteApiSubscribeUserId -
	e.DELETE("/api/subscribe/:user_id", c.DeleteSubscribe)

	// GetApiHaikuId - get_haiku
	e.GET("/api/haiku/:haiku_id", c.GetHaiku)

	// GetApiTimeline - timeline
	e.GET("/api/timeline", c.GetTimeline)

	// GetApiUser - user_info
	e.GET("/api/users/:user_id", c.GetUser)

	// GetTop - top
	e.GET("/api/top", c.GetTop)

	// PostApiSignup -
	e.POST("/api/signup", c.PostSignup)

	// PostHaiku -
	e.POST("/api/post-haiku", c.PostHaiku)

	// PostSignin -
	e.POST("/api/signin", c.PostSignin)

	// PostSubscribe -
	e.POST("/api/subscribe/:user_id", c.PostSubscribe)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
