/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:    ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// PerformActionRequest struct for PerformActionRequest
type PerformActionRequest struct {
	// Action to perform. See Device.actions for possible actions.
	Type string `json:"type"`
	// May be required to perform actions under certain conditions
	ForceDelete *bool `json:"force_delete,omitempty"`
	// When type is `reinstall`, enabling fast deprovisioning will bypass full disk wiping.
	DeprovisionFast *bool `json:"deprovision_fast,omitempty"`
	// When type is `reinstall`, preserve the existing data on all disks except the operating-system disk.
	PreserveData *bool `json:"preserve_data,omitempty"`
	// When type is `reinstall`, use this `operating_system` (defaults to the current `operating system`)
	OperatingSystem *string `json:"operating_system,omitempty"`
	// When type is `reinstall`, use this `ipxe_script_url` (`operating_system` must be `custom_ipxe`, defaults to the current `ipxe_script_url`)
	IpxeScriptUrl *string `json:"ipxe_script_url,omitempty"`
}

// NewPerformActionRequest instantiates a new PerformActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformActionRequest(type_ string) *PerformActionRequest {
	this := PerformActionRequest{}
	this.Type = type_
	return &this
}

// NewPerformActionRequestWithDefaults instantiates a new PerformActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformActionRequestWithDefaults() *PerformActionRequest {
	this := PerformActionRequest{}
	return &this
}

// GetType returns the Type field value
func (o *PerformActionRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PerformActionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PerformActionRequest) SetType(v string) {
	o.Type = v
}

// GetForceDelete returns the ForceDelete field value if set, zero value otherwise.
func (o *PerformActionRequest) GetForceDelete() bool {
	if o == nil || isNil(o.ForceDelete) {
		var ret bool
		return ret
	}
	return *o.ForceDelete
}

// GetForceDeleteOk returns a tuple with the ForceDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionRequest) GetForceDeleteOk() (*bool, bool) {
	if o == nil || isNil(o.ForceDelete) {
		return nil, false
	}
	return o.ForceDelete, true
}

// HasForceDelete returns a boolean if a field has been set.
func (o *PerformActionRequest) HasForceDelete() bool {
	if o != nil && !isNil(o.ForceDelete) {
		return true
	}

	return false
}

// SetForceDelete gets a reference to the given bool and assigns it to the ForceDelete field.
func (o *PerformActionRequest) SetForceDelete(v bool) {
	o.ForceDelete = &v
}

// GetDeprovisionFast returns the DeprovisionFast field value if set, zero value otherwise.
func (o *PerformActionRequest) GetDeprovisionFast() bool {
	if o == nil || isNil(o.DeprovisionFast) {
		var ret bool
		return ret
	}
	return *o.DeprovisionFast
}

// GetDeprovisionFastOk returns a tuple with the DeprovisionFast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionRequest) GetDeprovisionFastOk() (*bool, bool) {
	if o == nil || isNil(o.DeprovisionFast) {
		return nil, false
	}
	return o.DeprovisionFast, true
}

// HasDeprovisionFast returns a boolean if a field has been set.
func (o *PerformActionRequest) HasDeprovisionFast() bool {
	if o != nil && !isNil(o.DeprovisionFast) {
		return true
	}

	return false
}

// SetDeprovisionFast gets a reference to the given bool and assigns it to the DeprovisionFast field.
func (o *PerformActionRequest) SetDeprovisionFast(v bool) {
	o.DeprovisionFast = &v
}

// GetPreserveData returns the PreserveData field value if set, zero value otherwise.
func (o *PerformActionRequest) GetPreserveData() bool {
	if o == nil || isNil(o.PreserveData) {
		var ret bool
		return ret
	}
	return *o.PreserveData
}

// GetPreserveDataOk returns a tuple with the PreserveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionRequest) GetPreserveDataOk() (*bool, bool) {
	if o == nil || isNil(o.PreserveData) {
		return nil, false
	}
	return o.PreserveData, true
}

// HasPreserveData returns a boolean if a field has been set.
func (o *PerformActionRequest) HasPreserveData() bool {
	if o != nil && !isNil(o.PreserveData) {
		return true
	}

	return false
}

// SetPreserveData gets a reference to the given bool and assigns it to the PreserveData field.
func (o *PerformActionRequest) SetPreserveData(v bool) {
	o.PreserveData = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *PerformActionRequest) GetOperatingSystem() string {
	if o == nil || isNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionRequest) GetOperatingSystemOk() (*string, bool) {
	if o == nil || isNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *PerformActionRequest) HasOperatingSystem() bool {
	if o != nil && !isNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *PerformActionRequest) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *PerformActionRequest) GetIpxeScriptUrl() string {
	if o == nil || isNil(o.IpxeScriptUrl) {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformActionRequest) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || isNil(o.IpxeScriptUrl) {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *PerformActionRequest) HasIpxeScriptUrl() bool {
	if o != nil && !isNil(o.IpxeScriptUrl) {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *PerformActionRequest) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

func (o PerformActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ForceDelete) {
		toSerialize["force_delete"] = o.ForceDelete
	}
	if !isNil(o.DeprovisionFast) {
		toSerialize["deprovision_fast"] = o.DeprovisionFast
	}
	if !isNil(o.PreserveData) {
		toSerialize["preserve_data"] = o.PreserveData
	}
	if !isNil(o.OperatingSystem) {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if !isNil(o.IpxeScriptUrl) {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePerformActionRequest struct {
	value *PerformActionRequest
	isSet bool
}

func (v NullablePerformActionRequest) Get() *PerformActionRequest {
	return v.value
}

func (v *NullablePerformActionRequest) Set(val *PerformActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformActionRequest(val *PerformActionRequest) *NullablePerformActionRequest {
	return &NullablePerformActionRequest{value: val, isSet: true}
}

func (v NullablePerformActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
