package http

import (
	"encoding/json"
	"net/http"

	"github.com/caioverine/sport-pulse/internal/provider"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	simulatedProvider := provider.SimulatedProvider{}
	events := simulatedProvider.Events()

	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "failed to encode events", http.StatusInternalServerError)
		return
	}
}
