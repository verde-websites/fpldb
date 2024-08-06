package models

type Team struct {
	ID           int64  `json:"id" bun:"id,pk,autoincrement"`
	FplTrackerID int64  `json:"fpl_tracker_id" bun:"fpl_tracker_id"`
	TeamName     string `json:"team_name" bun:"team_name"`
	ShortName    string `json:"short_name" bun:"short_name"`
}
