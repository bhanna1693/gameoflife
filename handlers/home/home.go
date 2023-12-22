package home

import (
	"github.com/bhanna1693/gameoflife/components/home"
	"github.com/bhanna1693/gameoflife/services"
	"github.com/labstack/echo/v4"
)

func HandleHome(e echo.Context) error {
	return services.Render(e, home.Home())
}
