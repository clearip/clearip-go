package clearip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Client is wrapper for net/http client with additional variables
type Client struct {
	APIKey string
	URL    string
	HTTP   *http.Client
}

// NewClient returns custome http client
func NewClient(apiKey string) (*Client, error) {

	if len(apiKey) == 0 {
		return nil, fmt.Errorf("API key required")
	}

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return &Client{
		URL:    "https://api.clearip.io",
		APIKey: apiKey,
		HTTP:   httpClient,
	}, nil
}

// GetIpinfo send a get request for ip info
func (c *Client) GetIpinfo(ip string) (map[string]interface{}, error) {

	matches, _ := regexp.MatchString(`^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, ip)
	if !matches {
		return nil, fmt.Errorf("Ip required")
	}

	getIPInfoURL := strings.Join([]string{c.URL, "/ip/", ip, "/json?apikey=", c.APIKey}, "")
	resp, err := c.HTTP.Get(getIPInfoURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("auth error")
	}

	var response = make(map[string]interface{})

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&response); err != nil {
		return nil, fmt.Errorf("error parsing json")
	}

	return response, nil

}
