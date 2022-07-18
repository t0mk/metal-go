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

// FindOrganizationTransfers200Response struct for FindOrganizationTransfers200Response
type FindOrganizationTransfers200Response struct {
	Transfers []FindOrganizationTransfers200ResponseTransfersInner `json:"transfers,omitempty"`
}

// NewFindOrganizationTransfers200Response instantiates a new FindOrganizationTransfers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindOrganizationTransfers200Response() *FindOrganizationTransfers200Response {
	this := FindOrganizationTransfers200Response{}
	return &this
}

// NewFindOrganizationTransfers200ResponseWithDefaults instantiates a new FindOrganizationTransfers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindOrganizationTransfers200ResponseWithDefaults() *FindOrganizationTransfers200Response {
	this := FindOrganizationTransfers200Response{}
	return &this
}

// GetTransfers returns the Transfers field value if set, zero value otherwise.
func (o *FindOrganizationTransfers200Response) GetTransfers() []FindOrganizationTransfers200ResponseTransfersInner {
	if o == nil || o.Transfers == nil {
		var ret []FindOrganizationTransfers200ResponseTransfersInner
		return ret
	}
	return o.Transfers
}

// GetTransfersOk returns a tuple with the Transfers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindOrganizationTransfers200Response) GetTransfersOk() ([]FindOrganizationTransfers200ResponseTransfersInner, bool) {
	if o == nil || o.Transfers == nil {
		return nil, false
	}
	return o.Transfers, true
}

// HasTransfers returns a boolean if a field has been set.
func (o *FindOrganizationTransfers200Response) HasTransfers() bool {
	if o != nil && o.Transfers != nil {
		return true
	}

	return false
}

// SetTransfers gets a reference to the given []FindOrganizationTransfers200ResponseTransfersInner and assigns it to the Transfers field.
func (o *FindOrganizationTransfers200Response) SetTransfers(v []FindOrganizationTransfers200ResponseTransfersInner) {
	o.Transfers = v
}

func (o FindOrganizationTransfers200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Transfers != nil {
		toSerialize["transfers"] = o.Transfers
	}
	return json.Marshal(toSerialize)
}

type NullableFindOrganizationTransfers200Response struct {
	value *FindOrganizationTransfers200Response
	isSet bool
}

func (v NullableFindOrganizationTransfers200Response) Get() *FindOrganizationTransfers200Response {
	return v.value
}

func (v *NullableFindOrganizationTransfers200Response) Set(val *FindOrganizationTransfers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindOrganizationTransfers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindOrganizationTransfers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindOrganizationTransfers200Response(val *FindOrganizationTransfers200Response) *NullableFindOrganizationTransfers200Response {
	return &NullableFindOrganizationTransfers200Response{value: val, isSet: true}
}

func (v NullableFindOrganizationTransfers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindOrganizationTransfers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
