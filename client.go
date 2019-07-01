package iptrace

import "fmt"

// Client main client for clear ip
type Client struct {
	HTTPClient CLHTTPClient
	IPRepo     IPInfoRepository
	BaseURI    string
	APIKey     string
}

// BaseURI base url for iptrace
const defaultBaseURI = "https://api.iptrace.io"

// NewClient returns a new iptrace API client, configured with default HTTPClient.
func NewClient(apiKey string) (*Client, error) {

	if len(apiKey) == 0 {
		return nil, fmt.Errorf("API key required")
	}
	iptrace := Client{APIKey: apiKey, BaseURI: defaultBaseURI}
	iptrace.HTTPClient = NewHTTPClient(iptrace.APIKey, iptrace.BaseURI)
	iptrace.setup()
	return &iptrace, nil
}

func (c *Client) setup() {
	c.IPRepo = IPInfoAPI{HTTP: c.HTTPClient}

}
