package main

import (
	"fmt"
	"mygo/core/models"
	v1 "mygo/routers/v1"
	"mygo/services/sqlite"

	"github.com/gin-contrib/static"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	sqlite.CheckDB()
	// sqlite.FirstRun()
	// mysql.CheckDB()
	router = gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./data/dist/", false)))
	// router.NoRoute(noRouteHandler())
	version1 := router.Group("/api/v1")
	v1.InitRoutes(version1)
}

func main() {
	fmt.Println("Server Running on Port: ", 9090)
	router.Run("127.0.0.1:9090")
}

func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var statuscode int
		var message string = "Not Found"
		var data interface{} = nil
		var listError []models.ErrorModel = nil
		// var endpoint string = c.Request.URL.String()
		// var method string = c.Request.Method

		var tempEr models.ErrorModel
		tempEr.ErrorCode = 4041
		// tempEr.Hints = "Not Found !! \n Routes In Valid. You enter on invalid Page/Endpoint"
		// tempEr.Info = "visit http://localhost:9090/v1/docs to see the available routes"
		listError = append(listError, tempEr)
		statuscode = 404
		responseModel := &models.ResponseModel{
			Status:  statuscode,
			Message: message,
			Data:    data,
			Errors:  listError,
			// Endpoint: endpoint,
			// Method:   method,
		}
		var content gin.H = responseModel.NewResponse()
		c.JSON(statuscode, content)
	}
}
