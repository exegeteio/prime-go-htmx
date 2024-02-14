package main

import (
	server "github.com/exegeteio/prime-go-htmx"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Count struct {
	Count int
}

func main() {
	e := server.NewEchoServer()
	count := Count{Count: 0}

	e.GET("/", func(c echo.Context) error {
		count.Count = count.Count + 1
		return c.Render(http.StatusOK, "index", count)
	})

	server.Start(e)
}
