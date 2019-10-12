package smite

// CreateSessionResponse response when creating a session
type CreateSessionResponse struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}

// GetDataUsedResponse response when getting API usage count
type GetDataUsedResponse []struct {
	ActiveSessions     int         `json:"Active_Sessions"`
	ConcurrentSessions int         `json:"Concurrent_Sessions"`
	RequestLimitDaily  int         `json:"Request_Limit_Daily"`
	SessionCap         int         `json:"Session_Cap"`
	SessionTimeLimit   int         `json:"Session_Time_Limit"`
	TotalRequestsToday int         `json:"Total_Requests_Today"`
	TotalSessionsToday int         `json:"Total_Sessions_Today"`
	RetMsg             interface{} `json:"ret_msg"`
}

// GetGodsResponse response when getting all gods
type GetGodsResponse []God

// GetPlayerIDByNameResponse response when getting player ID by name
type GetPlayerIDByNameResponse []struct {
	PlayerID int         `json:"player_id"`
	Portal   string      `json:"portal"`
	PortalID string      `json:"portal_id"`
	RetMsg   interface{} `json:"ret_msg"`
}
