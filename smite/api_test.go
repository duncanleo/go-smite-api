package smite

import (
	"os"
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

func TestGetPlayerIDByName(t *testing.T) {
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
	}
}
