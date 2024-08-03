package dto

type FixtureFPLResponseDto struct {
	Code                 int            `json:"code"`
	Event                int            `json:"event"`
	Finished             bool           `json:"finished"`
	FinishedProvisional  bool           `json:"finished_provisional"`
	ID                   int            `json:"id"`
	KickoffTime          string         `json:"kickoff_time"`
	Minutes              int            `json:"minutes"`
	ProvisionalStartTime bool           `json:"provisional_start_time"`
	Started              bool           `json:"started"`
	TeamA                int            `json:"team_a"`
	TeamAScore           int            `json:"team_a_score"`
	TeamH                int            `json:"team_h"`
	TeamHScore           int            `json:"team_h_score"`
	Stats                []OriginalStat `json:"stats"`
	TeamHDifficulty      int            `json:"team_h_difficulty"`
	TeamADifficulty      int            `json:"team_a_difficulty"`
	PulseID              int            `json:"pulse_id"`
}
type OriginalStat struct {
	Identifier string `json:"identifier"`
	A          []struct {
		Value   int `json:"value"`
		Element int `json:"element"`
	} `json:"a"`
	H []struct {
		Value   int `json:"value"`
		Element int `json:"element"`
	} `json:"h"`
}
