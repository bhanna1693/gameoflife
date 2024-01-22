package home

import (
	"github.com/bhanna1693/gameoflife/internal/utils"
	"github.com/bhanna1693/gameoflife/web/views/home"
	"github.com/labstack/echo/v4"
)

func HandleHome(e echo.Context) error {
	var name string
	if e.QueryParam("name") == "" {
		name = "user"
	} else {
		name = e.QueryParam("name")
	}

	return utils.Render(e, home.Home(name))
}
