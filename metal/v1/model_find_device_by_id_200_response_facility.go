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

// FindDeviceById200ResponseFacility struct for FindDeviceById200ResponseFacility
type FindDeviceById200ResponseFacility struct {
	Address  *FindDeviceById200ResponseFacilityAddress `json:"address,omitempty"`
	Code     *string                                   `json:"code,omitempty"`
	Features []string                                  `json:"features,omitempty"`
	Id       *string                                   `json:"id,omitempty"`
	// IP ranges registered in facility. Can be used for GeoIP location
	IpRanges []string                                `json:"ip_ranges,omitempty"`
	Metro    *FindDeviceById200ResponseFacilityMetro `json:"metro,omitempty"`
	Name     *string                                 `json:"name,omitempty"`
}

// NewFindDeviceById200ResponseFacility instantiates a new FindDeviceById200ResponseFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindDeviceById200ResponseFacility() *FindDeviceById200ResponseFacility {
	this := FindDeviceById200ResponseFacility{}
	return &this
}

// NewFindDeviceById200ResponseFacilityWithDefaults instantiates a new FindDeviceById200ResponseFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindDeviceById200ResponseFacilityWithDefaults() *FindDeviceById200ResponseFacility {
	this := FindDeviceById200ResponseFacility{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetAddress() FindDeviceById200ResponseFacilityAddress {
	if o == nil || isNil(o.Address) {
		var ret FindDeviceById200ResponseFacilityAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetAddressOk() (*FindDeviceById200ResponseFacilityAddress, bool) {
	if o == nil || isNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given FindDeviceById200ResponseFacilityAddress and assigns it to the Address field.
func (o *FindDeviceById200ResponseFacility) SetAddress(v FindDeviceById200ResponseFacilityAddress) {
	o.Address = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *FindDeviceById200ResponseFacility) SetCode(v string) {
	o.Code = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetFeatures() []string {
	if o == nil || isNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetFeaturesOk() ([]string, bool) {
	if o == nil || isNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasFeatures() bool {
	if o != nil && !isNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *FindDeviceById200ResponseFacility) SetFeatures(v []string) {
	o.Features = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindDeviceById200ResponseFacility) SetId(v string) {
	o.Id = &v
}

// GetIpRanges returns the IpRanges field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetIpRanges() []string {
	if o == nil || isNil(o.IpRanges) {
		var ret []string
		return ret
	}
	return o.IpRanges
}

// GetIpRangesOk returns a tuple with the IpRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetIpRangesOk() ([]string, bool) {
	if o == nil || isNil(o.IpRanges) {
		return nil, false
	}
	return o.IpRanges, true
}

// HasIpRanges returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasIpRanges() bool {
	if o != nil && !isNil(o.IpRanges) {
		return true
	}

	return false
}

// SetIpRanges gets a reference to the given []string and assigns it to the IpRanges field.
func (o *FindDeviceById200ResponseFacility) SetIpRanges(v []string) {
	o.IpRanges = v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetMetro() FindDeviceById200ResponseFacilityMetro {
	if o == nil || isNil(o.Metro) {
		var ret FindDeviceById200ResponseFacilityMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetMetroOk() (*FindDeviceById200ResponseFacilityMetro, bool) {
	if o == nil || isNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasMetro() bool {
	if o != nil && !isNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given FindDeviceById200ResponseFacilityMetro and assigns it to the Metro field.
func (o *FindDeviceById200ResponseFacility) SetMetro(v FindDeviceById200ResponseFacilityMetro) {
	o.Metro = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FindDeviceById200ResponseFacility) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindDeviceById200ResponseFacility) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FindDeviceById200ResponseFacility) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FindDeviceById200ResponseFacility) SetName(v string) {
	o.Name = &v
}

func (o FindDeviceById200ResponseFacility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.IpRanges) {
		toSerialize["ip_ranges"] = o.IpRanges
	}
	if !isNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFindDeviceById200ResponseFacility struct {
	value *FindDeviceById200ResponseFacility
	isSet bool
}

func (v NullableFindDeviceById200ResponseFacility) Get() *FindDeviceById200ResponseFacility {
	return v.value
}

func (v *NullableFindDeviceById200ResponseFacility) Set(val *FindDeviceById200ResponseFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableFindDeviceById200ResponseFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableFindDeviceById200ResponseFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindDeviceById200ResponseFacility(val *FindDeviceById200ResponseFacility) *NullableFindDeviceById200ResponseFacility {
	return &NullableFindDeviceById200ResponseFacility{value: val, isSet: true}
}

func (v NullableFindDeviceById200ResponseFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindDeviceById200ResponseFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
