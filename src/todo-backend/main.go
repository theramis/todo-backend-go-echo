package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/todos", getAllTodosHandler)
	e.POST("/todos", createTodoHandler)

	e.Logger.Fatal(e.Start(":" + port))
}
