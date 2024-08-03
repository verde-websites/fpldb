package generate

import (
	"fmt"
	"os"
	"time"

	"github.com/verde-websites/fpldb/dto"
	"gopkg.in/yaml.v3"
)

func templateGameweekYamlFixture(bootstrapStaticResp dto.BootstrapStaticFPLResponseDto) error {
	data := []GameweekModel{}
	gameweekModel := GameweekModel{
		Model: "Gameweek",
		Rows:  []GameweekRow{},
	}
	gwNumber := 1
	for _, gameweek := range bootstrapStaticResp.Gameweeks {
		transferDeadlineTime, err := time.Parse("2006-01-02T15:04:05Z", gameweek.DeadlineTime)
		if err != nil {
			return err
		}
		fixtureID := fmt.Sprintf("gw%d", gwNumber)
		gameweekModel.Rows = append(gameweekModel.Rows, GameweekRow{
			FixtureID:        fixtureID,
			ID:               int64(gwNumber),
			GameweekNumber:   int64(gwNumber),
			Name:             gameweek.Name,
			TransferDeadline: transferDeadlineTime,
			GameweekActive:   gameweek.IsCurrent,
			GameweekFinished: gameweek.Finished,
			DataChecked:      gameweek.DataChecked,
		})
		gwNumber++
	}
	data = append(data, gameweekModel)
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		fmt.Printf("error marshalling gameweek yaml fixture: %v", err)
		return err
	}

	err = os.WriteFile("bunapp/embed/fixture/gameweek.yml", yamlData, 0644)
	if err != nil {
		fmt.Printf("error writing gameweek yaml fixture: %v", err)
		return err
	}
	return nil
}
