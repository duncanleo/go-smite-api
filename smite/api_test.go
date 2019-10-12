package smite

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func authClient() Client {
	godotenv.Load("../.env")
	return Client{DevID: os.Getenv("SMITE_DEV_ID"), AuthKey: os.Getenv("SMITE_AUTH_KEY"), Debug: true}
}

func TestPing(t *testing.T) {
	c := Client{}
	_, err := c.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestSession(t *testing.T) {
	c := authClient()
	createSessionResponse, err := c.CreateSession()
	if err != nil {
		t.Error(err)
	}
	_, err = c.TestSession(createSessionResponse.SessionID)
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetDataUsed(createSessionResponse.SessionID)
	if err != nil {
		t.Error(err)
	}
}

func TestGods(t *testing.T) {
	c := authClient()
	createSessionResponse, err := c.CreateSession()
	if err != nil {
		t.Error(err)
	}
	_, err = c.GetGods(createSessionResponse.SessionID)
	if err != nil {
		t.Error(err)
	}
}

func TestGetPlayerStats(t *testing.T) {
	c := authClient()
	createSessionResponse, err := c.CreateSession()
	if err != nil {
		t.Error(err)
	}
	playerResults, err := c.GetPlayerIDByName(createSessionResponse.SessionID, "Lassiz")
	if err != nil {
		t.Error(err)
	}
	if len(playerResults) == 0 {
		t.Errorf("No players returned by GetPlayerIDByName")
		return
	}
	player, err := c.GetPlayer(createSessionResponse.SessionID, playerResults[0].PlayerID)
	if err != nil {
		t.Error(err)
	}
	if len(player) == 0 {
		t.Errorf("No players returned by GetPlayer")
		return
	}
	if player[0].ID != playerResults[0].PlayerID {
		t.Errorf("PlayerID does not match in GetPlayer")
	}
	achievements, err := c.GetPlayerAchievements(createSessionResponse.SessionID, playerResults[0].PlayerID)
	if err != nil {
		t.Error(err)
	}
	if achievements.PlayerID != playerResults[0].PlayerID {
		t.Errorf("PlayerID does not match in GetPlayerAchievements")
	}
	godRanks, err := c.GetPlayerGodRanks(createSessionResponse.SessionID, playerResults[0].PlayerID)
	if len(godRanks) == 0 {
		t.Errorf("No god ranks returned by GetPlayerGodRanks")
		return
	}
	if godRanks[0].PlayerID != strconv.Itoa(playerResults[0].PlayerID) { // I have no idea why playerID is a string in this route
		t.Errorf("PlayerID does not match in GetPlayerGodRanks")
	}
	matchHistory, err := c.GetPlayerMatchHistory(createSessionResponse.SessionID, playerResults[0].PlayerID)
	if len(godRanks) == 0 {
		t.Errorf("No god ranks returned by GetPlayerMatchHistory")
		return
	}
	if matchHistory[0].PlayerID != playerResults[0].PlayerID { // I have no idea why playerID is a string in this route
		t.Errorf("PlayerID does not match in GetPlayerMatchHistory")
	}
}
