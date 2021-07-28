package main

import (
	"database/sql"
	h "echo_todoList/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

/*type Tasks struct {
	ID int `json:"id"`
	Name string `json:"name"`
}*/

func main()  {
	db,err:=sql.Open("mysql","test2:Cjswo.123@tcp(im.plea.kr:13306)/test2")
	if err!=nil{
		panic(err)
	}


	e:=echo.New()
	e.File("/","public/index.html")
	e.GET("/tasks",h.GetTasks(db))
	e.PUT("/save",h.PutTask(db))
	e.DELETE("/tasks/:id",h.DeleteTask(db))

	e.Logger.Fatal(e.Start(":8777"))
}

