package models

type Player struct {
	ID         int64  `json:"id" bun:"id,pk,autoincrement"`
	FirstName  string `json:"first_name" bun:"first_name"`
	SecondName string `json:"second_name" bun:"second_name"`
	WebName    string `json:"web_name" bun:"web_name"`
	Status     string `json:"status" bun:"status"`
	PlayerType string `json:"player_type" bun:"player_type"`
	TeamID     int64  `json:"team_id" bun:"team_id"` // Added to match SQL schema and added relation
	Team       *Team  `bun:"rel:belongs-to,join:team_id=id" json:"team,omitempty"`
}
