package home

import (
	"github.com/bhanna1693/gameoflife/internal/utils"
	"github.com/bhanna1693/gameoflife/web/views/home"
	"github.com/labstack/echo/v4"
)

func HandleHome(e echo.Context) error {
	name := utils.Ternary(e.QueryParam("name") == "", "user", e.QueryParam("name")).(string)
	return utils.Render(e, home.Home(name))
}
