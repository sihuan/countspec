package task

import (
	"mygo/core/models"
	"mygo/services/sqlite"

	"github.com/gin-gonic/gin"
)

// r.GET("/task/stdout/:id", task.GetTaskStdout)

func GetTaskStdout(c *gin.Context) {
	id := c.Param("id")
	var task models.QemuTask
	if err := sqlite.MyGODB.Model(&models.QemuTask{}).Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	stdout := task.Stdout
	basePath := task.Path
	// 读取文件并返回
	c.File(basePath + "/" + stdout)

}

// r.GET("/task/stdout/:id", task.GetTaskStdout)

func GetTaskStderr(c *gin.Context) {
	id := c.Param("id")
	var task models.QemuTask
	if err := sqlite.MyGODB.Model(&models.QemuTask{}).Where("id = ?", id).First(&task).Error; err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	stdout := task.Stderr
	basePath := task.Path
	// 读取文件并返回
	c.File(basePath + "/" + stdout)

}
