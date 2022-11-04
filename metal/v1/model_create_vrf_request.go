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

// CreateVrfRequest struct for CreateVrfRequest
type CreateVrfRequest struct {
	Description *string `json:"description,omitempty"`
	// A list of CIDR network addresses. Like [\"10.0.0.0/16\", \"2001:d78::/56\"]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /56 and /64. A VRF\\'s IP ranges must be defined in order to create VRF IP Reservations, which can then be used for Metal Gateways or Virtual Circuits.
	IpRanges []string `json:"ip_ranges,omitempty"`
	LocalAsn *int32   `json:"local_asn,omitempty"`
	// The UUID (or metro code) for the Metro in which to create this VRF.
	Metro string `json:"metro"`
	Name  string `json:"name"`
}

// NewCreateVrfRequest instantiates a new CreateVrfRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVrfRequest(metro string, name string) *CreateVrfRequest {
	this := CreateVrfRequest{}
	this.Metro = metro
	this.Name = name
	return &this
}

// NewCreateVrfRequestWithDefaults instantiates a new CreateVrfRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVrfRequestWithDefaults() *CreateVrfRequest {
	this := CreateVrfRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateVrfRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVrfRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateVrfRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateVrfRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *CreateVrfRequest) GetIpRanges() []string {
	if o == nil || isNil(o.IpRanges) {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVrfRequest) GetIpRangesOk() ([]string, bool) {
	if o == nil || isNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *CreateVrfRequest) HasIpRanges() bool {
	if o != nil && !isNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *CreateVrfRequest) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetLocalAsn returns the LocalAsn field value if set, zero value otherwise.
func (o *CreateVrfRequest) GetLocalAsn() int32 {
	if o == nil || isNil(o.LocalAsn) {
		var ret int32
		return ret
	}
	return *o.LocalAsn
}

// GetLocalAsnOk returns a tuple with the LocalAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVrfRequest) GetLocalAsnOk() (*int32, bool) {
	if o == nil || isNil(o.LocalAsn) {
		return nil, false
	}
	return o.LocalAsn, true
}

// HasLocalAsn returns a boolean if a field has been set.
func (o *CreateVrfRequest) HasLocalAsn() bool {
	if o != nil && !isNil(o.LocalAsn) {
		return true
	}

	return false
}

// SetLocalAsn gets a reference to the given int32 and assigns it to the LocalAsn field.
func (o *CreateVrfRequest) SetLocalAsn(v int32) {
	o.LocalAsn = &v
}

// GetMetro returns the Metro field value
func (o *CreateVrfRequest) GetMetro() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metro
}

// GetMetroOk returns a tuple with the Metro field value
// and a boolean to check if the value has been set.
func (o *CreateVrfRequest) GetMetroOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metro, true
}

// SetMetro sets field value
func (o *CreateVrfRequest) SetMetro(v string) {
	o.Metro = v
}

// GetName returns the Name field value
func (o *CreateVrfRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateVrfRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateVrfRequest) SetName(v string) {
	o.Name = v
}

func (o CreateVrfRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IpRanges) {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if !isNil(o.LocalAsn) {
		toSerialize["local_asn"] = o.LocalAsn
	}
	if true {
		toSerialize["metro"] = o.Metro
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateVrfRequest struct {
	value *CreateVrfRequest
	isSet bool
}

func (v NullableCreateVrfRequest) Get() *CreateVrfRequest {
	return v.value
}

func (v *NullableCreateVrfRequest) Set(val *CreateVrfRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVrfRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVrfRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVrfRequest(val *CreateVrfRequest) *NullableCreateVrfRequest {
	return &NullableCreateVrfRequest{value: val, isSet: true}
}

func (v NullableCreateVrfRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVrfRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
