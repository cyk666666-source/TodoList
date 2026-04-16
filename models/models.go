package models

import (
	"bubble/dao"
)

//和Model有关的放这一层

// Todo 结构体，所有地方都能用
type Todo struct {
	ID     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Status bool   `json:"status" form:"status"`
}

//Todo 的增删改查

func CreatATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoList(todolist *[]Todo) (err error) {
	err = dao.DB.Find(&todolist).Error
	return
}

func GetATodo(todo *Todo, id int) (err error) {
	err = dao.DB.First(&todo, id).Error
	return
}

func UpdateATodo(todo *Todo, id int) (err error) {
	err = GetATodo(todo, id)
	if err != nil {
		return
	}
	status := todo.Status
	todo.Status = !status
	err = dao.DB.Save(&todo).Error
	return err
}

func DeleteATodo(id int) (err error) {
	err = dao.DB.Delete(&Todo{}, "id=?", id).Error
	return err
}
