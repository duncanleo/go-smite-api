package smite

// God is a god character in SMITE
type God struct {
	Ability1                   GodAbility  `json:"Ability_1"`
	Ability2                   GodAbility  `json:"Ability_2"`
	Ability3                   GodAbility  `json:"Ability_3"`
	Ability4                   GodAbility  `json:"Ability_4"`
	Ability5                   GodAbility  `json:"Ability_5"`
	AttackSpeed                float64     `json:"AttackSpeed"`
	AttackSpeedPerLevel        float64     `json:"AttackSpeedPerLevel"`
	Cons                       string      `json:"Cons"`
	HP5PerLevel                float64     `json:"HP5PerLevel"`
	Health                     int         `json:"Health"`
	HealthPerFive              int         `json:"HealthPerFive"`
	HealthPerLevel             int         `json:"HealthPerLevel"`
	Lore                       string      `json:"Lore"`
	MP5PerLevel                float64     `json:"MP5PerLevel"`
	MagicProtection            int         `json:"MagicProtection"`
	MagicProtectionPerLevel    float64     `json:"MagicProtectionPerLevel"`
	MagicalPower               int         `json:"MagicalPower"`
	MagicalPowerPerLevel       float64     `json:"MagicalPowerPerLevel"`
	Mana                       int         `json:"Mana"`
	ManaPerFive                float64     `json:"ManaPerFive"`
	ManaPerLevel               int         `json:"ManaPerLevel"`
	Name                       string      `json:"Name"`
	OnFreeRotation             string      `json:"OnFreeRotation"`
	Pantheon                   string      `json:"Pantheon"`
	PhysicalPower              int         `json:"PhysicalPower"`
	PhysicalPowerPerLevel      float64     `json:"PhysicalPowerPerLevel"`
	PhysicalProtection         int         `json:"PhysicalProtection"`
	PhysicalProtectionPerLevel float64     `json:"PhysicalProtectionPerLevel"`
	Pros                       string      `json:"Pros"`
	Roles                      string      `json:"Roles"`
	Speed                      int         `json:"Speed"`
	Title                      string      `json:"Title"`
	Type                       string      `json:"Type"`
	AbilityDescription1        Description `json:"abilityDescription1"`
	AbilityDescription2        Description `json:"abilityDescription2"`
	AbilityDescription3        Description `json:"abilityDescription3"`
	AbilityDescription4        Description `json:"abilityDescription4"`
	AbilityDescription5        Description `json:"abilityDescription5"`
	BasicAttack                Description `json:"basicAttack"`
	GodAbility1URL             string      `json:"godAbility1_URL"`
	GodAbility2URL             string      `json:"godAbility2_URL"`
	GodAbility3URL             string      `json:"godAbility3_URL"`
	GodAbility4URL             string      `json:"godAbility4_URL"`
	GodAbility5URL             string      `json:"godAbility5_URL"`
	GodCardURL                 string      `json:"godCard_URL"`
	GodIconURL                 string      `json:"godIcon_URL"`
	ID                         int         `json:"id"`
	LatestGod                  string      `json:"latestGod"`
	RetMsg                     interface{} `json:"ret_msg"`
}

// GodAbility represents one of a god's abilities in SMITE
type GodAbility struct {
	Description Description `json:"Description"`
	ID          int         `json:"Id"`
	Summary     string      `json:"Summary"`
	URL         string      `json:"URL"`
}

// Item represents a god's item in SMITE
type Item struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// ItemDescription description of a god's item in SMITE
type ItemDescription struct {
	Cooldown    string `json:"cooldown"`
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Menuitems   []Item `json:"menuitems"`
	Rankitems   []Item `json:"rankitems"`
}

// Description generic description container
type Description struct {
	ItemDescription ItemDescription `json:"itemDescription"`
}

// RankedGame represents a ranked game
type RankedGame struct {
	Leaves   int         `json:"Leaves"`
	Losses   int         `json:"Losses"`
	Name     string      `json:"Name"`
	Points   int         `json:"Points"`
	PrevRank int         `json:"PrevRank"`
	Rank     int         `json:"Rank"`
	RankStat float64     `json:"Rank_Stat"`
	Season   int         `json:"Season"`
	Tier     int         `json:"Tier"`
	Trend    int         `json:"Trend"`
	Wins     int         `json:"Wins"`
	PlayerID interface{} `json:"player_id"`
	RetMsg   interface{} `json:"ret_msg"`
}
