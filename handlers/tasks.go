package handlers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"

	m "echo_todoList/models"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, m.GetTasks(db))
	}
}


func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var task m.Task
		// Map imcoming JSON body to the new Task
		c.Bind(&task)

		log.Println(task)

		// Add a task using model
		id, err := m.PutTask(db, task.Name)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}


func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id,_:=strconv.Atoi(c.Param("id"))
		_,err:=m.DeleteTask(db,id)
		//Return a JSON response on success
		if err==nil{
			return c.JSON(http.StatusOK,H{
				"deleted":id,
			})
		} else {
			return err
		}
	}
}