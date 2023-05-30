package dms

import (
	"fmt"
)

// DMS(Degree, Minute, Seconds) coordinates
type DMS struct {
	Degrees   uint8
	Minutes   uint8
	Seconds   float64
	Direction string
}

func (d *DMS) String() string {
	return fmt.Sprintf(`%d° %d' %.3f" %s`, d.Degrees, d.Minutes, d.Seconds, d.Direction)
}

// NewDMS converts Decimal Degrees to Degree, Minute, Seconds coordinates
func NewDMS(latitude, longitude float64) (*DMS, *DMS, error) {
	var latDirection string
	var lonDirection string

	if latitude > 0 {
		latDirection = "N"
	} else {
		latDirection = "S"
		latitude = -latitude
	}

	if longitude > 0 {
		lonDirection = "E"
	} else {
		lonDirection = "W"
		longitude = -longitude
	}

	if latitude >= 90 || longitude >= 180 {
		return nil, nil, fmt.Errorf("Latitude must be less than 90 and longitude must be less than 180.")
	}

	latitudeDegrees := uint8(latitude)
	latitudeMinutes := uint8((latitude - float64(latitudeDegrees)) * 60)
	latitudeSeconds := (latitude - float64(latitudeDegrees) - float64(latitudeMinutes)/60) * 3600

	longitudeDegrees := uint8(longitude)
	longitudeMinutes := uint8((longitude - float64(longitudeDegrees)) * 60)
	longitudeSeconds := (longitude - float64(longitudeDegrees) - float64(longitudeMinutes)/60) * 3600

	return &DMS{
			Degrees:   latitudeDegrees,
			Minutes:   latitudeMinutes,
			Seconds:   latitudeSeconds,
			Direction: latDirection,
		},
		&DMS{
			Degrees:   longitudeDegrees,
			Minutes:   longitudeMinutes,
			Seconds:   longitudeSeconds,
			Direction: lonDirection,
		},
		nil
}
