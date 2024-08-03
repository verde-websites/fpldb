package models

type Team struct {
	ID        int64  `json:"id" bun:"id,pk,autoincrement"`
	TeamName  string `json:"team_name" bun:"team_name"`
	ShortName string `json:"short_name" bun:"short_name"`
}
