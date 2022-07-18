/*
Metal API

This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// UpdateMembershipRequest struct for UpdateMembershipRequest
type UpdateMembershipRequest struct {
	Role []string `json:"role,omitempty"`
}

// NewUpdateMembershipRequest instantiates a new UpdateMembershipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMembershipRequest() *UpdateMembershipRequest {
	this := UpdateMembershipRequest{}
	return &this
}

// NewUpdateMembershipRequestWithDefaults instantiates a new UpdateMembershipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMembershipRequestWithDefaults() *UpdateMembershipRequest {
	this := UpdateMembershipRequest{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UpdateMembershipRequest) GetRole() []string {
	if o == nil || o.Role == nil {
		var ret []string
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMembershipRequest) GetRoleOk() ([]string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateMembershipRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given []string and assigns it to the Role field.
func (o *UpdateMembershipRequest) SetRole(v []string) {
	o.Role = v
}

func (o UpdateMembershipRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateMembershipRequest struct {
	value *UpdateMembershipRequest
	isSet bool
}

func (v NullableUpdateMembershipRequest) Get() *UpdateMembershipRequest {
	return v.value
}

func (v *NullableUpdateMembershipRequest) Set(val *UpdateMembershipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMembershipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMembershipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMembershipRequest(val *UpdateMembershipRequest) *NullableUpdateMembershipRequest {
	return &NullableUpdateMembershipRequest{value: val, isSet: true}
}

func (v NullableUpdateMembershipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMembershipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
