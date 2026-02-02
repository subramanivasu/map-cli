package nominatim

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL        = "https://nominatim.openstreetmap.org"
	defaultTimeout = 10 * time.Second
	userAgent      = "map-cli/1.0"
)

type Client struct {
	httpClient *http.Client
}

type LocationResult struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"lon"`
	Name      string `json:"name"`
	Address   string `json:"display_name"`
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}


func (c *Client) Geocode(locationName string) (*LocationResult, error) {
	if strings.TrimSpace(locationName) == "" {
		return nil, fmt.Errorf("location name cannot be empty")
	}

	query := url.QueryEscape(strings.TrimSpace(locationName))
	endpoint := fmt.Sprintf("%s/search?q=%s&format=json", baseURL, query)

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call Nominatim API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("nominatim API returned status %d: %s", resp.StatusCode, string(body))
	}

	var results []LocationResult
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, fmt.Errorf("failed to parse Nominatim response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("location '%s' not found", locationName)
	}

	return &results[0], nil
}
