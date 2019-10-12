package smite

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	baseURL = "http://api.smitegame.com/smiteapi.svc"
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

// GetAuthedSubRoute make a HTTP GET request authenticated with developer's signature
func (c Client) GetAuthedSubRoute(route, subRoute string) (*http.Response, error) {
	return c.Get(fmt.Sprintf("%s/%s/%s/%s/%s", route, c.DevID, c.signature(strings.Replace(route, "json", "", 1)), subRoute, c.timestamp()))
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
	resp, err := c.GetAuthedSubRoute("testsessionjson", sessionID)

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
	resp, err := c.GetAuthedSubRoute("getdatausedjson", sessionID)

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
