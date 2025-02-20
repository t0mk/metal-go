/*
Metal API

# Introduction Equinix Metal provides a RESTful HTTP API which can be reached at <https://api.equinix.com/metal/v1>. This document describes the API and how to use it.  The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account. Every feature of the Equinix Metal web interface is accessible through the API.  The API docs are generated from the Equinix Metal OpenAPI specification and are officially hosted at <https://metal.equinix.com/developers/api>.  # Common Parameters  The Equinix Metal API uses a few methods to minimize network traffic and improve throughput. These parameters are not used in all API calls, but are used often enough to warrant their own section. Look for these parameters in the documentation for the API calls that support them.  ## Pagination  Pagination is used to limit the number of results returned in a single request. The API will return a maximum of 100 results per page. To retrieve additional results, you can use the `page` and `per_page` query parameters.  The `page` parameter is used to specify the page number. The first page is `1`. The `per_page` parameter is used to specify the number of results per page. The maximum number of results differs by resource type.  ## Sorting  Where offered, the API allows you to sort results by a specific field. To sort results use the `sort_by` query parameter with the root level field name as the value. The `sort_direction` parameter is used to specify the sort direction, either either `asc` (ascending) or `desc` (descending).  ## Filtering  Filtering is used to limit the results returned in a single request. The API supports filtering by certain fields in the response. To filter results, you can use the field as a query parameter.  For example, to filter the IP list to only return public IPv4 addresses, you can filter by the `type` field, as in the following request:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/projects/id/ips?type=public_ipv4 ```  Only IP addresses with the `type` field set to `public_ipv4` will be returned.  ## Searching  Searching is used to find matching resources using multiple field comparissons. The API supports searching in resources that define this behavior. The fields available for search differ by resource, as does the search strategy.  To search resources you can use the `search` query parameter.  ## Include and Exclude  For resources that contain references to other resources, sucha as a Device that refers to the Project it resides in, the Equinix Metal API will returns `href` values (API links) to the associated resource.  ```json {   ...   \"project\": {     \"href\": \"/metal/v1/projects/f3f131c8-f302-49ef-8c44-9405022dc6dd\"   } } ```  If you're going need the project details, you can avoid a second API request.  Specify the contained `href` resources and collections that you'd like to have included in the response using the `include` query parameter.  For example:  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=projects ```  The `include` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests where `href` resources are presented.  To have multiple resources include, use a comma-separated list (e.g. `?include=emails,projects,memberships`).  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=emails,projects,memberships ```  You may also include nested associations up to three levels deep using dot notation (`?include=memberships.projects`):  ```sh curl -H 'X-Auth-Token: my_authentication_token' \\   https://api.equinix.com/metal/v1/user?include=memberships.projects ```  To exclude resources, and optimize response delivery, use the `exclude` query parameter. The `exclude` parameter is generally accepted in `GET`, `POST`, `PUT`, and `PATCH` requests for fields with nested object responses. When excluded, these fields will be replaced with an object that contains only an `href` field.

API version: 1.0.0
Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the VirtualNetworkCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualNetworkCreateInput{}

// VirtualNetworkCreateInput struct for VirtualNetworkCreateInput
type VirtualNetworkCreateInput struct {
	Description *string `json:"description,omitempty"`
	// The UUID (or facility code) for the Facility in which to create this Virtual network.
	Facility *string `json:"facility,omitempty"`
	// The UUID (or metro code) for the Metro in which to create this Virtual Network.
	Metro *string `json:"metro,omitempty"`
	// VLAN ID between 2-3999. Must be unique for the project within the Metro in which this Virtual Network is being created. If no value is specified, the next-available VLAN ID in the range 1000-1999 will be automatically selected.
	Vxlan *int32 `json:"vxlan,omitempty"`
}

// NewVirtualNetworkCreateInput instantiates a new VirtualNetworkCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualNetworkCreateInput() *VirtualNetworkCreateInput {
	this := VirtualNetworkCreateInput{}
	return &this
}

// NewVirtualNetworkCreateInputWithDefaults instantiates a new VirtualNetworkCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualNetworkCreateInputWithDefaults() *VirtualNetworkCreateInput {
	this := VirtualNetworkCreateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VirtualNetworkCreateInput) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VirtualNetworkCreateInput) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VirtualNetworkCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *VirtualNetworkCreateInput) GetFacility() string {
	if o == nil || isNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkCreateInput) GetFacilityOk() (*string, bool) {
	if o == nil || isNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *VirtualNetworkCreateInput) HasFacility() bool {
	if o != nil && !isNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *VirtualNetworkCreateInput) SetFacility(v string) {
	o.Facility = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *VirtualNetworkCreateInput) GetMetro() string {
	if o == nil || isNil(o.Metro) {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkCreateInput) GetMetroOk() (*string, bool) {
	if o == nil || isNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *VirtualNetworkCreateInput) HasMetro() bool {
	if o != nil && !isNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *VirtualNetworkCreateInput) SetMetro(v string) {
	o.Metro = &v
}

// GetVxlan returns the Vxlan field value if set, zero value otherwise.
func (o *VirtualNetworkCreateInput) GetVxlan() int32 {
	if o == nil || isNil(o.Vxlan) {
		var ret int32
		return ret
	}
	return *o.Vxlan
}

// GetVxlanOk returns a tuple with the Vxlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetworkCreateInput) GetVxlanOk() (*int32, bool) {
	if o == nil || isNil(o.Vxlan) {
		return nil, false
	}
	return o.Vxlan, true
}

// HasVxlan returns a boolean if a field has been set.
func (o *VirtualNetworkCreateInput) HasVxlan() bool {
	if o != nil && !isNil(o.Vxlan) {
		return true
	}

	return false
}

// SetVxlan gets a reference to the given int32 and assigns it to the Vxlan field.
func (o *VirtualNetworkCreateInput) SetVxlan(v int32) {
	o.Vxlan = &v
}

func (o VirtualNetworkCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualNetworkCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !isNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !isNil(o.Vxlan) {
		toSerialize["vxlan"] = o.Vxlan
	}
	return toSerialize, nil
}

type NullableVirtualNetworkCreateInput struct {
	value *VirtualNetworkCreateInput
	isSet bool
}

func (v NullableVirtualNetworkCreateInput) Get() *VirtualNetworkCreateInput {
	return v.value
}

func (v *NullableVirtualNetworkCreateInput) Set(val *VirtualNetworkCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualNetworkCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualNetworkCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualNetworkCreateInput(val *VirtualNetworkCreateInput) *NullableVirtualNetworkCreateInput {
	return &NullableVirtualNetworkCreateInput{value: val, isSet: true}
}

func (v NullableVirtualNetworkCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualNetworkCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
