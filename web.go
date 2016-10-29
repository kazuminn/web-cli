package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	//flag 処理
	defaultPort := ":1323"

	var (
		defaultport = flag.String("defaultport", defaultPort, ":port")
	)
	port := *defaultport
	flag.Parse()

	//web server
	e := echo.New()
	e.GET("/", get)
	e.Run(standard.New(port))
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
