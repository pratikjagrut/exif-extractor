package exif

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func GetLatLong(path string) (lat float64, long float64, err error) {
	// Open the image file.
	f, err := os.Open(path)
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("Exif: Open File: %v", err)
	}
	defer f.Close() // Ensure the file is closed once we're done with it

	// Decode the EXIF data from the image.
	x, err := exif.Decode(f)
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("Exif: Decode: %v", err)
	}

	// Fetch the GPS data if it exists
	lat, long, err = x.LatLong()
	if err != nil {
		return 0.0, 0.0, fmt.Errorf("Exif: Fetch Lat Long: %v", err)
	}
	return
}
