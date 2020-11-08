package main

import (
	"errors"
	"ttnmapper-ingress-api/types"
)

func CheckData(packet types.TtnMapperUplinkMessage) error {

	// Accuracy too low when satellites less than 4
	if !IsZeroOfUnderlyingType(packet.Satellites) {
		if packet.Satellites < 4 {
			return errors.New("less than 4 satellites")
		}
	}

	// Accuracy value must be below 10 metre
	if !IsZeroOfUnderlyingType(packet.AccuracyMeters) {
		if packet.AccuracyMeters > 10 {
			return errors.New("accuracy too low")
		}
	}

	// HDOP must be lower than 5
	if !IsZeroOfUnderlyingType(packet.Hdop) {
		if packet.Hdop > 5 {
			return errors.New("hdop is too high")
		}
	}

	// Latitude
	if IsZeroOfUnderlyingType(packet.Latitude) {
		return errors.New("latitude not set")
	}
	if packet.Latitude >= 90 || packet.Latitude <= -90 {
		return errors.New("latitude out of range")
	}

	// Longitude
	if IsZeroOfUnderlyingType(packet.Longitude) {
		return errors.New("longitude not set")
	}
	if packet.Longitude >= 180 || packet.Longitude <= -180 {
		return errors.New("longitude out of range")
	}

	// Null island
	if packet.Longitude < 1 && packet.Longitude > -1 && packet.Latitude < 1 && packet.Latitude > -1 {
		return errors.New("not accepting coordinates on null island")
	}

	return nil
}

func SanitizeData(packet *types.TtnMapperUplinkMessage) {
	// clamp altitude to ground if not set
	if IsZeroOfUnderlyingType(packet.Altitude) {
		packet.Altitude = 0
	}

	// handle overflow for a small range just below 2^16
	if packet.Altitude > (2^16-1000) && packet.Altitude < 2^16 {
		packet.Altitude = packet.Altitude - (2 ^ 16) // Negative altitude
	}

	// round lat to 6 decimal places?
	// round lon to 6 decimal places?

	// Some single channel gateways send frequency in Hz, not MHz
	// TTNv3 also sends the frequency in Herz, not MHz like V2 - change to Hz here
	// Below 1MHz assume the value is passed in MHz not Hz, so convert to Hz
	if packet.Frequency < 1000000 {
		packet.Frequency = packet.Frequency * 1000000
	}
}
