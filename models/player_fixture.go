package models

type PlayerFixture struct {
	ID                       int64    `json:"id" bun:"id,pk,autoincrement"`
	PlayerID                 int64    `json:"player_id" bun:"player_id"`
	PlayerFplTrackerID       int64    `json:"player_fpl_tracker_id" bun:"player_fpl_tracker_id"`
	FixtureID                int64    `json:"fixture_id" bun:"fixture_id"`
	GameWeekID               int64    `json:"game_week_id" bun:"game_week_id"`
	TeamID                   int64    `json:"team_id" bun:"team_id"`
	Minutes                  int64    `json:"minutes" bun:"minutes"`
	CleanSheet               bool     `json:"clean_sheet" bun:"clean_sheet"`
	GoalsScored              int64    `json:"goals_scored" bun:"goals_scored"`
	GoalsConceded            int64    `json:"goals_conceded" bun:"goals_conceded"`
	Assists                  int64    `json:"assists" bun:"assists"`
	Saves                    int64    `json:"saves" bun:"saves"`
	OwnGoals                 int64    `json:"own_goals" bun:"own_goals"`
	PenaltiesSaved           int64    `json:"penalties_saved" bun:"penalties_saved"`
	PenaltiesMissed          int64    `json:"penalties_missed" bun:"penalties_missed"`
	YellowCards              int64    `json:"yellow_cards" bun:"yellow_cards"`
	RedCards                 int64    `json:"red_cards" bun:"red_cards"`
	BonusPoints              int64    `json:"bonus_points" bun:"bonus_points"`
	BPSPoints                int64    `json:"bps_points" bun:"bps_points"`
	Influence                float64  `json:"influence" bun:"influence"`
	Creativity               float64  `json:"creativity" bun:"creativity"`
	Threat                   float64  `json:"threat" bun:"threat"`
	ICTIndex                 float64  `json:"ict_index" bun:"ict_index"`
	Started                  bool     `json:"started" bun:"started"`
	ExpectedGoals            float64  `json:"expected_goals" bun:"expected_goals"`
	ExpectedAssists          float64  `json:"expected_assists" bun:"expected_assists"`
	ExpectedGoalInvolvements float64  `json:"expected_goal_involvements" bun:"expected_goal_involvements"`
	ExpectedGoalsConceded    float64  `json:"expected_goals_conceded" bun:"expected_goals_conceded"`
	TotalPoints              int64    `json:"total_points" bun:"total_points"`
	Player                   *Player  `bun:"rel:has-one,join:player_id=id" json:"player,omitempty"`
	Fixture                  *Fixture `bun:"rel:has-one,join:fixture_id=id" json:"fixture,omitempty"`
	Team                     *Team    `bun:"rel:has-one,join:team_id=id" json:"team,omitempty"`
}
