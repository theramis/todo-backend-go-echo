package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func deleteAllTodosHandler(c echo.Context) (err error) {
	todoRepository.DeleteAll()
	return c.NoContent(http.StatusNoContent)
}

func getTodoHandler(c echo.Context) (err error) {
	if id, err := getTodoId(c); err != nil {
		return err
	} else {
		if todo, err := todoRepository.Get(id); err != nil {
			return c.String(http.StatusNotFound, "Todo note was not found")
		} else {
			return c.JSON(http.StatusOK, todo)
		}
	}
}

func deleteTodoHandler(c echo.Context) (err error) {
	if id, err := getTodoId(c); err != nil {
		return err
	} else {
		if err := todoRepository.Delete(id); err != nil {
			return c.String(http.StatusNotFound, "Todo note was not found")
		} else {
			return c.NoContent(http.StatusNoContent)
		}
	}
}

func updateTodoHandler(c echo.Context) (err error) {
	if id, err := getTodoId(c); err != nil {
		return err
	} else {
		todo, err := todoRepository.Get(id)
		if err != nil {
			return c.String(http.StatusNotFound, "Todo note was not found")
		}
		if err = c.Bind(todo); err != nil {
			return err
		}

		if err := todoRepository.Update(todo); err != nil {
			return c.String(http.StatusNotFound, "Todo note was not found")
		} else {
			return c.NoContent(http.StatusNoContent)
		}
	}
}

func getTodoId(c echo.Context) (id int, err error) {
	rawId := c.Param("id")
	id, err = strconv.Atoi(rawId)
	return
}
