package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	strava "github.com/strava/go.strava"
)

func main() {
	accessToken := "as3gi5wh3p548gh04598hg0493"

	if accessToken == "" {
		fmt.Println("No access token provided")
		return
	}

	router := gin.Default()
	router.GET("/activities", getActivities)

	router.Run(":8080")
}

func getActivities(c *gin.Context) {
	activitiesFile, err := os.Open("mockData/getActivitiesResponse.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer activitiesFile.Close()

	byteResult, _ := ioutil.ReadAll(activitiesFile)

	var activities []*strava.ActivitySummary

	json.Unmarshal(byteResult, &activities)

	c.IndentedJSON(http.StatusOK, activities)
}
