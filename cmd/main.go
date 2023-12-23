package main

import (
	"github.com/bhanna1693/gameoflife/handlers/hello"
	"github.com/bhanna1693/gameoflife/handlers/home"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "assets")
	e.GET("/", func(e echo.Context) error {
		return home.HandleHome(e)
	})
	e.GET("/hello", func(e echo.Context) error {
		return hello.HandleHello(e)
	})
	e.Logger.Fatal(e.Start(":80"))
}
