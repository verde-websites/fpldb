package generate

import (
	"fmt"
	"os"

	"github.com/verde-websites/fpldb/dto"
	"gopkg.in/yaml.v3"
)

func templatePlayerYamlFixture(bootstrapStaticResp dto.BootstrapStaticFPLResponseDto) error {
	playerData := []PlayerModel{}
	playerFplTrackerData := []PlayerFplTrackerModel{}
	playerModel := PlayerModel{
		Model: "Player",
		Rows:  []PlayerRow{},
	}

	playerFplTrackerModel := PlayerFplTrackerModel{
		Model: "PlayerFplTracker",
		Rows:  []PlayerFplTrackerRow{},
	}

	for _, player := range bootstrapStaticResp.Players {
		playerID := fmt.Sprintf("player%d", player.Code)
		playerTrackerID := fmt.Sprintf("playerFplTracker%d", player.Code)
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
		playerFplTrackerModel.Rows = append(playerFplTrackerModel.Rows, PlayerFplTrackerRow{
			PlayerFplTrackerId: playerTrackerID,
			SeasonID:           1,
			PlayerID:           int64(player.Code),
			PlayerTrackerID:    int64(player.ID),
		})
	}
	playerData = append(playerData, playerModel)
	playerFplTrackerData = append(playerFplTrackerData, playerFplTrackerModel)
	playerYamlData, err := yaml.Marshal(playerData)
	if err != nil {
		fmt.Printf("error marshalling player yaml fixture: %v", err)
		return err
	}

	playerFplTackerYamlData, err := yaml.Marshal(playerFplTrackerData)
	if err != nil {
		fmt.Printf("error marshalling playerFplTracker yaml fixture: %v", err)
		return err
	}
	// create the player yaml fixture file

	err = os.WriteFile("bunapp/embed/fixture/player.yml", playerYamlData, 0644)
	if err != nil {
		fmt.Printf("error writing player yaml fixture: %v", err)
		return err
	}

	err = os.WriteFile("bunapp/embed/fixture/playerFplTracker.yml", playerFplTackerYamlData, 0644)
	if err != nil {
		fmt.Printf("error writing playerFplTracker yaml fixture: %v", err)
		return err
	}
	return nil
}
