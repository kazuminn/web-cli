package main

import (
    "net/http"

    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func main() {
    e := echo.New()
    e.GET("/", get)
    e.Run(standard.New(":1323"))
}

func get(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}
