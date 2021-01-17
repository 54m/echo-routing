package main

import (
	"github.com/54m/echo-routing/example/router"
	"github.com/54m/echo-routing/output"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// NOTE: Get
	{
		e.GET("/", router.Index)

		e.GET("/check", router.Check)
	}

	// NOTE: Post
	{
		e.POST("/ping", router.Ping)
	}

	output.Do(e.Routes())
}
