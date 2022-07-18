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

// CheckCapacityForMetro200Response struct for CheckCapacityForMetro200Response
type CheckCapacityForMetro200Response struct {
	Servers []CheckCapacityForMetro200ResponseServersInner `json:"servers,omitempty"`
}

// NewCheckCapacityForMetro200Response instantiates a new CheckCapacityForMetro200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckCapacityForMetro200Response() *CheckCapacityForMetro200Response {
	this := CheckCapacityForMetro200Response{}
	return &this
}

// NewCheckCapacityForMetro200ResponseWithDefaults instantiates a new CheckCapacityForMetro200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckCapacityForMetro200ResponseWithDefaults() *CheckCapacityForMetro200Response {
	this := CheckCapacityForMetro200Response{}
	return &this
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *CheckCapacityForMetro200Response) GetServers() []CheckCapacityForMetro200ResponseServersInner {
	if o == nil || o.Servers == nil {
		var ret []CheckCapacityForMetro200ResponseServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCapacityForMetro200Response) GetServersOk() ([]CheckCapacityForMetro200ResponseServersInner, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *CheckCapacityForMetro200Response) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []CheckCapacityForMetro200ResponseServersInner and assigns it to the Servers field.
func (o *CheckCapacityForMetro200Response) SetServers(v []CheckCapacityForMetro200ResponseServersInner) {
	o.Servers = v
}

func (o CheckCapacityForMetro200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableCheckCapacityForMetro200Response struct {
	value *CheckCapacityForMetro200Response
	isSet bool
}

func (v NullableCheckCapacityForMetro200Response) Get() *CheckCapacityForMetro200Response {
	return v.value
}

func (v *NullableCheckCapacityForMetro200Response) Set(val *CheckCapacityForMetro200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckCapacityForMetro200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckCapacityForMetro200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckCapacityForMetro200Response(val *CheckCapacityForMetro200Response) *NullableCheckCapacityForMetro200Response {
	return &NullableCheckCapacityForMetro200Response{value: val, isSet: true}
}

func (v NullableCheckCapacityForMetro200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckCapacityForMetro200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
