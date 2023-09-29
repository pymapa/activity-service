package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pymapa/activity-service/client"
)

// GetActivities returns a list of strava activities
func GetActivities(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("strava_access_token")

	if accessToken == "" {
		fmt.Println("No access token provided")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No access token provided"})
		return
	}

	fmt.Println("Access token: " + accessToken)

	activities := client.GetActivities(accessToken)

	if activities == nil {
		ctx.IndentedJSON(http.StatusNoContent, gin.H{"message": "No activities found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, activities)
}
