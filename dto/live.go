package dto

type LiveFPLResponseDto struct {
	Players []*LivePlayerFPLResponseDto `json:"elements"`
}
type LivePlayerFPLResponseDto struct {
	PlayerTrackerId int                        `json:"id"`
	GameweekStats   LiveGameweekGameweekStats  `json:"stats"`
	FixtureStats    []LiveGameweekFixtureStats `json:"explain"`
}

type LiveGameweekGameweekStats struct {
	Minutes                  int    `json:"minutes"`
	GoalsScored              int    `json:"goals_scored"`
	Assists                  int    `json:"assists"`
	CleanSheets              int    `json:"clean_sheets"`
	GoalsConceded            int    `json:"goals_conceded"`
	OwnGoals                 int    `json:"own_goals"`
	PenaltiesSaved           int    `json:"penalties_saved"`
	PenaltiesMissed          int    `json:"penalties_missed"`
	YellowCards              int    `json:"yellow_cards"`
	RedCards                 int    `json:"red_cards"`
	Saves                    int    `json:"saves"`
	Bonus                    int    `json:"bonus"`
	BPS                      int    `json:"bps"`
	Influence                string `json:"influence"`
	Creativity               string `json:"creativity"`
	Threat                   string `json:"threat"`
	ICTIndex                 string `json:"ict_index"`
	Starts                   int    `json:"starts"`
	ExpectedGoals            string `json:"expected_goals"`
	ExpectedAssists          string `json:"expected_assists"`
	ExpectedGoalInvolvements string `json:"expected_goal_involvements"`
	ExpectedGoalsConceded    string `json:"expected_goals_conceded"`
	TotalPoints              int    `json:"total_points"`
}

type LiveGameweekFixtureStats struct {
	FixtureID int                       `json:"fixture"`
	Stats     []LiveGameweekFixtureStat `json:"stats"`
}

type LiveGameweekFixtureStat struct {
	Identifier string `json:"identifier"`
	Points     int    `json:"points"`
	Value      int    `json:"value"`
}
