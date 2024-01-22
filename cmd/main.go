package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	gameoflifedb "github.com/bhanna1693/gameoflife/internal/database/gameoflife"
	"github.com/bhanna1693/gameoflife/internal/handlers/gameoflife"
	"github.com/bhanna1693/gameoflife/internal/handlers/home"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	// Open a database connection
	db, err := sql.Open("sqlite3", "your-database-file.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is established
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	err = gameoflifedb.CreateGameOfLifeTable(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Static("/static", "web/static")
	e.GET("/", func(e echo.Context) error {
		return home.HandleHome(e)
	})
	e.GET("/gameoflife", func(e echo.Context) error {
		return gameoflife.HandleGameOfLife(e, db)
	})
	e.POST("/gameoflife", func(e echo.Context) error {
		return gameoflife.HandleGameOfLife(e, db)
	})
	e.POST("/gameoflife/process-board/:id", func(e echo.Context) error {
		return gameoflife.HandleGameOfLifeBoard(e, db)
	})
	e.GET("/gameoflife/results", func(e echo.Context) error {
		return gameoflife.HandleGameOfLifeResults(e, db)
	})
	e.DELETE("/gameoflife/:id", func(e echo.Context) error {
		return gameoflife.HandleGameOfLifeDelete(e, db)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
