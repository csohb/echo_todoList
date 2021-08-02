package models

import (
	"gorm.io/gorm"
	"log"
)


type Task struct {
	ID        uint   `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	User      string `json:"user"`
}

// Get Response
type TaskResponse struct {
	Tasks []Task `json:"tasks"`
}

// Update
type UpdateRequest struct {
	Completed bool `json:"completed"`
}

// Post
type PostRequest struct {
	Todo string `json:"todo"`
	User string `json:"user"`
}



func SelectData(db *gorm.DB, user string) ([]TaskModel, error) {
	task:=make([]TaskModel,0)
	if err:=db.Where("user = ?",user).Find(&task).Error; err!=nil{
		log.Printf("db load err : %s", err)
	}
	log.Printf("select data : %s", task)
	return task,nil
}

func InsertData(db *gorm.DB, req PostRequest) (error) {
	task:=TaskModel{}
	task.Completed=false
	task.Todo=req.Todo
	task.User=req.User

	if err:=db.Create(&task).Error; err!=nil{
		log.Printf("create err : %+v", err)
		return err
	}
	return nil
}

func UpdateData(db *gorm.DB, req UpdateRequest, id int) (error) {
	task:=TaskModel{}
	task.Completed=req.Completed

	if err:=db.Model(&task).Where("id = ?",id).Update("completed", req.Completed).Error; err!=nil{
		log.Printf("update task failed : %+v", err)
		return err
	}

	return nil
}

func DeleteData(db *gorm.DB, id int) (error) {
	task:=TaskModel{}
	if err:=db.Where("id = ? ", id).Delete(&task).Error; err!=nil{
		return err
	}
	return nil
}