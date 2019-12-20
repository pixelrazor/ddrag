package ddrago

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ddragonURL = "https://ddragon.leagueoflegends.com"
)

type Client struct {
	client *http.Client
}

func New(client *http.Client) Client {
	return Client{client: client}
}

func (c Client) dispatchAndUnmarshal(endpoint string, out interface{}) error {
	resp, err := c.client.Get(endpoint)
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

func (c Client) Versions() (versions []Version, err error) {
	err = c.dispatchAndUnmarshal(ddragonURL+"/api/versions.json", &versions)
	return
}