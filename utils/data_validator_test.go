package utils

import (
	"testing"
	"ttnmapper-ingress-api/types"
)

func TestSanitizeFrequency(t *testing.T) {
	if SanitizeFrequency(868000000) != 868000000 {
		t.Fatalf("868000000")
	}

	if SanitizeFrequency(868.000) != 868000000 {
		t.Fatalf("868.000")
	}

	if SanitizeFrequency(868400000000000) != 868400000 {
		t.Fatalf("868400000000000")
	}

	if SanitizeFrequency(869099976) != 869100000 {
		t.Fatalf("869099976")
	}
}

func TestCheckData(t *testing.T) {
	packet := types.TtnMapperUplinkMessage{
		Latitude:       0,
		Longitude:      0,
		Altitude:       0,
		AccuracyMeters: 0,
		Satellites:     0,
		Hdop:           0,
	}

	err := CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect empty coordinates")
	}

	packet.Latitude = 91
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect latitude out of range")
	}

	packet.Latitude = -91
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect latitude out of range")
	}

	packet.Latitude = 89
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect longitude not set")
	}

	packet.Longitude = 181
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect longitude not set")
	}

	packet.Longitude = -181
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect longitude not set")
	}

	packet.Longitude = 179
	packet.Satellites = 3
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect too few satellites")
	}
	packet.Satellites = 0

	packet.AccuracyMeters = 11
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect too high accuracy meters")
	}
	packet.AccuracyMeters = 0

	packet.Hdop = 6
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect too high hdop")
	}
	packet.Hdop = 0

	packet.Latitude = 0.5
	packet.Longitude = 0.5
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect null island")
	}

	packet.Latitude = -0.5
	packet.Longitude = -0.5
	err = CheckData(packet)
	if err == nil {
		t.Fatalf("didn't detect null island")
	}
}

func TestSanitizeData(t *testing.T) {
	packet := types.TtnMapperUplinkMessage{
		Latitude:       0,
		Longitude:      0,
		Altitude:       0,
		AccuracyMeters: 0,
		Satellites:     0,
		Hdop:           0,
	}

	SanitizeData(&packet)
	if packet.Latitude != 0 || packet.Longitude != 0 ||
		packet.Altitude != 0 || packet.AccuracyMeters != 0 ||
		packet.Satellites != 0 || packet.Hdop != 0 {
		t.Fatalf("zero values changed to non zero")
	}

	packet.Latitude = 33.3333
	SanitizeData(&packet)
	if packet.Latitude != 33.3333 {
		t.Fatalf("valid latitude changed")
	}

	packet.Latitude = -33.3333
	SanitizeData(&packet)
	if packet.Latitude != -33.3333 {
		t.Fatalf("valid latitude changed")
	}

	packet.Longitude = 18.00234
	SanitizeData(&packet)
	if packet.Longitude != 18.00234 {
		t.Fatalf("valid longitude changed")
	}

	packet.Longitude = -18.00234
	SanitizeData(&packet)
	if packet.Longitude != -18.00234 {
		t.Fatalf("valid longitude changed")
	}

	packet.Latitude = 91
	SanitizeData(&packet)
	if packet.Latitude != 0 {
		t.Fatalf("out of range latitude not set to 0")
	}

	packet.Latitude = -91
	SanitizeData(&packet)
	if packet.Latitude != 0 {
		t.Fatalf("out of range latitude not set to 0")
	}

	packet.Longitude = 181
	SanitizeData(&packet)
	if packet.Longitude != 0 {
		t.Fatalf("out of range longitude not set to 0")
	}

	packet.Longitude = -181
	SanitizeData(&packet)
	if packet.Longitude != 0 {
		t.Fatalf("out of range longitude not set to 0")
	}

	packet.Latitude = 0.5
	SanitizeData(&packet)
	if packet.Latitude != 0 {
		t.Fatalf("null island latitude not set to 0")
	}

	packet.Latitude = -0.5
	SanitizeData(&packet)
	if packet.Latitude != 0 {
		t.Fatalf("null island latitude not set to 0")
	}

	packet.Longitude = 0.5
	SanitizeData(&packet)
	if packet.Longitude != 0 {
		t.Fatalf("null island longitude not set to 0")
	}

	packet.Longitude = -0.5
	SanitizeData(&packet)
	if packet.Longitude != 0 {
		t.Fatalf("null island longitude not set to 0")
	}

	packet.Altitude = -300
	SanitizeData(&packet)
	if packet.Altitude != -300 {
		t.Fatalf("small negative altitude changed")
	}

	packet.Altitude = 65530
	SanitizeData(&packet)
	if packet.Altitude != -6 {
		t.Fatalf("16 bit overflow not handled: %f", packet.Altitude)
	}

	packet.Altitude = 65540
	SanitizeData(&packet)
	if packet.Altitude != 65540 {
		t.Fatalf("16 bit overflow not handled correctly: %f", packet.Altitude)
	}

	packet.Altitude = 64536 // 64537 will become -999
	SanitizeData(&packet)
	if packet.Altitude != 64536 {
		t.Fatalf("16 bit overflow not handled correctly: %f", packet.Altitude)
	}
}
