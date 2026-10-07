package provider

import "github.com/caioverine/sport-pulse/internal/domain"

type SimulatedProvider struct {
}

func (p SimulatedProvider) Events() []domain.Event {
	return []domain.Event{
		{
			ID:      "event-001",
			Type:    "MATCH_STARTED",
			MatchID: "match-001",
			Minute:  0,
		},
		{
			ID:      "event-002",
			Type:    "GOAL",
			MatchID: "match-001",
			Minute:  23,
		},
		{
			ID:      "event-003",
			Type:    "YELLOW_CARD",
			MatchID: "match-001",
			Minute:  41,
		},
	}
}
