package tarball

import (
	"mygo/core/models"
	"mygo/services/sqlite"

	"github.com/gin-gonic/gin"
)

// r.GET("/tarball/config/:id", tarball.GetTarballConfig)

func GetTarballConfig(c *gin.Context) {
	id := c.Param("id")
	var tarball models.Tarball
	if err := sqlite.MyGODB.Model(&models.Tarball{}).Where("id = ?", id).First(&tarball).Error; err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	config := tarball.ConfigPath
	// 读取文件并返回
	c.File(config)
}
