package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
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
	e.DELETE("/api/:haiku_id", c.DeleteApiHaikuId)

	// DeleteApiSubscribeUserId - 
	e.DELETE("/api/subscribe/:user_id", c.DeleteApiSubscribeUserId)

	// GetApiHaikuId - get_haiku
	e.GET("/api/:haiku_id", c.GetApiHaikuId)

	// GetApiTimeline - timeline
	e.GET("/api/timeline", c.GetApiTimeline)

	// GetApiUser - user_info
	e.GET("/api/users/:user_id", c.GetApiUser)

	// GetTop - top
	e.GET("/api/top", c.GetTop)

	// PostApiSignup - 
	e.POST("/api/signup", c.PostApiSignup)

	// PostSignin - 
	e.POST("/api/signin", c.PostSignin)

	// PostSubscribe - 
	e.POST("/api/subscribe/:user_id", c.PostSubscribe)

	// PostUser - 
	e.POST("/api/post_haiku", c.PostUser)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}