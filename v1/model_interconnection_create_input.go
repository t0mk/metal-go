/*
 * Metal API
 *
 * This is the API for Equinix Metal Product. Interact with your devices, user account, and projects.
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// InterconnectionCreateInput struct for InterconnectionCreateInput
type InterconnectionCreateInput struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// A Facility ID or code.
	Facility string `json:"facility"`
	// A Metro ID or code. Required for creating a connection, unless creating with facility.
	Metro *string `json:"metro,omitempty"`
	// Either 'shared' or 'dedicated'.
	Type string `json:"type"`
	// Either 'primary' or 'redundant'.
	Redundancy string `json:"redundancy"`
	ContactEmail *string `json:"contact_email,omitempty"`
	Project *string `json:"project,omitempty"`
	// A connection speed, in bps, mbps, or gbps. Ex: '100000000' or '100 mbps'.
	Speed *string `json:"speed,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
}

// NewInterconnectionCreateInput instantiates a new InterconnectionCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterconnectionCreateInput(name string, facility string, type_ string, redundancy string) *InterconnectionCreateInput {
	this := InterconnectionCreateInput{}
	this.Name = name
	this.Facility = facility
	this.Type = type_
	this.Redundancy = redundancy
	return &this
}

// NewInterconnectionCreateInputWithDefaults instantiates a new InterconnectionCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterconnectionCreateInputWithDefaults() *InterconnectionCreateInput {
	this := InterconnectionCreateInput{}
	return &this
}

// GetName returns the Name field value
func (o *InterconnectionCreateInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InterconnectionCreateInput) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterconnectionCreateInput) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterconnectionCreateInput) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterconnectionCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetFacility returns the Facility field value
func (o *InterconnectionCreateInput) GetFacility() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetFacilityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Facility, true
}

// SetFacility sets field value
func (o *InterconnectionCreateInput) SetFacility(v string) {
	o.Facility = v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *InterconnectionCreateInput) GetMetro() string {
	if o == nil || o.Metro == nil {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetMetroOk() (*string, bool) {
	if o == nil || o.Metro == nil {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *InterconnectionCreateInput) HasMetro() bool {
	if o != nil && o.Metro != nil {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *InterconnectionCreateInput) SetMetro(v string) {
	o.Metro = &v
}

// GetType returns the Type field value
func (o *InterconnectionCreateInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InterconnectionCreateInput) SetType(v string) {
	o.Type = v
}

// GetRedundancy returns the Redundancy field value
func (o *InterconnectionCreateInput) GetRedundancy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Redundancy
}

// GetRedundancyOk returns a tuple with the Redundancy field value
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetRedundancyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Redundancy, true
}

// SetRedundancy sets field value
func (o *InterconnectionCreateInput) SetRedundancy(v string) {
	o.Redundancy = v
}

// GetContactEmail returns the ContactEmail field value if set, zero value otherwise.
func (o *InterconnectionCreateInput) GetContactEmail() string {
	if o == nil || o.ContactEmail == nil {
		var ret string
		return ret
	}
	return *o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetContactEmailOk() (*string, bool) {
	if o == nil || o.ContactEmail == nil {
		return nil, false
	}
	return o.ContactEmail, true
}

// HasContactEmail returns a boolean if a field has been set.
func (o *InterconnectionCreateInput) HasContactEmail() bool {
	if o != nil && o.ContactEmail != nil {
		return true
	}

	return false
}

// SetContactEmail gets a reference to the given string and assigns it to the ContactEmail field.
func (o *InterconnectionCreateInput) SetContactEmail(v string) {
	o.ContactEmail = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *InterconnectionCreateInput) GetProject() string {
	if o == nil || o.Project == nil {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetProjectOk() (*string, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *InterconnectionCreateInput) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *InterconnectionCreateInput) SetProject(v string) {
	o.Project = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *InterconnectionCreateInput) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *InterconnectionCreateInput) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *InterconnectionCreateInput) SetSpeed(v string) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InterconnectionCreateInput) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterconnectionCreateInput) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InterconnectionCreateInput) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InterconnectionCreateInput) SetTags(v []string) {
	o.Tags = &v
}

func (o InterconnectionCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["facility"] = o.Facility
	}
	if o.Metro != nil {
		toSerialize["metro"] = o.Metro
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["redundancy"] = o.Redundancy
	}
	if o.ContactEmail != nil {
		toSerialize["contact_email"] = o.ContactEmail
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Speed != nil {
		toSerialize["speed"] = o.Speed
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInterconnectionCreateInput struct {
	value *InterconnectionCreateInput
	isSet bool
}

func (v NullableInterconnectionCreateInput) Get() *InterconnectionCreateInput {
	return v.value
}

func (v *NullableInterconnectionCreateInput) Set(val *InterconnectionCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableInterconnectionCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableInterconnectionCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterconnectionCreateInput(val *InterconnectionCreateInput) *NullableInterconnectionCreateInput {
	return &NullableInterconnectionCreateInput{value: val, isSet: true}
}

func (v NullableInterconnectionCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterconnectionCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


