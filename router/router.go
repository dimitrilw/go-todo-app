package router

import (
	"net/http"

	"github.com/dimitrilw/go-todo-app/controller"
	"github.com/gin-gonic/gin" // cspell:disable-line
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLFiles("view/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("todo", controller.GetTodoList)
		v1.POST("todo", controller.CreateTodo)
		v1.GET("todo/:id", controller.GetTodo)
		v1.PUT("todo/:id", controller.UpdateTodo)
		v1.DELETE("todo/:id", controller.DeleteTodo)
	}

	return r
}
