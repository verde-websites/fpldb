package generate

/*
- The purpose of this is to pull data from the FPL bootstrap static endpoint
  and generate relevant fixtures to store in the repo.
- The fixtures should be stored in the bunapp/embed/fixtures directory
*/

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/verde-websites/fpldb/dto"
)

func GenerateDatabaseFixtures() error {

	client := &http.Client{}
	resp, err := client.Get("https://fantasy.premierleague.com/api/bootstrap-static/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var bootstrapStaticResp dto.BootstrapStaticFPLResponseDto
	if err := json.Unmarshal(body, &bootstrapStaticResp); err != nil {
		return err
	}

	fixtureResp, err := client.Get("https://fantasy.premierleague.com/api/fixtures/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fixtureBody, err := io.ReadAll(fixtureResp.Body)
	if err != nil {
		return err
	}

	var fixturesResp []dto.FixtureFPLResponseDto
	if err := json.Unmarshal(fixtureBody, &fixturesResp); err != nil {
		return err
	}

	// template a gameweek yaml fixture entry in the embed fixture dir from the bootstrap static response
	if err := templateGameweekYamlFixture(bootstrapStaticResp); err != nil {
		return err
	}
	// template a team yaml fixture entry in the embed fixture dir from the bootstrap static response
	teamFplTrackerRows, err := templateTeamYamlFixture(bootstrapStaticResp)
	if err != nil {
		return err
	}

	// template a player yaml fixture entry in the embed fixture dir from the bootstrap static response
	if err := templatePlayerYamlFixture(bootstrapStaticResp); err != nil {
		return err
	}

	if err := GenerateFixtures(fixturesResp, teamFplTrackerRows); err != nil {
		return err
	}
	return nil
}

type TeamModel struct {
	Model string    `yaml:"model"`
	Rows  []TeamRow `yaml:"rows"`
}
type TeamFplTrackerModel struct {
	Model string              `yaml:"model"`
	Rows  []TeamFplTrackerRow `yaml:"rows"`
}

type PlayerModel struct {
	Model string      `yaml:"model"`
	Rows  []PlayerRow `yaml:"rows"`
}

type PlayerFplTrackerModel struct {
	Model string                `yaml:"model"`
	Rows  []PlayerFplTrackerRow `yaml:"rows"`
}

type GameweekModel struct {
	Model string        `yaml:"model"`
	Rows  []GameweekRow `yaml:"rows"`
}

type GameweekRow struct {
	FixtureID        string    `yaml:"_id"`
	ID               int64     `yaml:"id"`
	GameweekNumber   int64     `yaml:"gameweek_number"`
	Name             string    `yaml:"name"`
	TransferDeadline time.Time `yaml:"transfer_deadline"`
	GameweekActive   bool      `yaml:"gameweek_active"`
	GameweekFinished bool      `yaml:"gameweek_finished"`
	DataChecked      bool      `yaml:"data_checked"`
}

type TeamRow struct {
	TeamID    string `yaml:"_id"`
	ID        int64  `yaml:"id"`
	TeamName  string `yaml:"team_name"`
	ShortName string `yaml:"short_name"`
}

type TeamFplTrackerRow struct {
	TeamFplTrackerId string `yaml:"_id"`
	SeasonID         int64  `yaml:"season_id"`
	TeamID           int64  `yaml:"team_id"`
	TeamTrackerID    int64  `yaml:"team_tracker_id"`
}

type PlayerRow struct {
	PlayerID   string `yaml:"_id"`
	ID         int64  `yaml:"id"`
	FirstName  string `yaml:"first_name"`
	SecondName string `yaml:"second_name"`
	WebName    string `yaml:"web_name"`
	PlayerType int64  `yaml:"player_type"`
	Status     string `yaml:"status"`
	TeamId     int64  `yaml:"team_id"`
}

type PlayerFplTrackerRow struct {
	PlayerFplTrackerId string `yaml:"_id"`
	SeasonID           int64  `yaml:"season_id"`
	PlayerID           int64  `yaml:"player_id"`
	PlayerTrackerID    int64  `yaml:"player_tracker_id"`
}
