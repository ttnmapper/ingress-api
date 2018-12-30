package main

import (
	"errors"
	"log"
	"ttnmapper-ingress-api/types"
)

func CheckData(packet types.TtnMapperUplinkMessage) error {

	// Accuracy too low when satellites less than 4
	if !IsZeroOfUnderlyingType(packet.TtnMSatellites) {
		if packet.TtnMSatellites < 4 {
			return errors.New("less than 4 satellites")
		}
	}

	// Accuracy value must be below 10 metre
	if !IsZeroOfUnderlyingType(packet.TtnMAccuracy) {
		if packet.TtnMAccuracy > 10 {
			return errors.New("accuracy too low")
		}
	}

	// HDOP must be lower than 5
	if !IsZeroOfUnderlyingType(packet.TtnMHdop) {
		if packet.TtnMHdop > 5 {
			return errors.New("hdop is too high")
		}
	}

	// Latitude
	if IsZeroOfUnderlyingType(packet.TtnMLatitude) {
		log.Print(packet.TtnMLatitude)
		return errors.New("latitude not set")
	}

	if packet.TtnMLatitude >= 90 || packet.TtnMLatitude <= -90 || packet.TtnMLatitude == 0 {
		return errors.New("latitude out of range")
	}

	// Longitude
	if IsZeroOfUnderlyingType(packet.TtnMLongitude) {
		return errors.New("longitude not set")
	}

	if packet.TtnMLongitude >= 180 || packet.TtnMLongitude <= -180 || packet.TtnMLongitude == 0 {
		return errors.New("longitude out of range")
	}

	// Null island
	if packet.TtnMLongitude < 1 && packet.TtnMLongitude > -1 && packet.TtnMLatitude < 1 && packet.TtnMLatitude > -1 {
		return errors.New("not accepting coordinates on null island")
	}

	return nil
}

func SanitizeData(packet types.TtnMapperUplinkMessage) {
	// clamp altitude to ground if not set
	if IsZeroOfUnderlyingType(packet.TtnMAltitude) {
		packet.TtnMAltitude = 0
	}

	// handle overflow for a small range just below 2^16
	if packet.TtnMAltitude > 2^16 && packet.TtnMAltitude < (2^16-1000) {
		packet.TtnMAltitude = (2 ^ 16) - packet.TtnMAltitude
	}

	// round lat to 6 decimal places
	// round lon to 6 decimal places
}
