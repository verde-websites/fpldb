package generate

import (
	"fmt"
	"os"
	"time"

	"github.com/verde-websites/fpldb/dto"
	"gopkg.in/yaml.v3"
)

type FixtureModel struct {
	Model string       `yaml:"model"`
	Rows  []FixtureRow `yaml:"rows"`
}

type FixtureRow struct {
	FixtureID            string    `yaml:"_id"`
	ID                   int64     `yaml:"id"`
	SeasonID             int64     `yaml:"season_id"`
	GameweekID           int64     `yaml:"game_week_id"`
	HomeTeamID           int64     `yaml:"home_team_id"`
	AwayTeamID           int64     `yaml:"away_team_id"`
	KickoffTime          time.Time `yaml:"kickoff_time"`
	Minutes              int64     `yaml:"minutes"`
	Finished             bool      `yaml:"finished"`
	FinishedProvisional  bool      `yaml:"finished_provisional"`
	ProvisionalStartTime bool      `yaml:"provisional_start_time"`
	Started              bool      `yaml:"started"`
	HomeTeamScore        int64     `yaml:"home_team_score"`
	AwayTeamScore        int64     `yaml:"away_team_score"`
}

func GenerateFixtures(fplFixturesResp []dto.FixtureFPLResponseDto, teamData []TeamRow) error {
	teamToTrackerID := make(map[int64]int64)
	for _, team := range teamData {
		teamToTrackerID[team.FplTrackerID] = team.ID
	}
	fixtureData := []FixtureModel{}
	for _, fixture := range fplFixturesResp {
		fixtureID := fmt.Sprintf("fixture%d", fixture.ID)
		fixtureModel := FixtureModel{
			Model: "Fixture",
			Rows:  []FixtureRow{},
		}
		kickoffTime, err := time.Parse("2006-01-02T15:04:05Z", fixture.KickoffTime)
		if err != nil {
			return err
		}

		fixtureModel.Rows = append(fixtureModel.Rows, FixtureRow{
			FixtureID:            fixtureID,
			ID:                   int64(fixture.ID),
			SeasonID:             1,
			GameweekID:           int64(fixture.Event),
			HomeTeamID:           int64(teamToTrackerID[int64(fixture.TeamH)]),
			AwayTeamID:           int64(teamToTrackerID[int64(fixture.TeamA)]),
			KickoffTime:          kickoffTime,
			Minutes:              int64(fixture.Minutes),
			Finished:             fixture.Finished,
			FinishedProvisional:  fixture.FinishedProvisional,
			ProvisionalStartTime: fixture.ProvisionalStartTime,
			Started:              fixture.Started,
			HomeTeamScore:        int64(fixture.TeamHScore),
			AwayTeamScore:        int64(fixture.TeamAScore),
		})
		fixtureData = append(fixtureData, fixtureModel)
	}

	yamlData, err := yaml.Marshal(fixtureData)
	if err != nil {
		fmt.Printf("error marshalling fixture yaml fixture: %v", err)
		return err
	}

	err = os.WriteFile("bunapp/embed/fixture/fixture.yml", yamlData, 0644)
	if err != nil {
		fmt.Printf("error writing fixture yaml fixture: %v", err)
		return err
	}
	return nil
}
