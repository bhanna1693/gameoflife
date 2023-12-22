package hello

import (
	"github.com/bhanna1693/gameoflife/components/hello"
	"github.com/bhanna1693/gameoflife/services"
	"github.com/labstack/echo/v4"
)

func HandleHello(e echo.Context) error {
	name := e.QueryParam("name")
	if name == "" {
		name = "user"
	}
	return services.Render(e, hello.Hello(name))
}
