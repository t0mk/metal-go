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

// checks if the IPReservationRequestInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPReservationRequestInput{}

// IPReservationRequestInput struct for IPReservationRequestInput
type IPReservationRequestInput struct {
	Comments               *string                `json:"comments,omitempty"`
	Customdata             map[string]interface{} `json:"customdata,omitempty"`
	Details                *string                `json:"details,omitempty"`
	Facility               *string                `json:"facility,omitempty"`
	FailOnApprovalRequired *bool                  `json:"fail_on_approval_required,omitempty"`
	// The code of the metro you are requesting the IP reservation in.
	Metro    *string  `json:"metro,omitempty"`
	Quantity int32    `json:"quantity"`
	Tags     []string `json:"tags,omitempty"`
	Type     string   `json:"type"`
}

// NewIPReservationRequestInput instantiates a new IPReservationRequestInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPReservationRequestInput(quantity int32, type_ string) *IPReservationRequestInput {
	this := IPReservationRequestInput{}
	this.Quantity = quantity
	this.Type = type_
	return &this
}

// NewIPReservationRequestInputWithDefaults instantiates a new IPReservationRequestInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPReservationRequestInputWithDefaults() *IPReservationRequestInput {
	this := IPReservationRequestInput{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetComments() string {
	if o == nil || isNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetCommentsOk() (*string, bool) {
	if o == nil || isNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasComments() bool {
	if o != nil && !isNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IPReservationRequestInput) SetComments(v string) {
	o.Comments = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetCustomdata() map[string]interface{} {
	if o == nil || isNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasCustomdata() bool {
	if o != nil && !isNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *IPReservationRequestInput) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetDetails() string {
	if o == nil || isNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetDetailsOk() (*string, bool) {
	if o == nil || isNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasDetails() bool {
	if o != nil && !isNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *IPReservationRequestInput) SetDetails(v string) {
	o.Details = &v
}

// GetFacility returns the Facility field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetFacility() string {
	if o == nil || isNil(o.Facility) {
		var ret string
		return ret
	}
	return *o.Facility
}

// GetFacilityOk returns a tuple with the Facility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetFacilityOk() (*string, bool) {
	if o == nil || isNil(o.Facility) {
		return nil, false
	}
	return o.Facility, true
}

// HasFacility returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasFacility() bool {
	if o != nil && !isNil(o.Facility) {
		return true
	}

	return false
}

// SetFacility gets a reference to the given string and assigns it to the Facility field.
func (o *IPReservationRequestInput) SetFacility(v string) {
	o.Facility = &v
}

// GetFailOnApprovalRequired returns the FailOnApprovalRequired field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetFailOnApprovalRequired() bool {
	if o == nil || isNil(o.FailOnApprovalRequired) {
		var ret bool
		return ret
	}
	return *o.FailOnApprovalRequired
}

// GetFailOnApprovalRequiredOk returns a tuple with the FailOnApprovalRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetFailOnApprovalRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.FailOnApprovalRequired) {
		return nil, false
	}
	return o.FailOnApprovalRequired, true
}

// HasFailOnApprovalRequired returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasFailOnApprovalRequired() bool {
	if o != nil && !isNil(o.FailOnApprovalRequired) {
		return true
	}

	return false
}

// SetFailOnApprovalRequired gets a reference to the given bool and assigns it to the FailOnApprovalRequired field.
func (o *IPReservationRequestInput) SetFailOnApprovalRequired(v bool) {
	o.FailOnApprovalRequired = &v
}

// GetMetro returns the Metro field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetMetro() string {
	if o == nil || isNil(o.Metro) {
		var ret string
		return ret
	}
	return *o.Metro
}

// GetMetroOk returns a tuple with the Metro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetMetroOk() (*string, bool) {
	if o == nil || isNil(o.Metro) {
		return nil, false
	}
	return o.Metro, true
}

// HasMetro returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasMetro() bool {
	if o != nil && !isNil(o.Metro) {
		return true
	}

	return false
}

// SetMetro gets a reference to the given string and assigns it to the Metro field.
func (o *IPReservationRequestInput) SetMetro(v string) {
	o.Metro = &v
}

// GetQuantity returns the Quantity field value
func (o *IPReservationRequestInput) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *IPReservationRequestInput) SetQuantity(v int32) {
	o.Quantity = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IPReservationRequestInput) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IPReservationRequestInput) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IPReservationRequestInput) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *IPReservationRequestInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IPReservationRequestInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IPReservationRequestInput) SetType(v string) {
	o.Type = v
}

func (o IPReservationRequestInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPReservationRequestInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !isNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !isNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !isNil(o.Facility) {
		toSerialize["facility"] = o.Facility
	}
	if !isNil(o.FailOnApprovalRequired) {
		toSerialize["fail_on_approval_required"] = o.FailOnApprovalRequired
	}
	if !isNil(o.Metro) {
		toSerialize["metro"] = o.Metro
	}
	toSerialize["quantity"] = o.Quantity
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableIPReservationRequestInput struct {
	value *IPReservationRequestInput
	isSet bool
}

func (v NullableIPReservationRequestInput) Get() *IPReservationRequestInput {
	return v.value
}

func (v *NullableIPReservationRequestInput) Set(val *IPReservationRequestInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIPReservationRequestInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIPReservationRequestInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPReservationRequestInput(val *IPReservationRequestInput) *NullableIPReservationRequestInput {
	return &NullableIPReservationRequestInput{value: val, isSet: true}
}

func (v NullableIPReservationRequestInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPReservationRequestInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
