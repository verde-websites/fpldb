package models

type PlayerFplTracker struct {
	SeasonID        int64   `json:"season_id" bun:"season_id,pk"`
	PlayerID        int64   `json:"player_id" bun:"player_id,pk"`                 // Our player ID
	PlayerTrackerId int64   `json:"player_tracker_id" bun:"player_tracker_id,pk"` // FPL player ID (New per season)
	Season          *Season `bun:"rel:belongs-to,join:season_id=id" json:"season,omitempty"`
	Player          *Player `bun:"rel:belongs-to,join:player_id=id" json:"player,omitempty"`
}
