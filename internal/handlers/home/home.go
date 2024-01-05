package home

import (
	"github.com/bhanna1693/gameoflife/internal/utils"
	"github.com/bhanna1693/gameoflife/web/templates/home"
	"github.com/labstack/echo/v4"
)

func HandleHome(e echo.Context) error {
	return utils.Render(e, home.Home())
}
