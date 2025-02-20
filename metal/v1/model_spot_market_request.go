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
	"time"
)

// checks if the SpotMarketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpotMarketRequest{}

// SpotMarketRequest struct for SpotMarketRequest
type SpotMarketRequest struct {
	CreatedAt   *time.Time              `json:"created_at,omitempty"`
	DevicesMax  *int32                  `json:"devices_max,omitempty"`
	DevicesMin  *int32                  `json:"devices_min,omitempty"`
	EndAt       *time.Time              `json:"end_at,omitempty"`
	Facilities  *Href                   `json:"facilities,omitempty"`
	Href        *string                 `json:"href,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	Instances   *Href                   `json:"instances,omitempty"`
	MaxBidPrice *float32                `json:"max_bid_price,omitempty"`
	Metro       *SpotMarketRequestMetro `json:"metro,omitempty"`
	Project     *Href                   `json:"project,omitempty"`
}

// NewSpotMarketRequest instantiates a new SpotMarketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpotMarketRequest() *SpotMarketRequest {
	this := SpotMarketRequest{}
	return &this
}

// NewSpotMarketRequestWithDefaults instantiates a new SpotMarketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpotMarketRequestWithDefaults() *SpotMarketRequest {
	this := SpotMarketRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SpotMarketRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDevicesMax returns the DevicesMax field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetDevicesMax() int32 {
	if o == nil || isNil(o.DevicesMax) {
		var ret int32
		return ret
	}
	return *o.DevicesMax
}

// GetDevicesMaxOk returns a tuple with the DevicesMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetDevicesMaxOk() (*int32, bool) {
	if o == nil || isNil(o.DevicesMax) {
		return nil, false
	}
	return o.DevicesMax, true
}

// HasDevicesMax returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasDevicesMax() bool {
	if o != nil && !isNil(o.DevicesMax) {
		return true
	}

	return false
}

// SetDevicesMax gets a reference to the given int32 and assigns it to the DevicesMax field.
func (o *SpotMarketRequest) SetDevicesMax(v int32) {
	o.DevicesMax = &v
}

// GetDevicesMin returns the DevicesMin field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetDevicesMin() int32 {
	if o == nil || isNil(o.DevicesMin) {
		var ret int32
		return ret
	}
	return *o.DevicesMin
}

// GetDevicesMinOk returns a tuple with the DevicesMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetDevicesMinOk() (*int32, bool) {
	if o == nil || isNil(o.DevicesMin) {
		return nil, false
	}
	return o.DevicesMin, true
}

// HasDevicesMin returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasDevicesMin() bool {
	if o != nil && !isNil(o.DevicesMin) {
		return true
	}

	return false
}

// SetDevicesMin gets a reference to the given int32 and assigns it to the DevicesMin field.
func (o *SpotMarketRequest) SetDevicesMin(v int32) {
	o.DevicesMin = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetEndAt() time.Time {
	if o == nil || isNil(o.EndAt) {
		var ret time.Time
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetEndAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndAt) {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasEndAt() bool {
	if o != nil && !isNil(o.EndAt) {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given time.Time and assigns it to the EndAt field.
func (o *SpotMarketRequest) SetEndAt(v time.Time) {
	o.EndAt = &v
}

// GetFacilities returns the Facilities field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetFacilities() Href {
	if o == nil || isNil(o.Facilities) {
		var ret Href
		return ret
	}
	return *o.Facilities
}

// GetFacilitiesOk returns a tuple with the Facilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetFacilitiesOk() (*Href, bool) {
	if o == nil || isNil(o.Facilities) {
		return nil, false
	}
	return o.Facilities, true
}

// HasFacilities returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasFacilities() bool {
	if o != nil && !isNil(o.Facilities) {
		return true
	}

	return false
}

// SetFacilities gets a reference to the given Href and assigns it to the Facilities field.
func (o *SpotMarketRequest) SetFacilities(v Href) {
	o.Facilities = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SpotMarketRequest) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SpotMarketRequest) SetId(v string) {
	o.Id = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetInstances() Href {
	if o == nil || isNil(o.Instances) {
		var ret Href
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetInstancesOk() (*Href, bool) {
	if o == nil || isNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasInstances() bool {
	if o != nil && !isNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given Href and assigns it to the Instances field.
func (o *SpotMarketRequest) SetInstances(v Href) {
	o.Instances = &v
}

// GetMaxBidPrice returns the MaxBidPrice field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetMaxBidPrice() float32 {
	if o == nil || isNil(o.MaxBidPrice) {
		var ret float32
		return ret
	}
	return *o.MaxBidPrice
}

// GetMaxBidPriceOk returns a tuple with the MaxBidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetMaxBidPriceOk() (*float32, bool) {
	if o == nil || isNil(o.MaxBidPrice) {
		return nil, false
	}
	return o.MaxBidPrice, true
}

// HasMaxBidPrice returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasMaxBidPrice() bool {
	if o != nil && !isNil(o.MaxBidPrice) {
		return true
	}

	return false
}

// SetMaxBidPrice gets a reference to the given float32 and assigns it to the MaxBidPrice field.
func (o *SpotMarketRequest) SetMaxBidPrice(v float32) {
	o.MaxBidPrice = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetMetro() SpotMarketRequestMetro {
	if o == nil || isNil(o.Metro) {
		var ret SpotMarketRequestMetro
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetMetroOk() (*SpotMarketRequestMetro, bool) {
	if o == nil || isNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasMetro() bool {
	if o != nil && !isNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given SpotMarketRequestMetro and assigns it to the Metro field.
func (o *SpotMarketRequest) SetMetro(v SpotMarketRequestMetro) {
	o.Metro = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *SpotMarketRequest) GetProject() Href {
	if o == nil || isNil(o.Project) {
		var ret Href
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpotMarketRequest) GetProjectOk() (*Href, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *SpotMarketRequest) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Href and assigns it to the Project field.
func (o *SpotMarketRequest) SetProject(v Href) {
	o.Project = &v
}

func (o SpotMarketRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpotMarketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.DevicesMax) {
		toSerialize["devices_max"] = o.DevicesMax
	}
	if !isNil(o.DevicesMin) {
		toSerialize["devices_min"] = o.DevicesMin
	}
	if !isNil(o.EndAt) {
		toSerialize["end_at"] = o.EndAt
	}
	if !isNil(o.Facilities) {
		toSerialize["facilities"] = o.Facilities
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}
	if !isNil(o.MaxBidPrice) {
		toSerialize["max_bid_price"] = o.MaxBidPrice
	}
	if !isNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	return toSerialize, nil
}

type NullableSpotMarketRequest struct {
	value *SpotMarketRequest
	isSet bool
}

func (v NullableSpotMarketRequest) Get() *SpotMarketRequest {
	return v.value
}

func (v *NullableSpotMarketRequest) Set(val *SpotMarketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSpotMarketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSpotMarketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpotMarketRequest(val *SpotMarketRequest) *NullableSpotMarketRequest {
	return &NullableSpotMarketRequest{value: val, isSet: true}
}

func (v NullableSpotMarketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpotMarketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
