package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pymapa/activity-service/api"
)

func main() {

	router := gin.Default()
	router.GET("/activities", api.GetActivities)

	router.Run(":8080")
}
