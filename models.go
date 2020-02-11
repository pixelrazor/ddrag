package ddrago

import (
	"fmt"
	"image"
)

// Versions represents a given version number in the form of major.minor.patch
type Version string

// Language represents a language constant used for localization
type Language string

// ChampionID is the ID of a champions (normally, but not always, just the champion name)
type ChampionID string

// All supported languages
const (
	Czech                 Language = "cs_CZ"
	Greek                 Language = "el_GR"
	Polish                Language = "pl_PL"
	Romanian              Language = "ro_RO"
	Hungarian             Language = "hu_HU"
	English_UnitedKingdom Language = "en_GB"
	German                Language = "de_DE"
	Spanish               Language = "es_ES"
	Italian               Language = "it_IT"
	French                Language = "fr_FR"
	Japanese              Language = "ja_JP"
	Korean                Language = "ko_KR"
	Spanish_Mexico        Language = "es_MX"
	Spanish_Argentina     Language = "es_AR"
	Portuguese            Language = "pt_BR"
	English_UnitedStates  Language = "en_US"
	English_Australia     Language = "en_AU"
	Russian               Language = "ru_RU"
	Turkish               Language = "tr_TR"
	Malay                 Language = "ms_MY"
	English_Philippines   Language = "en_PH"
	English_Singapore     Language = "en_SG"
	Thai                  Language = "th_TH"
	Vietnamese            Language = "vn_VN"
	Indonesian            Language = "id_ID"
	Chinese_Malaysia      Language = "zh_MY"
	Chinese_China         Language = "zh_CN"
	Chinese_Taiwan        Language = "zh_TW"
)

type championBrief struct {
	Type    string                       `json:"type"`
	Format  string                       `json:"format"`
	Version string                       `json:"version"`
	Data    map[ChampionID]ChampionBrief `json:"data"`
}

type champions struct {
	Type    string                      `json:"type"`
	Format  string                      `json:"format"`
	Version Version                     `json:"version"`
	Data    map[ChampionID]ChampionInfo `json:"data"`
}

