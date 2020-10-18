package API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) ListMeetingParticipants(meetingId int) (apiResponse ListMeetingParticipantsResponse, err error) {

	endpoint := fmt.Sprintf("/articles/%d/participants", Email)
	httpMethod := http.MethodGet

	var b []byte
	b, err = client.executeRequest(endpoint, httpMethod)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &apiResponse)
	if err != nil {
		return
	}

	return
}
