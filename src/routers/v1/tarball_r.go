package v1

import (
	"mygo/controllers/v1/tarball"

	"github.com/gin-gonic/gin"
)

func SetTarballRouter(r *gin.RouterGroup) {
	r.GET("/tarball", tarball.GetTarballs)
	r.POST("/tarball", tarball.CreateTarball)
	r.GET("/tarball/config/:id", tarball.GetTarballConfig)
}
