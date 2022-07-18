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

// RequestBgpConfigRequest struct for RequestBgpConfigRequest
type RequestBgpConfigRequest struct {
	Asn            int32   `json:"asn"`
	DeploymentType string  `json:"deployment_type"`
	Md5            *string `json:"md5,omitempty"`
	UseCase        *string `json:"use_case,omitempty"`
}

// NewRequestBgpConfigRequest instantiates a new RequestBgpConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestBgpConfigRequest(asn int32, deploymentType string) *RequestBgpConfigRequest {
	this := RequestBgpConfigRequest{}
	this.Asn = asn
	this.DeploymentType = deploymentType
	return &this
}

// NewRequestBgpConfigRequestWithDefaults instantiates a new RequestBgpConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestBgpConfigRequestWithDefaults() *RequestBgpConfigRequest {
	this := RequestBgpConfigRequest{}
	return &this
}

// GetAsn returns the Asn field value
func (o *RequestBgpConfigRequest) GetAsn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *RequestBgpConfigRequest) GetAsnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *RequestBgpConfigRequest) SetAsn(v int32) {
	o.Asn = v
}

// GetDeploymentType returns the DeploymentType field value
func (o *RequestBgpConfigRequest) GetDeploymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value
// and a boolean to check if the value has been set.
func (o *RequestBgpConfigRequest) GetDeploymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentType, true
}

// SetDeploymentType sets field value
func (o *RequestBgpConfigRequest) SetDeploymentType(v string) {
	o.DeploymentType = v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *RequestBgpConfigRequest) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBgpConfigRequest) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *RequestBgpConfigRequest) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *RequestBgpConfigRequest) SetMd5(v string) {
	o.Md5 = &v
}

// GetUseCase returns the UseCase field value if set, zero value otherwise.
func (o *RequestBgpConfigRequest) GetUseCase() string {
	if o == nil || o.UseCase == nil {
		var ret string
		return ret
	}
	return *o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestBgpConfigRequest) GetUseCaseOk() (*string, bool) {
	if o == nil || o.UseCase == nil {
		return nil, false
	}
	return o.UseCase, true
}

// HasUseCase returns a boolean if a field has been set.
func (o *RequestBgpConfigRequest) HasUseCase() bool {
	if o != nil && o.UseCase != nil {
		return true
	}

	return false
}

// SetUseCase gets a reference to the given string and assigns it to the UseCase field.
func (o *RequestBgpConfigRequest) SetUseCase(v string) {
	o.UseCase = &v
}

func (o RequestBgpConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["asn"] = o.Asn
	}
	if true {
		toSerialize["deployment_type"] = o.DeploymentType
	}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.UseCase != nil {
		toSerialize["use_case"] = o.UseCase
	}
	return json.Marshal(toSerialize)
}

type NullableRequestBgpConfigRequest struct {
	value *RequestBgpConfigRequest
	isSet bool
}

func (v NullableRequestBgpConfigRequest) Get() *RequestBgpConfigRequest {
	return v.value
}

func (v *NullableRequestBgpConfigRequest) Set(val *RequestBgpConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestBgpConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestBgpConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestBgpConfigRequest(val *RequestBgpConfigRequest) *NullableRequestBgpConfigRequest {
	return &NullableRequestBgpConfigRequest{value: val, isSet: true}
}

func (v NullableRequestBgpConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestBgpConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
