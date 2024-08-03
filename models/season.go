package models

type Season struct {
	ID         int64       `json:"id" bun:"id,pk,autoincrement"`
	SeasonName string      `json:"season_name"`
	Gameweeks  []*Gameweek `bun:"rel:has-many,join:id=season_id" json:"gameweeks,omitempty"`
}
