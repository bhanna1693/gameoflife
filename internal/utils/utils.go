package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Ternary(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func Render(e echo.Context, component templ.Component) error {
	return component.Render(e.Request().Context(), e.Response())
}
