package clearip

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// IPInfoRepository holds the functions for dealing with ip info
type IPInfoRepository interface {
	GetAllDataByIP(ip string) (interface{}, error)
}

// IPInfoAPI implements IPInfoRepository
type IPInfoAPI struct {
	HTTP CLHTTPClient
}

// GetAllDataByIP send a get request for ip info
func (c IPInfoAPI) GetAllDataByIP(ip string) (interface{}, error) {
	var IPInfo interface{}

	matches, _ := regexp.MatchString(`^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, ip)
	if !matches {
		return nil, fmt.Errorf("Ip required")
	}

	getIPInfoURL := strings.Join([]string{"/ip/", ip, "/json?apikey=", c.HTTP.APIKey}, "")
	data, err := c.HTTP.Get(getIPInfoURL, nil)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &IPInfo)
	return IPInfo, err

}
