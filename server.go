package main

import (
	"htmx-go-intro/views"

	"github.com/labstack/echo/v4"
)

type State struct {
	count int
}

func main() {
	e := echo.New()
	e.Static("/public", "public")

	state := State{count: 0}
	e.GET("/ticket", func(c echo.Context) error {
		component := views.Ticket(state.count)
		state.count++
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.File("/", "public/index.html")

	e.Logger.Debug(e.Start(":4000"))
}
