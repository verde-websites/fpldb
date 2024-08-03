package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Gameweek struct {
	bun.BaseModel    `bun:"table:game_weeks,alias:u"`
	ID               int64     `json:"id" bun:"id,pk,autoincrement"`
	GameweekNumber   int64     `json:"gameweek_number" bun:"gameweek_number"`
	Name             string    `json:"name"`
	SeasonID         int64     `json:"season_id" bun:"season_id"`
	TransferDeadline time.Time `json:"transfer_deadline" bun:"transfer_deadline"`
	GameweekActive   bool      `json:"gameweek_active" bun:"gameweek_active"`
	GameweekFinished bool      `json:"gameweek_finished" bun:"gameweek_finished"`
	DataChecked      bool      `json:"data_checked" bun:"data_checked"`
}
