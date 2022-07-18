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

// UpdateCurrentUserRequest struct for UpdateCurrentUserRequest
type UpdateCurrentUserRequest struct {
	Customdata  map[string]interface{} `json:"customdata,omitempty"`
	FirstName   *string                `json:"first_name,omitempty"`
	LastName    *string                `json:"last_name,omitempty"`
	Password    *string                `json:"password,omitempty"`
	PhoneNumber *string                `json:"phone_number,omitempty"`
	Timezone    *string                `json:"timezone,omitempty"`
}

// NewUpdateCurrentUserRequest instantiates a new UpdateCurrentUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCurrentUserRequest() *UpdateCurrentUserRequest {
	this := UpdateCurrentUserRequest{}
	return &this
}

// NewUpdateCurrentUserRequestWithDefaults instantiates a new UpdateCurrentUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCurrentUserRequestWithDefaults() *UpdateCurrentUserRequest {
	this := UpdateCurrentUserRequest{}
	return &this
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *UpdateCurrentUserRequest) GetCustomdata() map[string]interface{} {
	if o == nil || o.Customdata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCurrentUserRequest) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || o.Customdata == nil {
		return nil, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *UpdateCurrentUserRequest) HasCustomdata() bool {
	if o != nil && o.Customdata != nil {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *UpdateCurrentUserRequest) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateCurrentUserRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCurrentUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UpdateCurrentUserRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateCurrentUserRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateCurrentUserRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCurrentUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UpdateCurrentUserRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateCurrentUserRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateCurrentUserRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCurrentUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateCurrentUserRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateCurrentUserRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *UpdateCurrentUserRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCurrentUserRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *UpdateCurrentUserRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *UpdateCurrentUserRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UpdateCurrentUserRequest) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCurrentUserRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UpdateCurrentUserRequest) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UpdateCurrentUserRequest) SetTimezone(v string) {
	o.Timezone = &v
}

func (o UpdateCurrentUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Customdata != nil {
		toSerialize["customdata"] = o.Customdata
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCurrentUserRequest struct {
	value *UpdateCurrentUserRequest
	isSet bool
}

func (v NullableUpdateCurrentUserRequest) Get() *UpdateCurrentUserRequest {
	return v.value
}

func (v *NullableUpdateCurrentUserRequest) Set(val *UpdateCurrentUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCurrentUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCurrentUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCurrentUserRequest(val *UpdateCurrentUserRequest) *NullableUpdateCurrentUserRequest {
	return &NullableUpdateCurrentUserRequest{value: val, isSet: true}
}

func (v NullableUpdateCurrentUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCurrentUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
