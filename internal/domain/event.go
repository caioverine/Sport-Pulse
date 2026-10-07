package domain

type Event struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	MatchID string `json:"matchId"`
	Minute  int    `json:"minute"`
}
