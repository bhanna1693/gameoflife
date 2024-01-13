package main

import (
	"database/sql"
	"fmt"
	"log"

	gameoflifedb "github.com/bhanna1693/gameoflife/internal/database/gameoflife"
	"github.com/bhanna1693/gameoflife/internal/handlers/gameoflife"
	"github.com/bhanna1693/gameoflife/internal/handlers/home"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

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
	e.Static("/static", "web/static")
	e.GET("/", func(e echo.Context) error {
		return home.HandleHome(e)
	})
	e.GET("/gameoflife", func(e echo.Context) error {
		return gameoflife.HandleGameOfLife(e)
	})
	e.POST("/gameoflife", func(e echo.Context) error {
		return gameoflife.HandleGameOfLife(e)
	})
	e.POST("/gameoflife/process-board", func(e echo.Context) error {
		return gameoflife.HandleGameOfLifeBoard(e)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
