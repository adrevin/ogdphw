package internalhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/app"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	memorystorage "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

// testing with in memory storage

const ContentType = "application/json"

func TestServer(t *testing.T) {
	logger := logger.New(zap.NewDevelopmentConfig())
	defer logger.Sync()

	t.Run("basic", func(t *testing.T) {
		storage := memorystorage.New()
		calendar := app.New(logger, storage)
		server := NewServer(logger, calendar, configuration.ServerConfiguration{})

		testServer := httptest.NewServer(http.HandlerFunc(server.handleCalendarRequest))
		defer testServer.Close()

		response, err := http.Get(testServer.URL + "/api/events/day/2020-01-01") //nolint:noctx
		require.NoError(t, err)
		defer response.Body.Close()
		require.Equal(t, http.StatusOK, response.StatusCode)
		require.Equal(t, ContentType, response.Header.Get("Content-Type"))

		out, err := io.ReadAll(response.Body)
		require.NoError(t, err)
		require.Equal(t, "[]", string(out))
	})

	t.Run("complex", func(t *testing.T) {
		storage := memorystorage.New()
		calendar := app.New(logger, storage)
		server := NewServer(logger, calendar, configuration.ServerConfiguration{})

		testServer := httptest.NewServer(http.HandlerFunc(server.handleCalendarRequest))
		defer testServer.Close()
		userID, err := uuid.NewUUID()
		require.NoError(t, err)

		// create
		eventRequest := EventRequest{
			Title:    "Title",
			Time:     time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			Duration: 3600,
			UserID:   userID,
		}
		request, err := json.Marshal(eventRequest)
		require.NoError(t, err)

		response, err := http.Post(testServer.URL+"/api/events", ContentType, bytes.NewReader(request)) //nolint:noctx
		require.NoError(t, err)
		defer response.Body.Close()
		require.Equal(t, http.StatusOK, response.StatusCode)
		require.Equal(t, "application/json", response.Header.Get("Content-Type"))

		out, err := io.ReadAll(response.Body)

		eID := &EventID{}
		json.Unmarshal(out, eID)
		require.NoError(t, err)
		require.NotEqual(t, uuid.Nil, eID.ID)

		response, err = http.Get( //nolint:noctx
			testServer.URL + "/api/events/day/" + eventRequest.Time.Format(time.DateOnly))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
		require.Equal(t, ContentType, response.Header.Get("Content-Type"))

		out, err = io.ReadAll(response.Body)
		response.Body.Close()
		events := make([]EventResponse, 0)
		json.Unmarshal(out, &events)
		require.NoError(t, err)

		require.Equal(t, eventRequest.Title, events[0].Title)
		require.Equal(t, eventRequest.Time, events[0].Time)
		require.Equal(t, eventRequest.UserID, events[0].UserID)
		require.Equal(t, time.Duration(eventRequest.Duration)*time.Second, events[0].Duration)

		// update
		updateRequest := EventRequest{
			Title:    "NewTitle",
			Time:     time.Date(7019, time.December, 1, 0, 0, 0, 0, time.UTC),
			Duration: 3300,
			UserID:   userID,
		}

		request, err = json.Marshal(updateRequest)
		require.NoError(t, err)

		req, err := http.NewRequest( //nolint:noctx
			http.MethodPut, testServer.URL+"/api/events/"+eID.ID.String(), bytes.NewReader(request))
		require.NoError(t, err)
		req.Header.Set("Content-Type", ContentType)
		client := &http.Client{}
		response, err = client.Do(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
		response.Body.Close()

		response, err = http.Get( //nolint:noctx
			testServer.URL + "/api/events/day/" + updateRequest.Time.Format(time.DateOnly))

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
		require.Equal(t, ContentType, response.Header.Get("Content-Type"))

		out, err = io.ReadAll(response.Body)
		response.Body.Close()
		events = make([]EventResponse, 0)
		json.Unmarshal(out, &events)
		require.NoError(t, err)

		require.Equal(t, updateRequest.Title, events[0].Title)
		require.Equal(t, updateRequest.Time, events[0].Time)
		require.Equal(t, updateRequest.UserID, events[0].UserID)
		require.Equal(t, time.Duration(updateRequest.Duration)*time.Second, events[0].Duration)

		// delete
		req, err = http.NewRequest( //nolint:noctx
			http.MethodDelete, testServer.URL+"/api/events/"+eID.ID.String(), bytes.NewReader(request))
		require.NoError(t, err)
		req.Header.Set("Content-Type", ContentType)
		client = &http.Client{}
		response, err = client.Do(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
		response.Body.Close()

		response, err = http.Get( //nolint:noctx
			testServer.URL + "/api/events/day/" + updateRequest.Time.Format(time.DateOnly))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, response.StatusCode)
		require.Equal(t, ContentType, response.Header.Get("Content-Type"))

		out, err = io.ReadAll(response.Body)
		response.Body.Close()
		events = make([]EventResponse, 0)
		json.Unmarshal(out, &events)
		require.NoError(t, err)
		require.Equal(t, 0, len(events))
	})
}
