package main

import (
	h "Todo_List/handlers"
	"Todo_List/migration"
	_ "database/sql"
	"github.com/labstack/echo/v4"
)



/*type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name,data)
}
*/
func main()  {

	// db Mirgration
	db:=migration.Migration()


	e:=echo.New()
	e.File("/","public/index.html")
	e.GET("/tasks",h.GetTasks(db))
	e.POST("/save",h.PostTask(db))
	e.PUT("/tasks/:id", h.PutTask(db))
	e.DELETE("/tasks/:id",h.DeleteTask(db))

	e.Logger.Fatal(e.Start(":8789"))
}

