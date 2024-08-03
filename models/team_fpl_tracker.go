package models

type TeamFplTracker struct {
	SeasonID      int64   `json:"season_id" bun:"season_id,pk"`
	TeamID        int64   `json:"team_id" bun:"team_id,pk"`
	TeamTrackerID int64   `json:"team_tracker_id" bun:"team_tracker_id,pk"`
	Season        *Season `bun:"rel:belongs-to,join:season_id=id" json:"season,omitempty"`
	Team          *Team   `bun:"rel:belongs-to,join:team_id=id" json:"team,omitempty"`
}
