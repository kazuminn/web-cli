package main

import (
		"flag"
		"fmt"
		"net/http"
		"os"

		"github.com/labstack/echo"
		"github.com/labstack/echo/engine/standard"
)

func main() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Fprintf(os.Stderr, "Error:\n%s", err)
				os.Exit(1)
			}
		}()
		os.Exit(_main())
}

func _main() int {
		// help
		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
			flag.PrintDefaults()
		}

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

		return 0
}


func get(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
}
