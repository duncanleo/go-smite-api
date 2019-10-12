package smite

// CreateSessionResponse response when creating a session
type CreateSessionResponse struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestamp"`
}
