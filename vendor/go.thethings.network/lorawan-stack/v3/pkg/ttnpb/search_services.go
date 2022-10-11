// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

// GetEntityIdentifiers returns the CollaboratorOf field as EntityIdentifiers.
func (req *SearchAccountsRequest) GetEntityIdentifiers() *EntityIdentifiers {
	if req == nil || req.CollaboratorOf == nil {
		return nil
	}
	switch v := req.CollaboratorOf.(type) {
	default:
		return nil
	case *SearchAccountsRequest_ApplicationIds:
		return v.ApplicationIds.GetEntityIdentifiers()
	case *SearchAccountsRequest_ClientIds:
		return v.ClientIds.GetEntityIdentifiers()
	case *SearchAccountsRequest_GatewayIds:
		return v.GatewayIds.GetEntityIdentifiers()
	case *SearchAccountsRequest_OrganizationIds:
		return v.OrganizationIds.GetEntityIdentifiers()
	}
}
