package hello

import (
	"github.com/bhanna1693/gameoflife/internal/utils"
	"github.com/bhanna1693/gameoflife/web/templates/hello"
	"github.com/labstack/echo/v4"
)

func HandleHello(e echo.Context) error {
	name := utils.Ternary(e.QueryParam("name") == "", "user", e.QueryParam("name")).(string)
	return utils.Render(e, hello.Hello(name))
}