package mappls

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL        = "https://tile.mappls.com/map/raster_tile/distanceA"
	defaultTimeout = 10 * time.Second
)

type Client struct {
	httpClient  *http.Client
	accessToken string
}

type DistanceResponse struct {
	ResponseCode int     `json:"responseCode"`
	Distance     float64 `json:"distance"`
	Unit         string  `json:"unit"`
}

func NewClient(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (c *Client) GetDistance(fromLat, fromLon, toLat, toLon string, unit string) (float64, error) {

	if err := validateCoordinates(fromLat, fromLon); err != nil {
		return 0, fmt.Errorf("invalid source coordinates: %w", err)
	}
	if err := validateCoordinates(toLat, toLon); err != nil {
		return 0, fmt.Errorf("invalid destination coordinates: %w", err)
	}

	if strings.TrimSpace(unit) == "" {
		unit = "K"
	}

	url := fmt.Sprintf("%s?from=%s,%s&to=%s,%s&unit=%s&access_token=%s",
		baseURL, fromLat, fromLon, toLat, toLon, unit, c.accessToken)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to call Mappls API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("mappls API returned status %d: %s", resp.StatusCode, string(body))
	}

	var result DistanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to parse Mappls response: %w", err)
	}

	if result.ResponseCode != http.StatusOK {
		return 0, fmt.Errorf("mappls API returned response code %d", result.ResponseCode)
	}

	distance := roundToTwoDecimals(result.Distance)
	return distance, nil
}

func validateCoordinates(lat, lon string) error {
	if strings.TrimSpace(lat) == "" || strings.TrimSpace(lon) == "" {
		return fmt.Errorf("latitude and longitude cannot be empty")
	}

	if _, err := strconv.ParseFloat(lat, 64); err != nil {
		return fmt.Errorf("invalid latitude: %w", err)
	}

	if _, err := strconv.ParseFloat(lon, 64); err != nil {
		return fmt.Errorf("invalid longitude: %w", err)
	}

	return nil
}

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
