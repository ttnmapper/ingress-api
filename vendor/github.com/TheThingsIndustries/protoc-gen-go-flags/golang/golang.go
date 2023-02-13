// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package golangplugin

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SetDuration converts time to durationpb.Duration
func SetDuration(t time.Duration) *durationpb.Duration {
	return durationpb.New(t)
}

// SetFieldMask converst string slice to fieldmaskpb.FieldMask
func SetFieldMask(paths []string) *fieldmaskpb.FieldMask {
	return &fieldmaskpb.FieldMask{Paths: paths}
}

// SetTimestamp converts time to timestamppb.Timestamp
func SetTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
