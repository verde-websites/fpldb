package models

import "time"

type Fixture struct {
	ID                   int64     `json:"id" bun:"id,pk,autoincrement"`
	FplTrackerID         int64     `json:"fpl_tracker_id" bun:"fpl_tracker_id"`
	GameWeekID           int64     `json:"game_week_id" bun:"game_week_id"`
	HomeTeamID           int64     `json:"home_team_id" bun:"home_team_id"`
	AwayTeamID           int64     `json:"away_team_id" bun:"away_team_id"`
	KickoffTime          time.Time `json:"kickoff_time" bun:"kickoff_time,type:timestamp with time zone"`
	Minutes              int       `json:"minutes" bun:"minutes"`
	Finished             bool      `json:"finished" bun:"finished"`
	FinishedProvisional  bool      `json:"finished_provisional" bun:"finished_provisional"`
	ProvisionalStartTime bool      `json:"provisional_start_time" bun:"provisional_start_time"`
	Started              bool      `json:"started" bun:"started"`
	HomeTeamScore        int64     `json:"home_team_score" bun:"home_team_score"`
	AwayTeamScore        int64     `json:"away_team_score" bun:"away_team_score"`
	GameWeek             *Gameweek `bun:"rel:belongs-to,join:game_week_id=id" json:"game_week,omitempty"`
	HomeTeam             *Team     `bun:"rel:belongs-to,join:home_team_id=id" json:"home_team,omitempty"`
	AwayTeam             *Team     `bun:"rel:belongs-to,join:away_team_id=id" json:"away_team,omitempty"`
}
