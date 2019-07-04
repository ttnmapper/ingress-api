package main

import (
	"errors"
	"strconv"
	"ttnmapper-ingress-api/types"
)

func ParsePayloadFields(packet *types.TtnMapperUplinkMessage) error {

	if packet.PayloadFields == nil {
		return errors.New("payload_fields not set")
	}

	if err := parseCayenneLpp(packet, packet.PayloadFields); err != nil {
		return err
	}

	locationKeys := [...]string{"location", "gps"}
	for _, v := range locationKeys {
		if val, ok := packet.PayloadFields[v]; ok {
			if locationObject, ok := val.(map[string]interface{}); ok {
				if err := extractFromRoot(packet, locationObject); err != nil {
					return err
				}
			}
		}
	}
	//if val, ok := packet.PayloadFields["location"]; ok {
	//	locationObject := val.(map[string]interface{})
	//	if err := extractFromRoot(packet, locationObject); err != nil {
	//		return err
	//	}
	//}

	if err := extractFromRoot(packet, packet.PayloadFields); err != nil {
		return err
	}

	return nil
}

func parseCayenneLpp(packet *types.TtnMapperUplinkMessage, data map[string]interface{}) error {

	// Cayenne LPP formats
	for i := 0; i < 10; i++ {
		key := "gps_" + strconv.Itoa(i)

		if val, ok := data[key]; ok {
			gpsObject := val.(map[string]interface{})

			if err := extractFromRoot(packet, gpsObject); err != nil {
				return err
			}

			packet.TtnMProvider = "Cayenne LPP"

		}
	}

	return nil
}

func extractFromRoot(packet *types.TtnMapperUplinkMessage, data map[string]interface{}) error {

	latitudeKeys := [...]string{"lat", "latitude", "Latitude", "latitudeDeg", "gps_lat"}
	for _, v := range latitudeKeys {
		if val, ok := data[v]; ok {

			switch val.(type) {
			case float64:
				packet.TtnMLatitude = val.(float64)

			case string:
				// value was sent as a string not as a number
				if f, err := strconv.ParseFloat(val.(string), 64); err == nil {
					packet.TtnMLatitude = f
				}

			default:
				// type can not be handled
			}
		}
	}

	longitudeKeys := [...]string{"lon", "lng", "long", "longitude", "Longitude", "longitudeDeg", "gps_lng"}
	for _, v := range longitudeKeys {
		if val, ok := data[v]; ok {

			switch val.(type) {
			case float64:
				packet.TtnMLongitude = val.(float64)

			case string:
				// value was sent as a string not as a number
				if f, err := strconv.ParseFloat(val.(string), 64); err == nil {
					packet.TtnMLongitude = f
				}

			default:
				// type can not be handled
			}
		}
	}

	altitudeKeys := [...]string{"alt", "altitude", "Altitude", "height", "gpsalt", "gps_alt"}
	for _, v := range altitudeKeys {
		if val, ok := data[v]; ok {

			switch val.(type) {
			case float64:
				packet.TtnMAltitude = val.(float64)

			case string:
				// value was sent as a string not as a number
				if f, err := strconv.ParseFloat(val.(string), 64); err == nil {
					packet.TtnMAltitude = f
				}

			default:
				// type can not be handled
			}
		}
	}

	/*
		1. sats
		2. accuracy
		3. hdop
		4. provider
		as last one overrides any earlier ones
	*/

	satelliteKeys := [...]string{"sats", "satellites", "numsat", "numsats"}
	for _, v := range satelliteKeys {
		if val, ok := data[v]; ok {
			packet.TtnMSatellites = int32(val.(float64))
			packet.TtnMProvider = v

			switch val.(type) {
			case float64:
				packet.TtnMSatellites = int32(val.(float64))

			case string:
				// value was sent as a string not as a number
				if i, err := strconv.ParseInt(val.(string), 10, 32); err == nil {
					packet.TtnMSatellites = int32(i)
				}

			default:
				// type can not be handled
			}
		}
	}

	accuracyKeys := [...]string{"acc", "accuracy", "hacc"}
	for _, v := range accuracyKeys {
		if val, ok := data[v]; ok {

			packet.TtnMProvider = v

			switch val.(type) {
			case float64:
				packet.TtnMAccuracy = val.(float64)

			case string:
				// value was sent as a string not as a number
				if f, err := strconv.ParseFloat(val.(string), 64); err == nil {
					packet.TtnMAccuracy = f
				}

			default:
				// type can not be handled
			}
		}
	}

	hdopKeys := [...]string{"hdop", "gps_hdop"}
	for _, v := range hdopKeys {
		if val, ok := data[v]; ok {

			packet.TtnMProvider = v

			switch val.(type) {
			case float64:
				packet.TtnMHdop = val.(float64)

			case string:
				// value was sent as a string not as a number
				if f, err := strconv.ParseFloat(val.(string), 64); err == nil {
					packet.TtnMHdop = f
				}

			default:
				// type can not be handled
			}
		}
	}

	// Digital Matter Oyster and Yabby reports a "FixFailed" flag to indicate that the provided coordinates are cached
	// ones and not valid for mapping purposes.
	fixFailedKeys := [...]string{"fixFailed"}
	for _, v := range fixFailedKeys {
		if val, ok := data[v]; ok {

			switch val.(type) {
			case bool:
				if val == true {
					return errors.New("fixFailed is true - ignoring measurement")
				}

			default:
				// type can not be handled
			}
		}
	}

	if val, ok := data["provider"]; ok {
		packet.TtnMProvider = val.(string)
	}

	return nil
}
