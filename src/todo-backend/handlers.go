package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var todoRepository = NewInMemoryTodoRepository()

func getAllTodosHandler(c echo.Context) error {
	todos := todoRepository.GetAll()
	return c.JSON(http.StatusOK, todos)
}

func createTodoHandler(c echo.Context) (err error) {
	todo := new(Todo)
	if err = c.Bind(todo); err != nil {
		return err
	}

	todoRepository.Create(todo)
	todo.SetUrl(c)
	return c.JSON(http.StatusCreated, todo)
}
