// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnpb

import (
	"strings"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

// TopLevelFields returns the unique top level fields of the given paths.
func TopLevelFields(paths []string) []string {
	seen := make(map[string]struct{}, len(paths))
	out := make([]string, 0, len(paths))
	for _, path := range paths {
		parts := strings.SplitN(path, ".", 2)
		if _, ok := seen[parts[0]]; ok {
			continue
		}
		seen[parts[0]] = struct{}{}
		out = append(out, parts[0])
	}
	return out
}

// BottomLevelFields returns the unique bottom level fields of the given paths.
func BottomLevelFields(paths []string) []string {
	seen := make(map[string]struct{}, len(paths))
	for _, path := range paths {
		prefix := path
		if i := strings.LastIndex(prefix, "."); i >= 0 {
			prefix = prefix[:i]
		}
		if _, ok := seen[prefix]; ok {
			delete(seen, prefix)
		}
		seen[path] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	return out
}

// HasOnlyAllowedFields returns whether the given requested paths only contains paths that are allowed.
// The requested fields (i.e. `a.b`) may be of a lower level than the allowed path (i.e. `a`).
func HasOnlyAllowedFields(requested []string, allowed ...string) bool {
nextRequested:
	for _, requested := range requested {
		for _, allowed := range allowed {
			if requested == allowed || strings.HasPrefix(requested, allowed+".") {
				continue nextRequested
			}
		}
		return false
	}
	return true
}

// HasAnyField returns whether the given requested paths contain any of the given fields.
// The requested fields (i.e. `a.b`) may be of a higher level than the search path (i.e. `a.b.c`).
func HasAnyField(requested []string, search ...string) bool {
	for _, requested := range requested {
		for _, search := range search {
			if requested == search || strings.HasPrefix(search, requested+".") {
				return true
			}
		}
	}
	return false
}

// FlattenPaths flattens the paths by the given paths to flatten.
// When paths contains `a.b.c` and flatten contains `a.b`, the result will be `a.b`.
func FlattenPaths(paths, flatten []string) []string {
	res := make([]string, 0, len(paths))
	flattened := make(map[string]bool)
	for _, path := range paths {
		for _, flatten := range flatten {
			if flatten == path || strings.HasPrefix(path, flatten+".") {
				if !flattened[flatten] {
					res = append(res, flatten)
					flattened[flatten] = true
				}
			} else {
				res = append(res, path)
			}
		}
	}
	return res
}

var errMissingField = errors.Define("missing_field", "field `{field}` is missing")

// RequireFields returns nil if the given requested paths contain all of the given fields and error otherwise.
// The requested fields (i.e. `a.b`) may be of a higher level than the search path (i.e. `a.b.c`).
func RequireFields(requested []string, search ...string) error {
	for _, s := range search {
		if !HasAnyField(requested, s) {
			return errMissingField.WithAttributes("field", s)
		}
	}
	return nil
}

var errProhibitedField = errors.Define("prohibited_field", "field `{field}` is prohibited")

// ProhibitFields returns nil if the given requested paths contain none of the given fields and error otherwise.
// The requested fields (i.e. `a.b`) may be of a higher level than the search path (i.e. `a.b.c`).
func ProhibitFields(requested []string, search ...string) error {
	for _, s := range search {
		if HasAnyField(requested, s) {
			return errProhibitedField.WithAttributes("field", s)
		}
	}
	return nil
}

// ContainsField returns true if the given paths contains the field path.
func ContainsField(path string, allowedPaths []string) bool {
	for _, allowedPath := range allowedPaths {
		if path == allowedPath {
			return true
		}
	}
	return false
}

// AllowedFields returns the paths from the given paths that are in the allowed paths.
func AllowedFields(paths, allowedPaths []string) []string {
	selectedPaths := make([]string, 0, len(paths))
	for _, path := range paths {
		if ContainsField(path, allowedPaths) {
			selectedPaths = append(selectedPaths, path)
			continue
		}
	}
	return selectedPaths
}

// AllowedBottomLevelFields returns the bottom level paths from the given paths that are in the allowed paths.
func AllowedBottomLevelFields(paths, allowedPaths []string) []string {
	allowedPaths = BottomLevelFields(allowedPaths)
	selectedPaths := make([]string, 0, len(allowedPaths))
outer:
	for _, allowedPath := range allowedPaths {
		for _, path := range paths {
			if allowedPath == path || strings.HasPrefix(allowedPath, path+".") {
				selectedPaths = append(selectedPaths, allowedPath)
				continue outer
			}
		}
	}
	return selectedPaths
}

// ExcludeFields returns the given paths without the given search paths to exclude.
func ExcludeFields(paths []string, excludePaths ...string) []string {
	if len(paths) == 0 {
		return paths
	}
	excluded := make([]string, 0, len(paths))
outer:
	for _, path := range paths {
		for _, excludePath := range excludePaths {
			if path == excludePath || strings.HasPrefix(path, excludePath+".") {
				continue outer
			}
		}
		excluded = append(excluded, path)
	}
	return excluded
}

// ExcludeSubFields returns the given paths without sub-fields of excludePaths.
func ExcludeSubFields(paths []string, excludePaths ...string) []string {
	if len(paths) == 0 {
		return paths
	}
	excluded := make([]string, 0, len(paths))
outer:
	for _, path := range paths {
		for _, excludePath := range excludePaths {
			if strings.HasPrefix(path, excludePath+".") {
				continue outer
			}
		}
		excluded = append(excluded, path)
	}
	return excluded
}

// AddFields appends all addPaths which are not already present in paths to paths and returns the result.
func AddFields(paths []string, addPaths ...string) []string {
	for _, p := range addPaths {
		if !HasAnyField(paths, p) {
			paths = append(paths, p)
		}
	}
	return ExcludeSubFields(paths, addPaths...)
}

// FieldsWithPrefix returns the paths with each the prefix prepended.
func FieldsWithPrefix(prefix string, paths ...string) []string {
	ret := make([]string, 0, len(paths))
	for _, p := range paths {
		ret = append(ret, prefix+"."+p)
	}
	return ret
}

// FieldsWithoutPrefix returns the paths under prefix with said prefix removed.
func FieldsWithoutPrefix(prefix string, paths ...string) []string {
	ret := make([]string, 0, len(paths))
	for _, p := range paths {
		if i := strings.Index(p, prefix+"."); i >= 0 {
			ret = append(ret, p[1+len(prefix):])
		}
	}
	return ret
}

// AddImplicitEndDeviceGetFields appends implicit EndDevice get paths to paths if not already present.
func AddImplicitEndDeviceGetFields(paths ...string) []string {
	return AddFields(paths,
		"created_at",
		"ids",
		"updated_at",
	)
}

// ApplyEndDeviceFieldMask applies fields specified by paths from src to dst and returns the result.
// If dst is nil, a new EndDevice is created.
func ApplyEndDeviceFieldMask(dst, src *EndDevice, paths ...string) (*EndDevice, error) {
	if dst == nil {
		dst = &EndDevice{}
	}
	return dst, dst.SetFields(src, paths...)
}

// FilterGetEndDevice returns a new EndDevice with only implicit fields and the ones specified by paths set.
func FilterGetEndDevice(pb *EndDevice, paths ...string) (*EndDevice, error) {
	return ApplyEndDeviceFieldMask(nil, pb, AddImplicitEndDeviceGetFields(paths...)...)
}

// AddImplicitSessionKeysGetFields appends implicit SessionKeys get paths to paths if not already present.
func AddImplicitSessionKeysGetFields(paths ...string) []string {
	return AddFields(paths,
		"session_key_id",
	)
}

// ApplySessionKeysFieldMask applies fields specified by paths from src to dst and returns the result.
// If dst is nil, a new SessionKeys is created.
func ApplySessionKeysFieldMask(dst, src *SessionKeys, paths ...string) (*SessionKeys, error) {
	if dst == nil {
		dst = &SessionKeys{}
	}
	return dst, dst.SetFields(src, paths...)
}

// FilterGetSessionKeys returns a new SessionKeys with only implicit fields and the ones specified by paths set.
func FilterGetSessionKeys(pb *SessionKeys, paths ...string) (*SessionKeys, error) {
	return ApplySessionKeysFieldMask(nil, pb, AddImplicitSessionKeysGetFields(paths...)...)
}
