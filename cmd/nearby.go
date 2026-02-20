package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/subramanivas/map-cli/pkg/config"
	"github.com/subramanivas/map-cli/pkg/mappls"
)

var (
	refLocation string
)

var nearbyCmd = &cobra.Command{
	Use:   "nearby [keywords...]",
	Short: "Search for nearby places",
	Long: `Search for nearby places using keywords and a reference location.

You must provide a reference location (latitude,longitude) to search nearby places.
You can provide multiple keywords separated by spaces or semicolon (for OR operator) or dollar sign (for AND operator).

Examples:
  map nearby coffee --refLocation 28.631460,77.217423
  map nearby coffee;tea --refLocation 28.631460,77.217423
  map nearby coffee $ food --refLocation 28.631460,77.217423`,
	RunE: runNearby,
}

func runNearby(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	if len(args) == 0 {
		return fmt.Errorf("please provide at least one keyword")
	}

	location := strings.TrimSpace(refLocation)
	if location == "" {
		return fmt.Errorf("reference location is required, use --refLocation flag with format: latitude,longitude (e.g., 28.631460,77.217423)")
	}

	keywords := strings.Join(args, ";")

	maplsClient := mappls.NewClient(cfg.MaplsAccessToken)

	locations, err := maplsClient.NearbySearch(keywords, location)
	if err != nil {
		return fmt.Errorf("failed to search nearby places: %w", err)
	}

	if len(locations) == 0 {
		fmt.Println("No places found")
		return nil
	}

	limit := 5
	if len(locations) < limit {
		limit = len(locations)
	}

	fmt.Printf("Found %d places (showing top %d):\n\n", len(locations), limit)

	for i := 0; i < limit; i++ {
		loc := locations[i]
		fmt.Printf("%d. %s\n", i+1, loc.PlaceName)
		fmt.Printf("   Address: %s\n", loc.PlaceAddress)
		fmt.Printf("   Distance: %d m\n", loc.Distance)
		if loc.MobileNo != "" {
			fmt.Printf("   Phone: %s\n", loc.MobileNo)
		}
		if loc.Email != "" {
			fmt.Printf("   Email: %s\n", loc.Email)
		}
		fmt.Println()
	}

	return nil
}

func init() {
	rootCmd.AddCommand(nearbyCmd)
	nearbyCmd.Flags().StringVar(&refLocation, "refLocation", "", "Reference location in format: latitude,longitude (e.g., 28.631460,77.217423)")
}
