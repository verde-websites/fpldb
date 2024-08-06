package generate

import (
	"fmt"
	"os"

	"github.com/verde-websites/fpldb/dto"
	"gopkg.in/yaml.v3"
)

func templateTeamYamlFixture(bootstrapStaticResp dto.BootstrapStaticFPLResponseDto) ([]TeamFplTrackerRow, error) {
	teamData := []TeamModel{}
	teamFplTrackerData := []TeamFplTrackerModel{}
	teamModel := TeamModel{
		Model: "Team",
		Rows:  []TeamRow{},
	}

	TeamFplTrackerModel := TeamFplTrackerModel{
		Model: "TeamFplTracker",
		Rows:  []TeamFplTrackerRow{},
	}

	for _, team := range bootstrapStaticResp.Teams {
		teamID := fmt.Sprintf("team%d", team.Code)
		teamModel.Rows = append(teamModel.Rows, TeamRow{
			TeamID:       teamID,
			ID:           int64(team.Code),
			FplTrackerID: int64(team.ID),
			TeamName:     team.Name,
			ShortName:    team.ShortName,
		})

		TeamFplTrackerModel.Rows = append(TeamFplTrackerModel.Rows, TeamFplTrackerRow{
			TeamFplTrackerId: teamID,
			SeasonID:         1,
			TeamID:           int64(team.Code),
			TeamTrackerID:    int64(team.ID),
		})
	}
	teamData = append(teamData, teamModel)
	teamFplTrackerData = append(teamFplTrackerData, TeamFplTrackerModel)
	yamlTeamData, err := yaml.Marshal(teamData)
	if err != nil {
		fmt.Printf("error marshalling team yaml fixture: %v", err)
		return nil, err
	}

	teamFplTrackerYamlData, err := yaml.Marshal(teamFplTrackerData)
	if err != nil {
		fmt.Printf("error marshalling teamFplTracker yaml fixture: %v", err)
		return nil, err
	}
	// create the team yaml fixture file

	err = os.WriteFile("bunapp/embed/fixture/team.yml", yamlTeamData, 0644)
	if err != nil {
		fmt.Printf("error writing team yaml fixture: %v", err)
		return nil, err
	}

	err = os.WriteFile("bunapp/embed/fixture/teamFplTracker.yml", teamFplTrackerYamlData, 0644)
	if err != nil {
		fmt.Printf("error writing teamFplTracker yaml fixture: %v", err)
		return nil, err
	}
	return TeamFplTrackerModel.Rows, nil
}
