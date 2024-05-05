package tarball

import (
	"mygo/core/models"
	"mygo/services/sqlite"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTarballs(c *gin.Context) {
	var tarballs []models.Tarball
	if err := sqlite.MyGODB.Model(&models.Tarball{}).Order("id desc").Preload("QemuTasks").Find(&tarballs).Error; err != nil {
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
		Data:    tarballs,
		Errors:  nil,
	}
	c.JSON(status, resp)
}
