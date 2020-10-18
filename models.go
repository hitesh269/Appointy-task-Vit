package models

import "time"

type Meeting struct {
	ID           bson.ObjectID  `bson:"_id" json:"id"`
	Title        string         `bson:"title" json:"title"`
	Created      time.Time      `bson:"created" json:"created"`
	StartTime    time.Time      `bson:"start_time" json:"start_time"`
	EndTime      time.Time      `bson:"end_time" json:"end_time"`
	Participants []Participants `bson:"participants" json:"participants"`
}
type Participant struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	RSVP  string `bson:"rsvp" json:"rsvp"`
}
