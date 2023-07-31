//go:build integration_tests

package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestServerIsOnline(t *testing.T) {
	t.Run("ServerIsOnline", func(t *testing.T) {
		resp, err := client.Get("http://localhost:5000/hello") //nolint:noctx
		defer resp.Body.Close()                                //nolint:govet, staticcheck
		CheckResponse(t, resp, err)
	})
}

func TestGetEvents(t *testing.T) {
	t.Run("day events", func(t *testing.T) {
		resp, err := client.Get("http://localhost:5000/api/events/day/2024-01-01") //nolint:noctx
		defer resp.Body.Close()                                                    //nolint:govet,staticcheck
		CheckResponse(t, resp, err)
		GetResponseEvents(t, resp)
	})

	t.Run("week events", func(t *testing.T) {
		resp, err := client.Get("http://localhost:5000/api/events/week/2024-01-01") //nolint:noctx
		defer resp.Body.Close()                                                     //nolint:govet,staticcheck
		CheckResponse(t, resp, err)
		GetResponseEvents(t, resp)
	})

	t.Run("month events", func(t *testing.T) {
		resp, err := client.Get("http://localhost:5000/api/events/month/2024-01-01") //nolint:noctx
		defer resp.Body.Close()                                                      //nolint:govet,staticcheck
		CheckResponse(t, resp, err)
		GetResponseEvents(t, resp)
	})
}

func TestPostEvents(t *testing.T) {
	t.Run("Post test events", func(t *testing.T) {
		userID, err := uuid.NewUUID()
		require.NoError(t, err)
		PostEvent(t, NewTestEvent(userID, 1))  // first week, mo
		PostEvent(t, NewTestEvent(userID, 2))  // first week, tu
		PostEvent(t, NewTestEvent(userID, 10)) // second week, we
		PostEvent(t, NewTestEvent(userID, 11)) // second week, th
	})
}

func TestLogic(t *testing.T) {
	t.Run("Post test events", func(t *testing.T) {
	})
}

func PostEvent(t *testing.T, eventRequest internalhttp.EventRequest) { //nolint:thelper
	t.Run(fmt.Sprintf("post event %s", eventRequest.Time), func(t *testing.T) {
		eventRequestJSON, err := json.Marshal(eventRequest)
		require.NoError(t, err)
		resp, err := client.Post( //nolint:noctx
			"http://localhost:5000/api/events/", "application/json", bytes.NewBuffer(eventRequestJSON))
		defer resp.Body.Close() //nolint:govet,staticcheck
		CheckResponse(t, resp, err)
	})
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
