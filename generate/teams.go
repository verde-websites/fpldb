package generate

import (
	"fmt"
	"os"

	"github.com/verde-websites/fpldb/dto"
	"gopkg.in/yaml.v3"
)

func templateTeamYamlFixture(bootstrapStaticResp dto.BootstrapStaticFPLResponseDto) ([]TeamRow, error) {
	teamData := []TeamModel{}
	teamFplSeasonData := []TeamFplSeasonModel{}
	teamModel := TeamModel{
		Model: "Team",
		Rows:  []TeamRow{},
	}

	TeamFplSeasonModel := TeamFplSeasonModel{
		Model: "TeamFplSeason",
		Rows:  []TeamFplSeasonRow{},
	}

	for _, team := range bootstrapStaticResp.Teams {
		teamID := fmt.Sprintf("team%d", team.Code)
		TeamFplSeasonID := fmt.Sprintf("team%d-season%d", team.Code, 1)
		teamModel.Rows = append(teamModel.Rows, TeamRow{
			TeamID:       teamID,
			ID:           int64(team.Code),
			FplTrackerID: int64(team.ID),
			TeamName:     team.Name,
			ShortName:    team.ShortName,
		})

		TeamFplSeasonModel.Rows = append(TeamFplSeasonModel.Rows, TeamFplSeasonRow{
			TeamFplSeasonID: TeamFplSeasonID,
			SeasonID:        1,
			TeamID:          int64(team.Code),
		})
	}
	teamData = append(teamData, teamModel)
	teamFplSeasonData = append(teamFplSeasonData, TeamFplSeasonModel)
	yamlTeamData, err := yaml.Marshal(teamData)
	if err != nil {
		fmt.Printf("error marshalling team yaml fixture: %v", err)
		return nil, err
	}

	teamFplSeasonYamlData, err := yaml.Marshal(teamFplSeasonData)
	if err != nil {
		fmt.Printf("error marshalling teamFplSeason yaml fixture: %v", err)
		return nil, err
	}
	// create the team yaml fixture file

	err = os.WriteFile("bunapp/embed/fixture/team.yml", yamlTeamData, 0644)
	if err != nil {
		fmt.Printf("error writing team yaml fixture: %v", err)
		return nil, err
	}

	err = os.WriteFile("bunapp/embed/fixture/teamFplSeason.yml", teamFplSeasonYamlData, 0644)
	if err != nil {
		fmt.Printf("error writing teamFplSeason yaml fixture: %v", err)
		return nil, err
	}
	return teamModel.Rows, nil
}
