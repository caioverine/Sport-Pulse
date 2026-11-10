package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func eventsStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")

	controller := http.NewResponseController(w)

	if err := controller.Flush(); err != nil {
		return
	}

	simulatedProvider := provider.SimulatedProvider{}
	events := simulatedProvider.Events()
	next := 0

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return

		case <-ticker.C:
			if next >= len(events) {
				if _, err := fmt.Fprint(w, ": heartbeat\n\n"); err != nil {
					return
				}
			} else {
				event := events[next]

				payload, err := json.Marshal(event)
				if err != nil {
					return
				}

				_, err = fmt.Fprintf(
					w,
					"id: %s\nevent: football-event\ndata: %s\n\n",
					event.ID,
					payload,
				)
				if err != nil {
					return
				}

				next++
			}

			if err := controller.Flush(); err != nil {
				return
			}
		}
	}
}
