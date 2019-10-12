package smite

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL = "http://api.smitegame.com/smiteapi.svc"

	// LanguageCodeEnglish API language code for English
	LanguageCodeEnglish = "1"
)

// Client represents an API client for SMITE API
type Client struct {
	BaseURL string
	DevID   string
	AuthKey string
	Debug   bool
}

func (c Client) baseURL() string {
	if len(c.BaseURL) > 0 {
		return c.BaseURL
	}
	return baseURL
}

// Get make a HTTP GET request to a route.
// Used by built-in helper functions but exported for future-proofing.
func (c Client) Get(route string) (*http.Response, error) {
	var url = fmt.Sprintf("%s/%s", c.baseURL(), route)
	if c.Debug {
		log.Printf("HTTP GET -> %s\n", url)
	}
	return http.Get(url)
}

// GetAuthed make a HTTP GET request authenticated with developer's signature
func (c Client) GetAuthed(route string) (*http.Response, error) {
	return c.Get(fmt.Sprintf("%s/%s/%s/%s", route, c.DevID, c.signature(strings.Replace(route, "json", "", 1)), c.timestamp()))
}

// GetAuthedSecondaryRoute make a HTTP GET request authenticated with developer's signature
func (c Client) GetAuthedSecondaryRoute(route, secondaryRoute string) (*http.Response, error) {
	return c.Get(fmt.Sprintf("%s/%s/%s/%s/%s", route, c.DevID, c.signature(strings.Replace(route, "json", "", 1)), secondaryRoute, c.timestamp()))
}

// GetAuthedTertiaryRoute make a HTTP GET request authenticated with developer's signature
func (c Client) GetAuthedTertiaryRoute(route, secondaryRoute, tertiaryRoute string) (*http.Response, error) {
	return c.Get(fmt.Sprintf("%s/%s/%s/%s/%s/%s", route, c.DevID, c.signature(strings.Replace(route, "json", "", 1)), secondaryRoute, c.timestamp(), tertiaryRoute))
}

func (c Client) timestamp() string {
	return time.Now().UTC().Format("20060102150405")
}

func (c Client) signature(route string) string {
	md5Hash := md5.New()
	io.WriteString(md5Hash, c.DevID)
	io.WriteString(md5Hash, route)
	io.WriteString(md5Hash, c.AuthKey)
	io.WriteString(md5Hash, c.timestamp())
	return fmt.Sprintf("%x", md5Hash.Sum(nil))
}

// Ping A quick way of validating access to the Hi-Rez API.
func (c Client) Ping() (string, error) {
	resp, err := c.Get("pingjson")

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d", resp.StatusCode)
	}
	var result string
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	return result, err
}

// CreateSession create a session
func (c Client) CreateSession() (CreateSessionResponse, error) {
	var result CreateSessionResponse
	resp, err := c.GetAuthed("createsessionjson")

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	// API doesn't use HTTP status codes :(
	if result.RetMsg != "Approved" {
		return result, fmt.Errorf("Session not created: %s", result.RetMsg)
	}
	return result, err
}

// TestSession test if the current developer signature is valid
func (c Client) TestSession(sessionID string) (string, error) {
	var result string
	resp, err := c.GetAuthedSecondaryRoute("testsessionjson", sessionID)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	// API doesn't use HTTP status codes :(
	if !strings.Contains(result, "This was a successful test") {
		return result, fmt.Errorf("Session invalid: %s", result)
	}
	return result, err
}

// GetDataUsed get current API usage limits
func (c Client) GetDataUsed(sessionID string) (GetDataUsedResponse, error) {
	var result GetDataUsedResponse
	resp, err := c.GetAuthedSecondaryRoute("getdatausedjson", sessionID)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	// API doesn't use HTTP status codes :(
	if len(result) > 0 && result[0].RetMsg == "Invalid session id." {
		return result, fmt.Errorf("Invalid session ID")
	}
	return result, err
}

// GetGods get details on all gods in SMITE
func (c Client) GetGods(sessionID string) (GetGodsResponse, error) {
	var result GetGodsResponse
	resp, err := c.GetAuthedTertiaryRoute("getgodsjson", sessionID, LanguageCodeEnglish)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	// API doesn't use HTTP status codes :(
	if len(result) > 0 && result[0].RetMsg == "Invalid session id." {
		return result, fmt.Errorf("Invalid session ID")
	}
	return result, err
}

// GetPlayerIDByName Search for a player by name
func (c Client) GetPlayerIDByName(sessionID, playerName string) (GetPlayerIDByNameResponse, error) {
	var result GetPlayerIDByNameResponse
	resp, err := c.GetAuthedTertiaryRoute("getplayeridbynamejson", sessionID, playerName)

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	// API doesn't use HTTP status codes :(
	if len(result) > 0 && result[0].RetMsg == "Invalid session id." {
		return result, fmt.Errorf("Invalid session ID")
	}
	return result, err
}

// GetPlayer Get data about a player
func (c Client) GetPlayer(sessionID string, playerID int) (GetPlayerResponse, error) {
	var result GetPlayerResponse
	resp, err := c.GetAuthedTertiaryRoute("getplayerjson", sessionID, strconv.Itoa(playerID))

	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)

	// API doesn't use HTTP status codes :(
	if len(result) > 0 && result[0].RetMsg == "Invalid session id." {
		return result, fmt.Errorf("Invalid session ID")
	}
	return result, err
}
