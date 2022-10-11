// Copyright © 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

package flagsplugin

import "github.com/spf13/pflag"

// AddAlias creates a flag with different name that points to another flag.
// Setting this flag will update the value in the child flag.
func AddAlias(flagSet *pflag.FlagSet, flagName, aliasName string, opts ...FlagOption) {
	flag := flagSet.Lookup(flagName)

	alias := *flag
	alias.Name = aliasName
	alias.Value = &AliasValue{target: flag}
	alias.Usage = flag.Usage
	ApplyOptions(&alias, opts...)
	flagSet.AddFlag(&alias)
}

// AliasValue holds a pointer to the target flag.
type AliasValue struct {
	target *pflag.Flag
}

// Set implements pflag.Value interface.
func (av *AliasValue) Set(s string) error {
	err := av.target.Value.Set(s)
	if err != nil {
		return err
	}
	av.target.Changed = true
	return nil
}

// Type implements pflag.Value interface.
func (av *AliasValue) Type() string { return av.target.Value.Type() }

// String implements pflag.Value interface.
func (av *AliasValue) String() string { return av.target.Value.String() }
