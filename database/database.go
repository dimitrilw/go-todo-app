package database

import (
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dimitrilw/go-todo-app/model"
)

var DB *gorm.DB

func InitDB() {
	slog.Info("InitDB called")
	var err error
	DB, err = gorm.Open(sqlite.Open("database/todo.sqlite"), &gorm.Config{})
	if err != nil {
		slog.Error("failed to open DB", "err", err)
	}

	slog.Info("migrating DB Todo model")
	DB.AutoMigrate(&model.Todo{})
}

func GetTodoList(todo *[]model.Todo) error {
	slog.Info("GetTodoList called")
	return DB.Find(todo).Error
}

func CreateTodo(todo *model.Todo) error {
	slog.Info("CreateTodo", "todo", todo)
	return DB.Create(todo).Error
}

func GetTodo(todo *model.Todo, id string) error {
	slog.Info("GetTodo", "id", id)
	return DB.Where("id = ?", id).First(todo).Error
}

func UpdateTodo(todo *model.Todo, id string) error {
	slog.Info("UpdateTodo", "id", id, "todo", todo)
	DB.Save(todo)
	// need implement actual error handling in this function
	return nil
}

func DeleteTodo(todo *model.Todo, id string) error {
	slog.Info("DeleteTodo", "id", id, "todo", todo)
	DB.Where("id = ?", id).Delete(todo)
	// need implement actual error handling in this function
	return nil
}

func AddTestData() {
	task := model.Todo{Title: "Task 1", Description: "Description 1"}
	CreateTodo(&task)
	task = model.Todo{Title: "Task 2", Description: "Description 2"}
	CreateTodo(&task)
}
