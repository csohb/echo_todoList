package models

import (
	"database/sql"
	"log"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// gorm 사용해서 get, post, put, delete method 만들기
func GetTasks(db *sql.DB) TaskCollection {
	sql:="SELECT * FROM tasks"
	rows,err:=db.Query(sql)
	if err !=nil{
		panic(err)
	}
	defer rows.Close()

	result:=TaskCollection{}
	for rows.Next(){
		task:=Task{}
		err2:=rows.Scan(&task.ID,&task.Name)
		if err2!=nil{
			panic(err2)
		}
		result.Tasks=append(result.Tasks,task)
	}
	log.Println(result)
	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(name)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}


func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql:="DELETE FROM tasks WHERE id=?"

	stmt,err:=db.Prepare(sql)
	if err!=nil{
		panic(err)
	}

	//Replace ? with id
	result,err2:=stmt.Exec(id)
	if err2!=nil{
		panic(err2)
	}

	// RowsAffected returns the number of rows affected by an update, insert, or delete.
	return result.RowsAffected()
}