package controller

import (
	"net/http"

	"github.com/gin-gonic/gin" // cspell:disable-line

	"github.com/dimitrilw/go-todo-app/database"
	"github.com/dimitrilw/go-todo-app/model"
)

func GetTodoList(gc *gin.Context) {
	var todoList []model.Todo
	err := database.GetTodoList(&todoList)
	if err != nil {
		gc.AbortWithStatus(http.StatusNotFound)
	} else {
		gc.JSON(http.StatusOK, todoList)
	}
}

func CreateTodo(gc *gin.Context) {
	var todo model.Todo
	gc.BindJSON(&todo)
	err := database.CreateTodo(&todo)
	if err != nil {
		gc.AbortWithStatus(http.StatusNotFound)
	} else {
		gc.JSON(http.StatusOK, todo)
	}
}

func GetTodo(gc *gin.Context) {
	id := gc.Params.ByName("id")
	var todo model.Todo
	err := database.GetTodo(&todo, id)
	if err != nil {
		gc.AbortWithStatus(http.StatusNotFound)
	} else {
		gc.JSON(http.StatusOK, todo)
	}
}

func UpdateTodo(gc *gin.Context) {
	var todo model.Todo
	id := gc.Params.ByName("id")
	err := database.GetTodo(&todo, id)
	if err != nil {
		gc.JSON(http.StatusNotFound, todo)
	}
	gc.BindJSON(&todo)
	err = database.UpdateTodo(&todo, id)
	if err != nil {
		gc.AbortWithStatus(http.StatusNotFound)
	} else {
		gc.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(gc *gin.Context) {
	var todo model.Todo
	id := gc.Params.ByName("id")
	err := database.DeleteTodo(&todo, id)
	if err != nil {
		gc.AbortWithStatus(http.StatusNotFound)
	} else {
		gc.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
