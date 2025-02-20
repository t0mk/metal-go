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

// checks if the TransferRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferRequest{}

// TransferRequest struct for TransferRequest
type TransferRequest struct {
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	Href               *string    `json:"href,omitempty"`
	Id                 *string    `json:"id,omitempty"`
	Project            *Href      `json:"project,omitempty"`
	TargetOrganization *Href      `json:"target_organization,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}

// NewTransferRequest instantiates a new TransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferRequest() *TransferRequest {
	this := TransferRequest{}
	return &this
}

// NewTransferRequestWithDefaults instantiates a new TransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferRequestWithDefaults() *TransferRequest {
	this := TransferRequest{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TransferRequest) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TransferRequest) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TransferRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *TransferRequest) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *TransferRequest) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *TransferRequest) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransferRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransferRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TransferRequest) SetId(v string) {
	o.Id = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *TransferRequest) GetProject() Href {
	if o == nil || isNil(o.Project) {
		var ret Href
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetProjectOk() (*Href, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *TransferRequest) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Href and assigns it to the Project field.
func (o *TransferRequest) SetProject(v Href) {
	o.Project = &v
}

// GetTargetOrganization returns the TargetOrganization field value if set, zero value otherwise.
func (o *TransferRequest) GetTargetOrganization() Href {
	if o == nil || isNil(o.TargetOrganization) {
		var ret Href
		return ret
	}
	return *o.TargetOrganization
}

// GetTargetOrganizationOk returns a tuple with the TargetOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetTargetOrganizationOk() (*Href, bool) {
	if o == nil || isNil(o.TargetOrganization) {
		return nil, false
	}
	return o.TargetOrganization, true
}

// HasTargetOrganization returns a boolean if a field has been set.
func (o *TransferRequest) HasTargetOrganization() bool {
	if o != nil && !isNil(o.TargetOrganization) {
		return true
	}

	return false
}

// SetTargetOrganization gets a reference to the given Href and assigns it to the TargetOrganization field.
func (o *TransferRequest) SetTargetOrganization(v Href) {
	o.TargetOrganization = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TransferRequest) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TransferRequest) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TransferRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.TargetOrganization) {
		toSerialize["target_organization"] = o.TargetOrganization
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTransferRequest struct {
	value *TransferRequest
	isSet bool
}

func (v NullableTransferRequest) Get() *TransferRequest {
	return v.value
}

func (v *NullableTransferRequest) Set(val *TransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferRequest(val *TransferRequest) *NullableTransferRequest {
	return &NullableTransferRequest{value: val, isSet: true}
}

func (v NullableTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
