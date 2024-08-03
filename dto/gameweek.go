package dto

type BSGameweekFPLResponseDto struct {
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	DeadlineTime           string `json:"deadline_time"`
	ReleaseTime            string `json:"release_time"`
	AverageEntryScore      int    `json:"average_entry_score"`
	Finished               bool   `json:"finished"`
	DataChecked            bool   `json:"data_checked"`
	HighestScoringEntry    int    `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int64  `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int    `json:"deadline_time_game_offset"`
	HighestScore           int    `json:"highest_score"`
	IsPrevious             bool   `json:"is_previous"`
	IsCurrent              bool   `json:"is_current"`
	IsNext                 bool   `json:"is_next"`
	CupLeaguesCreated      bool   `json:"cup_leagues_created"`
	H2HKoMatchesCreated    bool   `json:"h2h_ko_matches_created"`
	RankedCount            int64  `json:"ranked_count"`
	ChipPlays              []struct {
		ChipName  string `json:"chip_name"`
		NumPlayed int    `json:"num_played"`
	} `json:"chip_plays"`
	MostSelected      int `json:"most_selected"`
	MostTransferredIn int `json:"most_transferred_in"`
	TopPlayer         int `json:"top_element"`
	TopPlayerInfo     struct {
		ID     int `json:"id"`
		Points int `json:"points"`
	} `json:"top_element_info"`
	TransfersMade     int `json:"transfers_made"`
	MostCaptained     int `json:"most_captained"`
	MostViceCaptained int `json:"most_vice_captained"`
}
