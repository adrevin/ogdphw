//go:build integration_tests

package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	internalhttp "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

const eventsApi = "http://localhost:5000/api/events/"

func TestServerIsOnline(t *testing.T) {
	t.Run("server is online", func(t *testing.T) {
		resp, err := client.Get("http://localhost:5000/hello") //nolint:noctx
		CheckResponse(t, resp, err)
	})
}

var userID *uuid.UUID
var eventRequest *internalhttp.EventRequest
var eventID *uuid.UUID

func TestAddEvent(t *testing.T) {
	t.Run("add event", func(t *testing.T) {
		userID, err := uuid.NewUUID()
		require.NoError(t, err)
		eventRequest = &internalhttp.EventRequest{
			Title:    "2024-01-01 event",
			Time:     time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration: 3600,
			UserID:   userID,
		}
		eventRequestJSON, err := json.Marshal(eventRequest)
		require.NoError(t, err)
		resp, err := client.Post( //nolint:noctx
			eventsApi, "application/json", bytes.NewBuffer(eventRequestJSON))
		CheckResponse(t, resp, err)

		content, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		eventIDResponse := internalhttp.EventID{}
		err = json.Unmarshal(content, &eventIDResponse)
		require.NoError(t, err)
		eventID = &eventIDResponse.ID
		resp.Body.Close()
	})
}

func TestGetEvents(t *testing.T) {
	var eventID uuid.UUID
	t.Run("day events", func(t *testing.T) {
		resp, err := client.Get(eventsApi + "day/2024-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		var responseEvents = GetResponseEvents(t, resp)
		resp.Body.Close()
		require.Equal(t, 1, len(responseEvents))
		responseEvent := responseEvents[0]
		require.Equal(t, eventRequest.Title, responseEvent.Title)
		require.Equal(t, eventRequest.Time, responseEvent.Time)
		require.Equal(t, time.Duration(eventRequest.Duration)*time.Second, responseEvent.Duration)
		require.Equal(t, eventRequest.UserID, responseEvent.UserID)
		eventID = responseEvent.ID
	})

	t.Run("week events", func(t *testing.T) {
		resp, err := client.Get(eventsApi + "week/2024-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		var responseEvents = GetResponseEvents(t, resp)
		require.Equal(t, 1, len(responseEvents))
		require.Equal(t, eventID, responseEvents[0].ID)
		resp.Body.Close()
	})

	t.Run("month events", func(t *testing.T) {
		resp, err := client.Get(eventsApi + "month/2024-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		var responseEvents = GetResponseEvents(t, resp)
		require.Equal(t, 1, len(responseEvents))
		require.Equal(t, eventID, responseEvents[0].ID)
		resp.Body.Close()
	})
}
func TestUpdateEvent(t *testing.T) {
	t.Run("update event", func(t *testing.T) {
		newUserID, err := uuid.NewUUID()
		require.NoError(t, err)

		eventRequest.Title = "New Title"
		eventRequest.Time = time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
		eventRequest.Duration = 1800
		eventRequest.UserID = newUserID

		eventRequestJSON, err := json.Marshal(eventRequest)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, eventsApi+eventID.String(), bytes.NewBuffer(eventRequestJSON))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		CheckResponse(t, resp, err)
		resp.Body.Close()

		resp, err = client.Get(eventsApi + "day/2025-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		var responseEvents = GetResponseEvents(t, resp)
		resp.Body.Close()
		require.Equal(t, 1, len(responseEvents))
		responseEvent := responseEvents[0]
		require.Equal(t, eventRequest.Title, responseEvent.Title)
		require.Equal(t, eventRequest.Time, responseEvent.Time)
		require.Equal(t, time.Duration(eventRequest.Duration)*time.Second, responseEvent.Duration)
		require.Equal(t, eventRequest.UserID, responseEvent.UserID)
	})
}

func TestDeleteEvent(t *testing.T) {
	t.Run("delete event", func(t *testing.T) {
		DeleteEvent(t, eventID.String())
		resp, err := client.Get(eventsApi + "day/2025-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		var responseEvents = GetResponseEvents(t, resp)
		resp.Body.Close()
		require.Equal(t, 0, len(responseEvents))
	})
}

var lastTestEvents []internalhttp.EventResponse

func TestDayWeekMonthLogic(t *testing.T) {
	t.Run("post test events", func(t *testing.T) {
		userID, err := uuid.NewUUID()
		require.NoError(t, err)
		PostEvent(t, NewTestEvent(userID, 1))  // first week, mo
		PostEvent(t, NewTestEvent(userID, 2))  // first week, tu
		PostEvent(t, NewTestEvent(userID, 29)) // last week, mo
		PostEvent(t, NewTestEvent(userID, 30)) // last week, tu
		PostEvent(t, NewTestEvent(userID, 31)) // last week, we

		resp, err := client.Get(eventsApi + "day/2024-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		var responseEvents = GetResponseEvents(t, resp)
		require.Equal(t, 1, len(responseEvents))
		resp.Body.Close()

		resp, err = client.Get(eventsApi + "week/2024-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		responseEvents = GetResponseEvents(t, resp)
		require.Equal(t, 2, len(responseEvents))
		resp.Body.Close()

		resp, err = client.Get(eventsApi + "week/2024-01-29") //nolint:noctx
		CheckResponse(t, resp, err)
		responseEvents = GetResponseEvents(t, resp)
		require.Equal(t, 3, len(responseEvents))
		resp.Body.Close()

		resp, err = client.Get(eventsApi + "month/2024-01-01") //nolint:noctx
		CheckResponse(t, resp, err)
		responseEvents = GetResponseEvents(t, resp)
		require.Equal(t, 5, len(responseEvents))
		resp.Body.Close()
		lastTestEvents = responseEvents
	})
}
func TestCleanTestResults(t *testing.T) {
	for _, event := range lastTestEvents {
		DeleteEvent(t, event.ID.String())
	}
}
func DeleteEvent(t *testing.T, eventID string) { //nolint:thelper
	req, err := http.NewRequest(http.MethodDelete, eventsApi+eventID, nil)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	CheckResponse(t, resp, err)
	resp.Body.Close()
}

func PostEvent(t *testing.T, eventRequest internalhttp.EventRequest) { //nolint:thelper
	eventRequestJSON, err := json.Marshal(eventRequest)
	require.NoError(t, err)
	resp, err := client.Post(eventsApi, "application/json", bytes.NewBuffer(eventRequestJSON)) //nolint:noctx
	CheckResponse(t, resp, err)
	resp.Body.Close()
}

func NewTestEvent(userID uuid.UUID, day int) internalhttp.EventRequest {
	return internalhttp.EventRequest{
		Title:    strconv.Itoa(day),
		Time:     time.Date(2024, time.January, day, 0, 0, 0, 0, time.UTC),
		Duration: 3600,
		UserID:   userID,
	}
}

func CheckResponse(t *testing.T, resp *http.Response, err error) { //nolint:thelper
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
}

func GetResponseEvents(t *testing.T, resp *http.Response) []internalhttp.EventResponse { //nolint:thelper
	content, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	events := make([]internalhttp.EventResponse, 0)
	err = json.Unmarshal(content, &events)
	require.NoError(t, err)
	return events
}
