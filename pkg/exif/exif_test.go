package exif

import (
	"math"
	"testing"
)

func TestGetLatLong(t *testing.T) {
	// The path to a test image file with known lat long
	path := "./image/bird.jpeg"

	// The known latitude and longitude of the image
	expectedLat := 48.564388888888885
	expectedLong := -123.43914166666667

	// Call the function to test
	lat, long, err := GetLatLong(path)

	// Check if there was an error
	if err != nil {
		t.Errorf("GetLatLong returned an error: %v", err)
	}

	// Check the latitude and longitude values.
	// We use a small tolerance to account for precision differences
	tolerance := 0.000001
	if math.Abs(lat-expectedLat) > tolerance || math.Abs(long-expectedLong) > tolerance {
		t.Errorf("GetLatLong returned incorrect data: got %v and %v for latitude and longitude respectively, want %v and %v",
			lat, long, expectedLat, expectedLong)
	}
}
