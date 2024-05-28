package handler

import (
	"net/http"

	"go-mongo-todos/pkg/in/port"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	getTodoUseCase     port.GetTodoUseCase
	getTodoListUseCase port.GetTodoListUseCase
	createTodoUseCase  port.CreateTodoUseCase
	deleteTodoUseCase  port.DeleteTodoUseCase
}

func (ctrl *TodoController) getTodo(ctx *gin.Context) {
	request := port.GetTodoRequest{
		ID: ctx.Param("todo_id"),
	}
	response, err := ctrl.getTodoUseCase.GetTodo(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response.Todo})
}

func (ctrl *TodoController) getTodoList(ctx *gin.Context) {
	response, err := ctrl.getTodoListUseCase.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response.TodoList})
}

func (ctrl *TodoController) createTodo(ctx *gin.Context) {

}

func (ctrl *TodoController) deleteTodo(ctx *gin.Context) {
}

func (ctrl *TodoController) RegisterRoute(r *gin.RouterGroup) {
	todoRouter := r.Group("/todos")
	todoRouter.GET("/:todo_id", ctrl.getTodo)
	todoRouter.GET("/", ctrl.getTodoList)
	todoRouter.POST("/", ctrl.createTodo)
	todoRouter.DELETE("/:todo_id", ctrl.deleteTodo)
}
