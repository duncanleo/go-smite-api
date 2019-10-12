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

// GetPlayerAchievementsResponse response when getting player achievements
type GetPlayerAchievementsResponse struct {
	AssistedKills        int         `json:"AssistedKills"`
	CampsCleared         int         `json:"CampsCleared"`
	Deaths               int         `json:"Deaths"`
	DivineSpree          int         `json:"DivineSpree"`
	DoubleKills          int         `json:"DoubleKills"`
	FireGiantKills       int         `json:"FireGiantKills"`
	FirstBloods          int         `json:"FirstBloods"`
	GodLikeSpree         int         `json:"GodLikeSpree"`
	GoldFuryKills        int         `json:"GoldFuryKills"`
	PlayerID             int         `json:"Id"`
	ImmortalSpree        int         `json:"ImmortalSpree"`
	KillingSpree         int         `json:"KillingSpree"`
	MinionKills          int         `json:"MinionKills"`
	Name                 string      `json:"Name"`
	PentaKills           int         `json:"PentaKills"`
	PhoenixKills         int         `json:"PhoenixKills"`
	PlayerKills          int         `json:"PlayerKills"`
	QuadraKills          int         `json:"QuadraKills"`
	RampageSpree         int         `json:"RampageSpree"`
	ShutdownSpree        int         `json:"ShutdownSpree"`
	SiegeJuggernautKills int         `json:"SiegeJuggernautKills"`
	TowerKills           int         `json:"TowerKills"`
	TripleKills          int         `json:"TripleKills"`
	UnstoppableSpree     int         `json:"UnstoppableSpree"`
	WildJuggernautKills  int         `json:"WildJuggernautKills"`
	RetMsg               interface{} `json:"ret_msg"`
}

// GetPlayerGodRanksResponse response when getting god ranks for a player
type GetPlayerGodRanksResponse []struct {
	Assists     int         `json:"Assists"`
	Deaths      int         `json:"Deaths"`
	Kills       int         `json:"Kills"`
	Losses      int         `json:"Losses"`
	MinionKills int         `json:"MinionKills"`
	Rank        int         `json:"Rank"`
	Wins        int         `json:"Wins"`
	Worshippers int         `json:"Worshippers"`
	God         string      `json:"god"`
	GodID       string      `json:"god_id"`
	PlayerID    string      `json:"player_id"`
	RetMsg      interface{} `json:"ret_msg"`
}

// GetPlayerMatchHistoryResponse response about a player's match history
type GetPlayerMatchHistoryResponse []struct {
	ActiveID1           int         `json:"ActiveId1"`
	ActiveID2           int         `json:"ActiveId2"`
	Active1             string      `json:"Active_1"`
	Active2             string      `json:"Active_2"`
	Active3             interface{} `json:"Active_3"`
	Assists             int         `json:"Assists"`
	Ban1                string      `json:"Ban1"`
	Ban10               string      `json:"Ban10"`
	Ban10ID             int         `json:"Ban10Id"`
	Ban1ID              int         `json:"Ban1Id"`
	Ban2                string      `json:"Ban2"`
	Ban2ID              int         `json:"Ban2Id"`
	Ban3                string      `json:"Ban3"`
	Ban3ID              int         `json:"Ban3Id"`
	Ban4                string      `json:"Ban4"`
	Ban4ID              int         `json:"Ban4Id"`
	Ban5                string      `json:"Ban5"`
	Ban5ID              int         `json:"Ban5Id"`
	Ban6                string      `json:"Ban6"`
	Ban6ID              int         `json:"Ban6Id"`
	Ban7                string      `json:"Ban7"`
	Ban7ID              int         `json:"Ban7Id"`
	Ban8                string      `json:"Ban8"`
	Ban8ID              int         `json:"Ban8Id"`
	Ban9                string      `json:"Ban9"`
	Ban9ID              int         `json:"Ban9Id"`
	Creeps              int         `json:"Creeps"`
	Damage              int         `json:"Damage"`
	DamageBot           int         `json:"Damage_Bot"`
	DamageDoneInHand    int         `json:"Damage_Done_In_Hand"`
	DamageMitigated     int         `json:"Damage_Mitigated"`
	DamageStructure     int         `json:"Damage_Structure"`
	DamageTaken         int         `json:"Damage_Taken"`
	DamageTakenMagical  int         `json:"Damage_Taken_Magical"`
	DamageTakenPhysical int         `json:"Damage_Taken_Physical"`
	Deaths              int         `json:"Deaths"`
	DistanceTraveled    int         `json:"Distance_Traveled"`
	FirstBanSide        string      `json:"First_Ban_Side"`
	God                 string      `json:"God"`
	GodID               int         `json:"GodId"`
	Gold                int         `json:"Gold"`
	Healing             int         `json:"Healing"`
	HealingBot          int         `json:"Healing_Bot"`
	HealingPlayerSelf   int         `json:"Healing_Player_Self"`
	ItemID1             int         `json:"ItemId1"`
	ItemID2             int         `json:"ItemId2"`
	ItemID3             int         `json:"ItemId3"`
	ItemID4             int         `json:"ItemId4"`
	ItemID5             int         `json:"ItemId5"`
	ItemID6             int         `json:"ItemId6"`
	Item1               string      `json:"Item_1"`
	Item2               string      `json:"Item_2"`
	Item3               string      `json:"Item_3"`
	Item4               string      `json:"Item_4"`
	Item5               string      `json:"Item_5"`
	Item6               string      `json:"Item_6"`
	KillingSpree        int         `json:"Killing_Spree"`
	Kills               int         `json:"Kills"`
	Level               int         `json:"Level"`
	MapGame             string      `json:"Map_Game"`
	Match               int         `json:"Match"`
	MatchQueueID        int         `json:"Match_Queue_Id"`
	MatchTime           string      `json:"Match_Time"`
	Minutes             int         `json:"Minutes"`
	MultiKillMax        int         `json:"Multi_kill_Max"`
	ObjectiveAssists    int         `json:"Objective_Assists"`
	Queue               string      `json:"Queue"`
	Region              string      `json:"Region"`
	Skin                string      `json:"Skin"`
	SkinID              int         `json:"SkinId"`
	Surrendered         int         `json:"Surrendered"`
	TaskForce           int         `json:"TaskForce"`
	Team1Score          int         `json:"Team1Score"`
	Team2Score          int         `json:"Team2Score"`
	TimeInMatchSeconds  int         `json:"Time_In_Match_Seconds"`
	WardsPlaced         int         `json:"Wards_Placed"`
	WinStatus           string      `json:"Win_Status"`
	WinningTaskForce    int         `json:"Winning_TaskForce"`
	PlayerID            int         `json:"playerId"`
	PlayerName          string      `json:"playerName"`
	RetMsg              interface{} `json:"ret_msg"`
}
