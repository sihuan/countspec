package task

import (
	"mygo/core/models"
	"mygo/services/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var qemuTasks []models.QemuTask

	if err := sqlite.MyGODB.Model(&models.QemuTask{}).Order("tarball_id desc").Order("name").Find(&qemuTasks).Error; err != nil {
		status := http.StatusInternalServerError
		resp := models.ResponseModel{
			Status:  status,
			Message: "failed",
			Data:    nil,
			Errors:  []models.ErrorModel{{Message: err.Error(), ErrorCode: 10086}},
		}
		c.JSON(status, resp)
		return
	}
	status := http.StatusOK
	resp := models.ResponseModel{
		Status:  status,
		Message: "success",
		Data:    qemuTasks,
		Errors:  nil,
	}
	c.JSON(status, resp)
}
