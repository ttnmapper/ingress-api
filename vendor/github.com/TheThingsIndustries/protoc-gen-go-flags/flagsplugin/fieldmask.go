// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import "strings"

var (
	toDash       = strings.NewReplacer("_", "-")
	toUnderscore = strings.NewReplacer("-", "_")
)
