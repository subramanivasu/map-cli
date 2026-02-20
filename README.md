# map-it

A command-line tool for map operations powered by [Mappls](https://www.mappls.com/) and [OpenStreetMap's Nominatim](https://nominatim.org/). Calculate distances between locations, discover nearby places, and explore geographical data—all from your terminal.

## Features

### Current Features

- **Aerial Distance Calculator** (`arieldist`)
  - Calculate the straight-line distance between two locations
  - Use location names for easy, human-readable queries
  - Use coordinates for precise calculations
  - Mix and match both approaches in a single command
  - Results displayed in kilometers

<<<<<<< HEAD
### Planned Features

- **Nearby Places Finder**
  - Search for specific types of places (cafes, restaurants, parks, etc.) in a given location
  - Get top results sorted by relevance
  - Filter by radius and other criteria
=======
- **Nearby Places Finder** (`nearby`)
  - Search for specific types of places (cafes, restaurants, parks, etc.) around a location
  - Get top 5 results sorted by distance
  - Support multiple keywords with OR and AND operators
  - Display place information including address, phone, and email
>>>>>>> f695d0e (Update README.md)

## Prerequisites

Before you get started, you'll need:

- **Go 1.22.2 or higher** – [Download here](https://golang.org/dl/)
- **Mappls API Access Token** – Get one free at [Mappls Developer Console](https://developer.mappls.com/)

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/subramanivas/map-it.git
cd map-it
```

### 2. Set Up Environment Variables

Create a `.env` file in the project root (or copy from `.env.example`):

```bash
cp .env.example .env
```

Edit `.env` and add your Mappls API token:

```
MAPPLS_ACCESS_TOKEN=your_access_token_here
```

### 3. Install Dependencies

```bash
go mod download
```

### 4. Build the Binary

```bash
go build -o map
```

This creates an executable called `map` that you can use immediately.

### 5. (Optional) Install Globally

To use the tool from anywhere on your system:

```bash
go install
```

This installs the binary to your `$GOPATH/bin` directory.

## Usage

### Aerial Distance Calculator

Calculate the straight-line distance between two locations in multiple ways:

#### Using Location Names

```bash
./map arieldist 'Yelehanka' 'Koramangala'
```

Output:
```
Distance: 12.50 km
```

#### Using Coordinates

Provide coordinates in the format `latitude;longitude`:

```bash
./map arieldist --from '13.115;77.607' --to '12.935;77.624'
```

#### Mix Location Names and Coordinates

You can combine both approaches:

```bash
# Source as location name, destination as coordinates
./map arieldist 'Bangalore Airport' --to '12.935;77.624'

# Source as coordinates, destination as location name
./map arieldist --from '13.115;77.607' 'Indiranagar'
```

#### Available Flags

- `--from` – Source coordinates in format `latitude;longitude` (optional if provided as argument)
- `--to` – Destination coordinates in format `latitude;longitude` (optional if provided as argument)

For more information:

```bash
./map arieldist --help
```

<<<<<<< HEAD
=======
### Nearby Places Finder

Search for nearby places (cafes, restaurants, hospitals, etc.) around a specific location.

#### Search by Keywords

```bash
./map nearby coffee --refLocation 28.631460,77.217423
```

Output:
```
Found 10 places (showing top 5):

1. Coffee
   Address: Outer Circle, Connaught Place, New Delhi, Delhi, 110001
   Distance: 82 m

2. Digging Cafe
   Address: 12 A, 1st Floor, Connaught Place, New Delhi, Delhi, 110001
   Distance: 84 m

...
```

#### Multiple Keywords with OR Operator (`;`)

Provide multiple arguments separated by spaces:

```bash
./map nearby coffee tea --refLocation 28.631460,77.217423
```

Or use quoted string with semicolon:

```bash
./map nearby "coffee;tea" --refLocation 28.631460,77.217423
```

#### Multiple Keywords with AND Operator (`$`)

Use a quoted string with dollar sign operator:

```bash
./map nearby "coffee $ food" --refLocation 28.631460,77.217423
```

#### Keyword Operators

- **OR operator (`;`)**: Results matching either keyword
  - Multiple arguments: `./map nearby coffee tea` → Searches for "coffee;tea"
  - Quoted string: `./map nearby "coffee;tea"`

- **AND operator (`$`)**: Results matching all keywords
  - Must use quoted string: `./map nearby "coffee $ food"`

#### Available Flags

- `--refLocation` – Reference location in format `latitude,longitude` **(REQUIRED)**
  - Example: `28.631460,77.217423`

For more information:

```bash
./map nearby --help
```

>>>>>>> f695d0e (Update README.md)
## Project Structure

```
map-it/
├── main.go                    # Application entry point
├── cmd/
│   ├── root.go               # Root command configuration
<<<<<<< HEAD
│   └── arieldist.go          # Aerial distance command implementation
=======
│   ├── arieldist.go          # Aerial distance command implementation
│   └── nearby.go             # Nearby places search command implementation
>>>>>>> f695d0e (Update README.md)
├── pkg/
│   ├── config/
│   │   └── config.go         # Configuration and environment setup
│   ├── mappls/
│   │   └── client.go         # Mappls API client
│   └── nominatim/
│       └── client.go         # Nominatim geocoding client
├── go.mod                     # Go module definition
├── .env.example               # Example environment variables
└── README.md                  # This file
```

## How It Works

1. **Input Processing**: The tool accepts locations as either human-readable place names or precise coordinates.

2. **Geocoding**: When you provide a place name (e.g., "Yelehanka"), the tool uses Nominatim's geocoding service to convert it to latitude and longitude coordinates.

3. **Distance Calculation**: With coordinates from both source and destination, the tool calls the Mappls API, which calculates the aerial (straight-line) distance over the Earth's surface.

4. **Result**: The distance is calculated in kilometers and displayed to you.

## Development

### Running Tests

```bash
go test ./...
```

### Building with Different OS Targets

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 go build -o map-linux

# Build for macOS
GOOS=darwin GOARCH=amd64 go build -o map-macos

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o map.exe
```

## API Credits

- **Mappls** – For distance calculation API
  - [Documentation](https://developer.mappls.com/)
  - Free tier available with registration

- **Nominatim (OpenStreetMap)** – For geocoding service
  - Completely free and open
  - [Usage Policy](https://nominatim.org/usage_policy.html)

## Troubleshooting

### "missing required environment variable: MAPPLS_ACCESS_TOKEN"

Make sure you've:

1. Created a `.env` file in the project root
2. Added your Mappls API token
3. Are running the tool from the project directory (or have the `.env` in your working directory)

### No Results from Location Names

- Ensure the location name is spelled correctly
- Try using the full address or more specific location names
- Check your internet connection (Nominatim requires online access)

### Build Errors

- Verify you have Go 1.22.2 or higher: `go version`
- Run `go mod download` to fetch dependencies
- Check that all required packages are available

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

MIT License – See LICENSE file for details.