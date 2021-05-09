// Copyright © 2017 The Things Network
// Use of this source code is governed by the MIT license that can be found in the LICENSE file.

package ttn_types

// GatewayMetadata contains metadata for each gateway that received a message
type GatewayMetadata struct {
	GtwID                  string   `json:"gtw_id,omitempty"`
	GtwTrusted             bool     `json:"gtw_trusted,omitempty"`
	Timestamp              uint32   `json:"timestamp,omitempty"`
	FineTimestamp          uint64   `json:"fine_timestamp,omitempty"`
	FineTimestampEncrypted []byte   `json:"fine_timestamp_encrypted,omitempty"`
	Time                   JSONTime `json:"time,omitempty"`
	Antenna                uint8    `json:"antenna,omitempty"`
	Channel                uint32   `json:"channel"`
	RSSI                   float32  `json:"rssi"`
	SNR                    float32  `json:"snr"`
	RFChain                uint32   `json:"rf_chain"`
	LocationMetadata
}
