package ddrago

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
)

const (
	ddragonURL = "https://ddragon.leagueoflegends.com"
)

// This is the default HTTP client used by the package. Override this value if you wish to use your own
var (
	HTTPClient = http.DefaultClient
)

// Errors that originate from this package
var (
	ErrNotFound = errors.New("not found")
)

func dispatchAndUnmarshal(endpoint string, out interface{}) error {
	resp, err := HTTPClient.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return fmt.Errorf("get %v gave status %v %v", endpoint, resp.StatusCode, resp.Status)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}

func fetchImage(endpoint string) (image.Image, error) {
	resp, err := HTTPClient.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("get %v gave status %v %v", endpoint, resp.StatusCode, resp.Status)
	}
	img, _ , err := image.Decode(resp.Body)
	return img, err
}

func Versions() (versions []Version, err error) {
	err = dispatchAndUnmarshal(ddragonURL+"/api/versions.json", &versions)
	return
}

func Languages() (languages []Language, err error) {
	err = dispatchAndUnmarshal(ddragonURL+"/cdn/languages.json", &languages)
	return
}

func ChampionsBrief(language Language, version Version) (map[ChampionID]ChampionBrief, error) {
	var cb championBrief
	err := dispatchAndUnmarshal(fmt.Sprintf(ddragonURL+"/cdn/%v/data/%v/champion.json", version, language), &cb)
	return cb.Data, err
}

func Champions(language Language, version Version) (map[ChampionID]ChampionInfo, error) {
	var cb champions
	err := dispatchAndUnmarshal(fmt.Sprintf(ddragonURL+"/cdn/%v/data/%v/championFull.json", version, language), &cb)
	return cb.Data, err
}

func Champion(language Language, version Version, id ChampionID) (ChampionInfo, error) {
	var cb champion
	err := dispatchAndUnmarshal(fmt.Sprintf(ddragonURL+"/cdn/%v/data/%v/champion/%v.json", version, language, id), &cb)
	if err != nil {
		return ChampionInfo{}, err
	}
	for _, c := range cb.Data {
		return c, nil
	}
	return ChampionInfo{}, ErrNotFound
}

func Splash(id ChampionID, skinNum int) (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/img/champion/splash/%v_%v.jpg", id, skinNum))
}

func LoadingScreen(id ChampionID, skinNum int) (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/img/champion/loading/%v_%v.jpg", id, skinNum))
}

func Tile(id ChampionID, skinNum int) (image.Image, error) {
	return fetchImage(fmt.Sprintf(ddragonURL+"/cdn/img/champion/tiles/%v_%v.jpg", id, skinNum))
}

