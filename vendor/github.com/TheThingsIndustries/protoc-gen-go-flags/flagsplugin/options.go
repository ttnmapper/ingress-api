// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import "github.com/spf13/pflag"

// FlagOption defines a function that operates on a flag object.
type FlagOption func(*pflag.Flag)

// ApplyOptions applies flag options to the passed flag.
func ApplyOptions(f *pflag.Flag, opts ...FlagOption) {
	for _, opt := range opts {
		opt(f)
	}
}

// WithHidden hides flag from help description.
func WithHidden(hidden bool) FlagOption {
	return func(f *pflag.Flag) {
		f.Hidden = hidden
	}
}
