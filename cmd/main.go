package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api/v1/posts", getPosts)

	e.Logger.Fatal(e.Start(":3000"))
}

func getPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"Post 1", "Post 2", "Post 3"})
}
