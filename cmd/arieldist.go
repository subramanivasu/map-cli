package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/subramanivas/map-cli/pkg/config"
	"github.com/subramanivas/map-cli/pkg/mappls"
	"github.com/subramanivas/map-cli/pkg/nominatim"
)

var (
	fromCoords string
	toCoords   string
)

var arieldistCmd = &cobra.Command{
	Use:   "arieldist [source location] [destination location]",
	Short: "Calculate aerial distance between two locations",
	Long: `Calculate the aerial distance between two locations using either location names or direct coordinates.

You can provide:
1. Location names as arguments: map arieldist 'Yelehanka' 'Koramangala'
2. Direct coordinates using flags: map arieldist --from '13.115;77.607' --to '12.935; 77.624'
3. Mix of both: source as argument, destination as coordinates, or vice versa

The format for coordinates is: latitude;longitude (e.g., 12.97;77.63)
The distance is calculated using the Mappls API and will be displayed in kilometers.`,
	RunE: runArieldist,
}

func runArieldist(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	nominatimClient := nominatim.NewClient()
	maplsClient := mappls.NewClient(cfg.MaplsAccessToken)

	srcLat, srcLon, err := getCoordinates(nominatimClient, args, 0, fromCoords, "source")
	if err != nil {
		return err
	}

	dstLat, dstLon, err := getCoordinates(nominatimClient, args, 1, toCoords, "destination")
	if err != nil {
		return err
	}

	distance, err := maplsClient.GetDistance(srcLat, srcLon, dstLat, dstLon, "K")
	if err != nil {
		return fmt.Errorf("failed to get distance: %w", err)
	}

	fmt.Printf("Distance: %.2f km\n", distance)
	return nil
}

func getCoordinates(
	nomClient *nominatim.Client,
	args []string,
	argIndex int,
	coordsFlag string,
	locationType string,
) (lat, lon string, err error) {
	if coordsStr := strings.TrimSpace(coordsFlag); coordsStr != "" {
		lat, lon, err := parseCoordinates(coordsStr)
		if err != nil {
			return "", "", fmt.Errorf("invalid %s coordinates format: %w", locationType, err)
		}
		return lat, lon, nil
	}

	if argIndex < len(args) {
		input := strings.TrimSpace(args[argIndex])
		if input == "" {
			return "", "", fmt.Errorf("%s location cannot be empty", locationType)
		}

		if strings.Contains(input, ";") {
			if lat, lon, err := parseCoordinates(input); err == nil {
				return lat, lon, nil
			}
		}

		result, err := nomClient.Geocode(input)
		if err != nil {
			return "", "", fmt.Errorf("failed to geocode %s location: %w", locationType, err)
		}

		return result.Latitude, result.Longitude, nil
	}

	return "", "", fmt.Errorf(
		"missing %s location: provide as argument (location name or 'lat;lon') or use --%s flag (format: lat;lon)",
		locationType, strings.ToLower(locationType))
}

func parseCoordinates(coordsStr string) (lat, lon string, err error) {
	parts := strings.Split(strings.TrimSpace(coordsStr), ";")
	if len(parts) != 2 {
		return "", "", fmt.Errorf("expected format 'latitude;longitude', got '%s'", coordsStr)
	}

	lat = strings.TrimSpace(parts[0])
	lon = strings.TrimSpace(parts[1])

	if lat == "" || lon == "" {
		return "", "", fmt.Errorf("latitude and longitude cannot be empty")
	}

	return lat, lon, nil
}

func init() {
	rootCmd.AddCommand(arieldistCmd)

	arieldistCmd.Flags().StringVar(&fromCoords, "from", "", "Source coordinates in format: latitude;longitude (e.g., 12.97;77.63)")
	arieldistCmd.Flags().StringVar(&toCoords, "to", "", "Destination coordinates in format: latitude;longitude (e.g., 13.08;80.27)")
}