type ChampionBrief struct {
	Version string     `json:"version"`
	ID      ChampionID `json:"id"`
	Key     string     `json:"key"`
	Name    string     `json:"name"`
	Title   string     `json:"title"`
	Blurb   string     `json:"blurb"`
	Info    struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   struct {
		Hp                   float64 `json:"hp"`
		Hpperlevel           float64 `json:"hpperlevel"`
		Mp                   float64 `json:"mp"`
		Mpperlevel           float64 `json:"mpperlevel"`
		Movespeed            float64 `json:"movespeed"`
		Armor                float64 `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           float64 `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
		Attackrange          float64 `json:"attackrange"`
		Hpregen              float64 `json:"hpregen"`
		Hpregenperlevel      float64 `json:"hpregenperlevel"`
		Mpregen              float64 `json:"mpregen"`
		Mpregenperlevel      float64 `json:"mpregenperlevel"`
		Crit                 float64 `json:"crit"`
		Critperlevel         float64 `json:"critperlevel"`
		Attackdamage         float64 `json:"attackdamage"`
		Attackdamageperlevel float64 `json:"attackdamageperlevel"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
		Attackspeed          float64 `json:"attackspeed"`
	} `json:"stats"`
}

type champion struct {
	Type    string                  `json:"type"`
	Format  string                  `json:"format"`
	Version Version                 `json:"version"`
	Data    map[string]ChampionInfo `json:"data"`
}

type ChampionInfo struct {
	version Version
	ID      ChampionID `json:"id"`
	Key     string     `json:"key"`
	Name    string     `json:"name"`
	Title   string     `json:"title"`
	Image   struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	} `json:"image"`
	Skins []struct {
		ID      string `json:"id"`
		Num     int    `json:"num"`
		Name    string `json:"name"`
		Chromas bool   `json:"chromas"`
	} `json:"skins"`
	Lore      string   `json:"lore"`
	Blurb     string   `json:"blurb"`
	Allytips  []string `json:"allytips"`
	Enemytips []string `json:"enemytips"`
	Tags      []string `json:"tags"`
	Partype   string   `json:"partype"`
	Info      struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
	} `json:"info"`
	Stats struct {
		Hp                   float64 `json:"hp"`
		Hpperlevel           float64 `json:"hpperlevel"`
		Mp                   float64 `json:"mp"`
		Mpperlevel           float64 `json:"mpperlevel"`
		Movespeed            float64 `json:"movespeed"`
		Armor                float64 `json:"armor"`
		Armorperlevel        float64 `json:"armorperlevel"`
		Spellblock           float64 `json:"spellblock"`
		Spellblockperlevel   float64 `json:"spellblockperlevel"`
		Attackrange          float64 `json:"attackrange"`
		Hpregen              float64 `json:"hpregen"`
		Hpregenperlevel      float64 `json:"hpregenperlevel"`
		Mpregen              float64 `json:"mpregen"`
		Mpregenperlevel      float64 `json:"mpregenperlevel"`
		Crit                 float64 `json:"crit"`
		Critperlevel         float64 `json:"critperlevel"`
		Attackdamage         float64 `json:"attackdamage"`
		Attackdamageperlevel float64 `json:"attackdamageperlevel"`
		Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
		Attackspeed          float64 `json:"attackspeed"`
	} `json:"stats"`
	Spells []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Tooltip     string `json:"tooltip"`
		Leveltip    struct {
			Label  []string `json:"label"`
			Effect []string `json:"effect"`
		} `json:"leveltip"`
		Maxrank      int       `json:"maxrank"`
		Cooldown     []float64 `json:"cooldown"`
		CooldownBurn string    `json:"cooldownBurn"`
		Cost         []int     `json:"cost"`
		CostBurn     string    `json:"costBurn"`
		Datavalues   struct {
		} `json:"datavalues"`
		Effect     []interface{} `json:"effect"`
		EffectBurn []interface{} `json:"effectBurn"`
		Vars       []interface{} `json:"vars"`
		CostType   string        `json:"costType"`
		Maxammo    string        `json:"maxammo"`
		Range      []int         `json:"range"`
		RangeBurn  string        `json:"rangeBurn"`
		Image      struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
		Resource string `json:"resource"`
	} `json:"spells"`
	Passive struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       struct {
			Full   string `json:"full"`
			Sprite string `json:"sprite"`
			Group  string `json:"group"`
			X      int    `json:"x"`
			Y      int    `json:"y"`
			W      int    `json:"w"`
			H      int    `json:"h"`
		} `json:"image"`
	} `json:"passive"`
	Recommended []struct {
		Champion            string      `json:"champion"`
		Title               string      `json:"title"`
		Map                 string      `json:"map"`
		Mode                string      `json:"mode"`
		Type                string      `json:"type"`
		CustomTag           string      `json:"customTag"`
		Sortrank            int         `json:"sortrank,omitempty"`
		ExtensionPage       bool        `json:"extensionPage"`
		UseObviousCheckmark bool        `json:"useObviousCheckmark,omitempty"`
		CustomPanel         interface{} `json:"customPanel"`
		Blocks              []struct {
			Type                string   `json:"type"`
			RecMath             bool     `json:"recMath"`
			RecSteps            bool     `json:"recSteps"`
			MinSummonerLevel    int      `json:"minSummonerLevel"`
			MaxSummonerLevel    int      `json:"maxSummonerLevel"`
			ShowIfSummonerSpell string   `json:"showIfSummonerSpell"`
			HideIfSummonerSpell string   `json:"hideIfSummonerSpell"`
			AppendAfterSection  string   `json:"appendAfterSection"`
			VisibleWithAllOf    []string `json:"visibleWithAllOf"`
			HiddenWithAnyOf     []string `json:"hiddenWithAnyOf"`
			Items               []struct {
				ID        string `json:"id"`
				Count     int    `json:"count"`
				HideCount bool   `json:"hideCount"`
			} `json:"items"`
		} `json:"blocks"`
	} `json:"recommended"`
}

func (ci ChampionBrief) FetchSquareImg() (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/%v/img/passive/%v", ci.Version, ci.Image.Full))
}

func (ci ChampionInfo) FetchPassiveImg() (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/%v/img/passive/%v", ci.version, ci.Passive.Image.Full))
}

func (ci ChampionInfo) FetchSquareImg() (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/%v/img/passive/%v", ci.version, ci.Image.Full))
}

func (ci ChampionInfo) FetchSpellImg(which int) (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/%v/img/passive/%v", ci.version, ci.Spells[which].Image.Full))
}
