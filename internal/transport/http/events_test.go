package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caioverine/sport-pulse/internal/domain"
)

func TestEventsHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/events", nil)
	recorder := httptest.NewRecorder()

	NewRouter().ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.StatusCode)
	}

	if contentType := response.Header.Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", contentType)
	}

	var events []domain.Event

	if err := json.NewDecoder(response.Body).Decode(&events); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(events) == 0 {
		t.Fatal("expected at least one event, got none")
	}

	if events[0].MatchID != "match-001" {
		t.Errorf("expected first event MatchID %q, got %q", "match-001", events[0].MatchID)
	}
}
