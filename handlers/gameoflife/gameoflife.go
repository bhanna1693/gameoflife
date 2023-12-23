package gameoflife

import (
	"github.com/bhanna1693/gameoflife/components/gameoflife"
	"github.com/bhanna1693/gameoflife/services"
	"github.com/labstack/echo/v4"
)

func HandleGameOfLife(e echo.Context) error {
	return services.Render(e, gameoflife.GameOfLife())
}
