package hello

import (
	"github.com/bhanna1693/gameoflife/components/hello"
	"github.com/bhanna1693/gameoflife/services"
	"github.com/bhanna1693/gameoflife/utils"
	"github.com/labstack/echo/v4"
)

func HandleHello(e echo.Context) error {
	name := utils.Ternary(e.QueryParam("name") == "", "user", e.QueryParam("name")).(string)
	return services.Render(e, hello.Hello(name))
}
