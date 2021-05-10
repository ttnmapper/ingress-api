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

package errors

import (
	"fmt"
	"regexp"

	"github.com/gotnospirit/messageformat"
	"go.thethings.network/lorawan-stack/v3/pkg/i18n"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Definition of a registered error.
type Definition struct {
	namespace              string
	name                   string
	messageFormat          string
	messageFormatArguments []string
	parsedMessageFormat    *messageformat.MessageFormat
	publicAttributes       []string
	code                   uint32 // 0 is invalid; so implies Unknown (code 2)
	grpcStatus             *status.Status
}

// DefinitionInterface is the interface of an error definition.
type DefinitionInterface interface {
	error
	fmt.Stringer
	Namespace() string
	Name() string
	FullName() string
	MessageFormat() string
	Code() uint32
}

// Namespace of the error.
func (d Definition) Namespace() string { return d.namespace }

// Name of the error.
func (d Definition) Name() string { return d.name }

// FullName returns the full name (namespace:name) of the error.
func (d Definition) FullName() string {
	namespace := d.namespace
	if namespace == "" {
		namespace = "unknown"
	}
	name := d.name
	if name == "" {
		name = "unknown"
	}
	return fmt.Sprintf("%s:%s", namespace, name)
}

// MessageFormat of the error.
func (d Definition) MessageFormat() string { return d.messageFormat }

// Code of the error.
// This code is consistent with google.golang.org/genproto/googleapis/rpc/code and google.golang.org/grpc/codes.
func (d Definition) Code() uint32 { return d.code }

func (d Definition) String() string {
	return fmt.Sprintf("error:%s (%s)", d.FullName(), d.messageFormat)
}

// Error implements the error interface.
func (d Definition) Error() string { return d.String() }

var messageFormatArgument = regexp.MustCompile(`\{[\s]*([a-z0-9_]+)`)

func messageFormatArguments(messageFormat string) (args []string) {
	for _, matches := range messageFormatArgument.FindAllStringSubmatch(messageFormat, -1) {
		if len(matches) == 2 {
			args = append(args, matches[1])
		}
	}
	m := make(map[string]struct{}, len(args))
	for _, arg := range args {
		m[arg] = struct{}{}
	}
	args = make([]string, 0, len(m))
	for arg := range m {
		args = append(args, arg)
	}
	return
}

func define(code uint32, name, messageFormat string, publicAttributes ...string) Definition {
	ns := namespace(3)
	if code == 0 {
		code = uint32(codes.Unknown)
	}

	def := Definition{
		namespace:              ns,
		name:                   name,
		messageFormat:          messageFormat,
		messageFormatArguments: messageFormatArguments(messageFormat),
		publicAttributes:       publicAttributes,
		code:                   code,
	}

	fullName := def.FullName()
	if Definitions[fullName] != nil {
		panic(fmt.Errorf("error %s already defined", fullName))
	}

	parsedMessageFormat, err := formatter.Parse(messageFormat)
	if err != nil {
		panic(fmt.Errorf("could not parse message format `%s` for %s: %s", messageFormat, fullName, err))
	}
	def.parsedMessageFormat = parsedMessageFormat

	// All message format arguments must be public:
nextArg:
	for _, arg := range def.messageFormatArguments {
		for _, attr := range def.publicAttributes {
			if arg == attr {
				continue nextArg
			}
		}
		def.publicAttributes = append(def.publicAttributes, arg)
	}

	def.setGRPCStatus() // store the (marshaled) gRPC status message.

	Definitions[fullName] = &def

	desc := i18n.Define(fmt.Sprintf("error:%s", fullName), def.messageFormat)
	desc.SetSource(2)

	return def
}

// Definitions of registered errors.
// Errors that are defined in init() funcs will be collected for translation.
var Definitions = make(map[string]*Definition)

// Define defines a registered error of type Unknown.
func Define(name, messageFormat string, publicAttributes ...string) Definition {
	return define(uint32(codes.Unknown), name, messageFormat, publicAttributes...)
}

// DefineInvalidArgument defines a registered error of type InvalidArgument.
func DefineInvalidArgument(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.InvalidArgument), name, messageFormat, publicAttributes...)
	return def
}

// DefineDeadlineExceeded defines a registered error of type DeadlineExceeded.
func DefineDeadlineExceeded(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.DeadlineExceeded), name, messageFormat, publicAttributes...)
	return def
}

// DefineCanceled defines a registered error of type Canceled.
func DefineCanceled(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.Canceled), name, messageFormat, publicAttributes...)
	return def
}

// DefineNotFound defines a registered error of type NotFound.
func DefineNotFound(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.NotFound), name, messageFormat, publicAttributes...)
	return def
}

// DefineAlreadyExists defines a registered error of type AlreadyExists.
func DefineAlreadyExists(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.AlreadyExists), name, messageFormat, publicAttributes...)
	return def
}

// DefinePermissionDenied defines a registered error of type PermissionDenied.
//
// It should be used when a client attempts to perform an authorized action
// using incorrect credentials or credentials with insufficient rights.
// If the client attempts to perform the action without providing any form
// of authentication, Unauthenticated should be used instead.
func DefinePermissionDenied(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.PermissionDenied), name, messageFormat, publicAttributes...)
	return def
}

// DefineResourceExhausted defines a registered error of type ResourceExhausted.
func DefineResourceExhausted(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.ResourceExhausted), name, messageFormat, publicAttributes...)
	return def
}

// DefineFailedPrecondition defines a registered error of type FailedPrecondition.
// Use Unavailable if the client can retry just the failing call.
// Use Aborted if the client should retry at a higher-level.
func DefineFailedPrecondition(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.FailedPrecondition), name, messageFormat, publicAttributes...)
	return def
}

// DefineAborted defines a registered error of type Aborted.
func DefineAborted(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.Aborted), name, messageFormat, publicAttributes...)
	return def
}

// OutOfRange - not used for now

// Unimplemented defines a registered error of type Unimplemented.
func DefineUnimplemented(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.Unimplemented), name, messageFormat, publicAttributes...)
	return def
}

// DefineInternal defines a registered error of type Internal.
func DefineInternal(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.Internal), name, messageFormat, publicAttributes...)
	return def
}

// DefineUnavailable defines a registered error of type Unavailable.
func DefineUnavailable(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.Unavailable), name, messageFormat, publicAttributes...)
	return def
}

// DefineDataLoss defines a registered error of type DataLoss.
func DefineDataLoss(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.DataLoss), name, messageFormat, publicAttributes...)
	return def
}

// DefineCorruption is the same as DefineDataLoss.
func DefineCorruption(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.DataLoss), name, messageFormat, publicAttributes...)
	return def
}

// DefineUnauthenticated defines a registered error of type Unauthenticated.
// It should be used when a client attempts to perform an authenticated action
// without providing any form of authentication.
// If the client attempts to perform the action using incorrect credentials
// or credentials with insufficient rights, PermissionDenied should be used instead.
func DefineUnauthenticated(name, messageFormat string, publicAttributes ...string) Definition {
	def := define(uint32(codes.Unauthenticated), name, messageFormat, publicAttributes...)
	return def
}

// New returns a new error from the definition. This is not required, but will
// add a stack trace for improved debugging.
func (d Definition) New() Error {
	return build(d, 0) // Don't refactor this to build(...).WithCause(...)
}
