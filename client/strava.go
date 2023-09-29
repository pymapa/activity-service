package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	strava "github.com/strava/go.strava"
)

func GetActivities(accessToken string) []*strava.ActivitySummary {

	fmt.Println("Access token: " + accessToken)

	// TODO: mockdata. Call strava API instead
	activitiesFile, err := os.Open("mockData/getActivitiesResponse.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer activitiesFile.Close()
	byteResult, _ := ioutil.ReadAll(activitiesFile)
	var activities []*strava.ActivitySummary
	json.Unmarshal(byteResult, &activities)

	return activities
}
