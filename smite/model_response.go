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

// GetPlayerResponse response when getting data about a player
type GetPlayerResponse []struct {
	ActivePlayerID             int         `json:"ActivePlayerId"`
	AvatarURL                  string      `json:"Avatar_URL"`
	CreatedDatetime            string      `json:"Created_Datetime"`
	HoursPlayed                int         `json:"HoursPlayed"`
	ID                         int         `json:"Id"`
	LastLoginDatetime          string      `json:"Last_Login_Datetime"`
	Leaves                     int         `json:"Leaves"`
	Level                      int         `json:"Level"`
	Losses                     int         `json:"Losses"`
	MasteryLevel               int         `json:"MasteryLevel"`
	MergedPlayers              interface{} `json:"MergedPlayers"`
	Name                       string      `json:"Name"`
	PersonalStatusMessage      string      `json:"Personal_Status_Message"`
	RankStatConquest           float64     `json:"Rank_Stat_Conquest"`
	RankStatConquestController float64     `json:"Rank_Stat_Conquest_Controller"`
	RankStatDuel               float64     `json:"Rank_Stat_Duel"`
	RankStatDuelController     float64     `json:"Rank_Stat_Duel_Controller"`
	RankStatJoust              float64     `json:"Rank_Stat_Joust"`
	RankStatJoustController    float64     `json:"Rank_Stat_Joust_Controller"`
	RankedConquest             RankedGame  `json:"RankedConquest"`
	RankedConquestController   RankedGame  `json:"RankedConquestController"`
	RankedDuel                 RankedGame  `json:"RankedDuel"`
	RankedDuelController       RankedGame  `json:"RankedDuelController"`
	RankedJoust                RankedGame  `json:"RankedJoust"`
	RankedJoustController      RankedGame  `json:"RankedJoustController"`
	Region                     string      `json:"Region"`
	TeamID                     int         `json:"TeamId"`
	TeamName                   string      `json:"Team_Name"`
	TierConquest               int         `json:"Tier_Conquest"`
	TierDuel                   int         `json:"Tier_Duel"`
	TierJoust                  int         `json:"Tier_Joust"`
	TotalAchievements          int         `json:"Total_Achievements"`
	TotalWorshippers           int         `json:"Total_Worshippers"`
	Wins                       int         `json:"Wins"`
	HzGamerTag                 interface{} `json:"hz_gamer_tag"`
	HzPlayerName               string      `json:"hz_player_name"`
	RetMsg                     interface{} `json:"ret_msg"`
}
