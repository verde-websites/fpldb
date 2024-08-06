package generate

import (
	"fmt"
	"os"

	"github.com/verde-websites/fpldb/dto"
	"gopkg.in/yaml.v3"
)

func templatePlayerYamlFixture(bootstrapStaticResp dto.BootstrapStaticFPLResponseDto) error {
	playerData := []PlayerModel{}
	playerFplSeasonData := []PlayerFplSeasonModel{}
	playerModel := PlayerModel{
		Model: "Player",
		Rows:  []PlayerRow{},
	}

	playerFplSeasonModel := PlayerFplSeasonModel{
		Model: "PlayerFplSeason",
		Rows:  []PlayerFplSeasonRow{},
	}

	for _, player := range bootstrapStaticResp.Players {
		playerID := fmt.Sprintf("player%d", player.Code)
		playerFplSeasonID := fmt.Sprintf("player%d-season%d", player.Code, 1)
		playerModel.Rows = append(playerModel.Rows, PlayerRow{
			PlayerID:     playerID,
			FplTrackerID: int64(player.ID),
			ID:           int64(player.Code),
			FirstName:    player.FirstName,
			SecondName:   player.SecondName,
			WebName:      player.WebName,
			PlayerType:   int64(player.PlayerType),
			Status:       player.Status,
			TeamId:       int64(player.TeamCode),
		})
		playerFplSeasonModel.Rows = append(playerFplSeasonModel.Rows, PlayerFplSeasonRow{
			PlayerFplSeasonID: playerFplSeasonID,
			SeasonID:          1,
			PlayerID:          int64(player.Code),
		})
	}
	playerData = append(playerData, playerModel)
	playerFplSeasonData = append(playerFplSeasonData, playerFplSeasonModel)
	playerYamlData, err := yaml.Marshal(playerData)
	if err != nil {
		fmt.Printf("error marshalling player yaml fixture: %v", err)
		return err
	}

	playerFplSeasonYamlData, err := yaml.Marshal(playerFplSeasonData)
	if err != nil {
		fmt.Printf("error marshalling playerFplSeason yaml fixture: %v", err)
		return err
	}
	// create the player yaml fixture file

	err = os.WriteFile("bunapp/embed/fixture/player.yml", playerYamlData, 0644)
	if err != nil {
		fmt.Printf("error writing player yaml fixture: %v", err)
		return err
	}

	err = os.WriteFile("bunapp/embed/fixture/playerFplSeason.yml", playerFplSeasonYamlData, 0644)
	if err != nil {
		fmt.Printf("error writing playerFplSeason yaml fixture: %v", err)
		return err
	}
	return nil
}
