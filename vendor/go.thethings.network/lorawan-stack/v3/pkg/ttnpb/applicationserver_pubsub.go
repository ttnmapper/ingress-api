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

// ApplicationPubSub_Provider is an alias to the interface identifying the PubSub provider types.
// This enables provider.RegisterProvider and provider.GetProvider to offer type safety guarantees.
// The underscore is maintained for consistency with the generated code.
type ApplicationPubSub_Provider = isApplicationPubSub_Provider

// All EntityType methods implement the IDStringer interface.

func (m *ApplicationPubSubIdentifiers) EntityType() string {
	return m.GetApplicationIds().EntityType()
}

func (m *ApplicationPubSub) EntityType() string {
	return m.GetIds().EntityType()
}

func (m *GetApplicationPubSubRequest) EntityType() string {
	return m.GetIds().EntityType()
}

func (m *ListApplicationPubSubsRequest) EntityType() string {
	return m.GetApplicationIds().EntityType()
}

func (m *SetApplicationPubSubRequest) EntityType() string {
	return m.GetPubsub().EntityType()
}

// All IDString methods implement the IDStringer interface.

func (m *ApplicationPubSubIdentifiers) IDString() string {
	return m.GetApplicationIds().IDString()
}

func (m *ApplicationPubSub) IDString() string {
	return m.GetIds().IDString()
}

func (m *GetApplicationPubSubRequest) IDString() string {
	return m.GetIds().IDString()
}

func (m *ListApplicationPubSubsRequest) IDString() string {
	return m.GetApplicationIds().IDString()
}

func (m *SetApplicationPubSubRequest) IDString() string {
	return m.GetPubsub().IDString()
}

// All ExtractRequestFields methods are used by github.com/grpc-ecosystem/go-grpc-middleware/tags.

func (m *ApplicationPubSubIdentifiers) ExtractRequestFields(dst map[string]interface{}) {
	m.GetApplicationIds().ExtractRequestFields(dst)
}

func (m *ApplicationPubSub) ExtractRequestFields(dst map[string]interface{}) {
	m.GetIds().ExtractRequestFields(dst)
}

func (m *GetApplicationPubSubRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetIds().ExtractRequestFields(dst)
}

func (m *ListApplicationPubSubsRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetApplicationIds().ExtractRequestFields(dst)
}

func (m *SetApplicationPubSubRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetPubsub().ExtractRequestFields(dst)
}
