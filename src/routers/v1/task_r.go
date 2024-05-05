package v1

import (
	"mygo/controllers/v1/task"

	"github.com/gin-gonic/gin"
)

func SetTaskRouter(r *gin.RouterGroup) {
	r.POST("/task", task.CreateTask)
	r.GET("/task", task.GetTasks)
	r.GET("/task/stdout/:id", task.GetTaskStdout)
	r.GET("/task/stderr/:id", task.GetTaskStderr)
	// r.DELETE("/task/:id", task.StopTask)
}
