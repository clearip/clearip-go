package clearip

import "fmt"

// Client main client for clear ip
type Client struct {
	HTTPClient CLHTTPClient
	IPInfo     IPInfoRepository
	BaseURI    string
	APIKey     string
}

// BaseURI base url for clearip
const defaultBaseURI = "https://api.clearip.io"

// NewClient returns a new ClearIp API client, configured with default HTTPClient.
func NewClient(apiKey string) (*Client, error) {

	if len(apiKey) == 0 {
		return nil, fmt.Errorf("API key required")
	}
	clearip := Client{APIKey: apiKey, BaseURI: defaultBaseURI}
	clearip.HTTPClient = NewHTTPClient(clearip.APIKey, clearip.BaseURI)
	clearip.setup()
	return &clearip, nil
}

func (c *Client) setup() {
	c.IPInfo = IPInfoAPI{HTTP: c.HTTPClient}

}
