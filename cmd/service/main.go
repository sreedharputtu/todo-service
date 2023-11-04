package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("cmd/service/templates/*.html")
	apiRouter := r.Group("/v1/api/todos")
	{
		apiRouter.GET("", getTodos)
		apiRouter.POST("", saveTodo)
	}
	uiRouter := r.Group("/v1/ui/todos")
	{
		uiRouter.GET("", getTodosPage)
		uiRouter.POST("", saveTodosPage)
	}

	r.Run(":8080")
}

func getTodosPage(c *gin.Context){
	log.Printf("todos:%v",todos)
	c.HTML(http.StatusOK , "index.html",gin.H{
		"Todos" : todos,
	})
}

func saveTodosPage(c *gin.Context){
	task := c.PostForm("task.name")
	todos = append(todos , Todo{
		Name : task ,
		Done : false,
	})
	log.Printf("todo list %v",todos)
	c.HTML(http.StatusOK , "index.html",gin.H{
		"Todos" : todos,
	})
}


func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func saveTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	todos = append(todos, todo)
	log.Printf("todos:%v", todos)
	c.JSON(http.StatusOK, todo)
}

type Todo struct {
	Name   string `json:"name"`
	Done 	 bool 	`json:"done"`
}

var todos []Todo
