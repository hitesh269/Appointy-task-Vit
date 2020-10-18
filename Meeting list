package API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (client Client) ListMeetings(userId string) (apiResponse ListMeetingsAPIResponse, err error) {

	endpoint := fmt.Sprintf("/meetings?start=%d &end=%", StartTime, EndTime)
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
