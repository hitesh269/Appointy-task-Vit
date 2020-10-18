package main

import (
	"fmt"
	"github.com/yash837/Appointy-Task"
	"log"
	"os"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/yash837/Appointy-Task/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func main() {

	createMeetingExample()
	listMeetingExample()
	getMeetingByIdExample()
	listMeetingParticipantssExample()
}
	
func listMeetingExample() {

	
	apiClient := API.NewClient(os.Getenv("API_URL"), os.Getenv("AUTH_TOKEN"))

	userId := os.Getenv("USER_ID")
	var resp API.ListMeetingsAPIResponse
	var err error

	resp, err = apiClient.ListMeetings(userId)
	if err != nil {
		log.Fatal(err)
	}

	for _, meeting := range resp.Meetings {
		fmt.Printf("id = %d, topic = %s, join url = %s, start time = %s\n", meeting.Id, meeting.Topic, meeting.JoinUrl, meeting.StartTime)
	}

}

func createMeetingExample() {

	
	apiClient := API.NewClient(os.Getenv("API_URL"),
		os.Getenv("AUTH_TOKEN"))

	userId := os.Getenv("USER_ID")

	var resp API.CreateMeetingResponse
	var err error

	resp, err = apiClient.CreateMeeting(userId,
		"Contributors Meeting for Project",
		meeting.MeetingTypeScheduled,
		"2020-05-31T22:00:00Z",
		30,
		"",
		"Asia/Singapore",
		"pass8888", 
		"Discuss next steps and ways to contribute for this project.",
		nil,
		nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created meeting : id = %d, topic = %s, join url = %s, start time = %s\n", resp.Id,
		resp.Topic, resp.JoinUrl, resp.StartTime)

}


func getMeetingByIdExample() {

	meetingId := 87853332664

	
	apiClient := API.NewClient(os.Getenv("API_URL"), os.Getenv("AUTH_TOKEN"))

	var err error

	var resp API.GetMeetingResponse
	resp, err = apiClient.GetMeeting(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved meeting : id = %d, topic = %s, join url = %s, start url = %s, start time = %s, status = %s\n", resp.Id,
		resp.Topic, resp.JoinUrl, resp.StartUrl, resp.StartTime, resp.Status)

}




func listMeetingParticipantsExample() {

	
	apiClient := API.NewClient(os.Getenv("API_URL"), os.Getenv("AUTH_TOKEN"))

	meetingId := 87853332664

	var resp API.ListMeetingRegistrantsResponse
	var err error

	resp, err = apiClient.ListMeetingParticipants(meetingId)
	if err != nil {
		log.Fatal(err)
	}

	for _, registrant := range resp.Registrants {
		fmt.Printf("registrant id = %s, email = %s, first name = %s, last name = %s\n",
			participants.Id, participants.Email, participants.FirstName, registrant.LastName)
	}

}
