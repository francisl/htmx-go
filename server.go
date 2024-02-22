package main

import (
	"htmx-go-intro/reviews"
	"htmx-go-intro/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/public", "public")

	e.GET("/reviews", func(c echo.Context) error {
		component := views.Reviews(reviews.All())
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.DELETE("/reviews/:id", func(c echo.Context) error {
		id := c.Param("id")
		err := reviews.Delete(id)

		if err != nil {
			e.Logger.Error(err)
			return c.String(http.StatusNotFound, "Not Found")
		}

		return c.HTML(http.StatusOK, "")
	})

	e.GET("/reviewform", func(c echo.Context) error {
		component := views.ReviewForm()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.POST("/reviewform", func(c echo.Context) error {
		review, _ := reviews.Add(reviews.CreateReview{
			Presentation: c.FormValue("presentation"),
			Rating:       c.FormValue("rating"),
			TheGood:      c.FormValue("good"),
			TheBad:       c.FormValue("bad"),
			Presenter:    c.FormValue("presenter"),
		})
		component := views.ReviewFormPost(review)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.File("/", "public/index.html")

	e.Logger.Debug(e.Start(":4000"))
}
