package handlers

import (
	models "Todo_List/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type H map[string]interface{}


func GetTasks(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := models.TaskResponse{Tasks: make([]models.Task, 0)}
		user:=c.QueryParam("user")
		log.Printf("user : %s", user)
		tasks:=make([]models.TaskModel,0)

		if v,err:=models.SelectData(db, user); err==nil{
			tasks=v
		} else {
			log.Println(err)
		}

		for _, v := range tasks {
			resp.Tasks = append(resp.Tasks, models.Task{
				ID:        v.ID,
				Todo:      v.Todo,
				Completed: v.Completed,
				User:      v.User,
			})
		}
		log.Printf("resp in handler : %+v", resp)
		return c.JSON(http.StatusOK, resp)
	}
}


func PostTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		req:=models.PostRequest{}
		// request body 받아옴
		if err := c.Bind(&req); err != nil {
			log.Printf("bind err : %+v", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}
		log.Printf("req : %v", req)
		if err:=models.InsertData(db, req); err!=nil{
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, H{
			"created": req,
		})
	}
}

// completed false -> true || true -> false
func PutTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		req:=models.UpdateRequest{}

		id, _ := strconv.Atoi(c.Param("id"))
		log.Printf("id : %d", id)
		// get request body
		if err := c.Bind(&req); err != nil {
			log.Printf("bind error : %s", err)
			return c.JSON(http.StatusInternalServerError, nil)
		}

		log.Printf("completed : %v", req.Completed)

		if err:=models.UpdateData(db,req,id); err!=nil{
			log.Printf("update data err : %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, H{
			"updated": req.Completed,
		})
	}
}

func DeleteTask(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := strconv.Atoi(c.Param("id"))
		log.Printf("id : %d", id)

		if err:=models.DeleteData(db, id); err!=nil{
			log.Printf("delete data err : %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, H{
			"deleted": id,
		})
	}
}
