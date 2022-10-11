// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package gogoplugin

import (
	"time"

	"github.com/gogo/protobuf/types"
)

// SetDuration converts time to *types.Duration
func SetDuration(t time.Duration) *types.Duration {
	return types.DurationProto(t)
}

// SetFieldMask converts time to *types.FieldMask
func SetFieldMask(paths []string) *types.FieldMask {
	return &types.FieldMask{Paths: paths}
}

// SetTimestamp converts time to *types.Timestamp
func SetTimestamp(t time.Time) *types.Timestamp {
	timestamp, err := types.TimestampProto(t)
	if err != nil {
		return nil
	}
	return timestamp
}
