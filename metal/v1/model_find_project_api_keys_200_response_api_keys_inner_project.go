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

// FindProjectAPIKeys200ResponseApiKeysInnerProject struct for FindProjectAPIKeys200ResponseApiKeysInnerProject
type FindProjectAPIKeys200ResponseApiKeysInnerProject struct {
	BgpConfig     *FindBatchById200ResponseDevicesInner  `json:"bgp_config,omitempty"`
	CreatedAt     *time.Time                             `json:"created_at,omitempty"`
	Customdata    map[string]interface{}                 `json:"customdata,omitempty"`
	Devices       []FindBatchById200ResponseDevicesInner `json:"devices,omitempty"`
	Id            *string                                `json:"id,omitempty"`
	Invitations   []FindBatchById200ResponseDevicesInner `json:"invitations,omitempty"`
	MaxDevices    map[string]interface{}                 `json:"max_devices,omitempty"`
	Members       []FindBatchById200ResponseDevicesInner `json:"members,omitempty"`
	Memberships   []FindBatchById200ResponseDevicesInner `json:"memberships,omitempty"`
	Name          *string                                `json:"name,omitempty"`
	NetworkStatus map[string]interface{}                 `json:"network_status,omitempty"`
	PaymentMethod *FindBatchById200ResponseDevicesInner  `json:"payment_method,omitempty"`
	SshKeys       []FindBatchById200ResponseDevicesInner `json:"ssh_keys,omitempty"`
	UpdatedAt     *time.Time                             `json:"updated_at,omitempty"`
	Volumes       []FindBatchById200ResponseDevicesInner `json:"volumes,omitempty"`
}

// NewFindProjectAPIKeys200ResponseApiKeysInnerProject instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindProjectAPIKeys200ResponseApiKeysInnerProject() *FindProjectAPIKeys200ResponseApiKeysInnerProject {
	this := FindProjectAPIKeys200ResponseApiKeysInnerProject{}
	return &this
}

// NewFindProjectAPIKeys200ResponseApiKeysInnerProjectWithDefaults instantiates a new FindProjectAPIKeys200ResponseApiKeysInnerProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindProjectAPIKeys200ResponseApiKeysInnerProjectWithDefaults() *FindProjectAPIKeys200ResponseApiKeysInnerProject {
	this := FindProjectAPIKeys200ResponseApiKeysInnerProject{}
	return &this
}

// GetBgpConfig returns the BgpConfig field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetBgpConfig() FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.BgpConfig) {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.BgpConfig
}

// GetBgpConfigOk returns a tuple with the BgpConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetBgpConfigOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.BgpConfig) {
		return nil, false
	}
	return o.BgpConfig, true
}

// HasBgpConfig returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasBgpConfig() bool {
	if o != nil && !isNil(o.BgpConfig) {
		return true
	}

	return false
}

// SetBgpConfig gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the BgpConfig field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetBgpConfig(v FindBatchById200ResponseDevicesInner) {
	o.BgpConfig = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomdata returns the Customdata field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCustomdata() map[string]interface{} {
	if o == nil || isNil(o.Customdata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Customdata
}

// GetCustomdataOk returns a tuple with the Customdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetCustomdataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Customdata) {
		return map[string]interface{}{}, false
	}
	return o.Customdata, true
}

// HasCustomdata returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasCustomdata() bool {
	if o != nil && !isNil(o.Customdata) {
		return true
	}

	return false
}

// SetCustomdata gets a reference to the given map[string]interface{} and assigns it to the Customdata field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetCustomdata(v map[string]interface{}) {
	o.Customdata = v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetDevices() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.Devices) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetDevicesOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.Devices) {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasDevices() bool {
	if o != nil && !isNil(o.Devices) {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Devices field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetDevices(v []FindBatchById200ResponseDevicesInner) {
	o.Devices = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetId(v string) {
	o.Id = &v
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetInvitations() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.Invitations) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetInvitationsOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.Invitations) {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasInvitations() bool {
	if o != nil && !isNil(o.Invitations) {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Invitations field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetInvitations(v []FindBatchById200ResponseDevicesInner) {
	o.Invitations = v
}

// GetMaxDevices returns the MaxDevices field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMaxDevices() map[string]interface{} {
	if o == nil || isNil(o.MaxDevices) {
		var ret map[string]interface{}
		return ret
	}
	return o.MaxDevices
}

// GetMaxDevicesOk returns a tuple with the MaxDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMaxDevicesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.MaxDevices) {
		return map[string]interface{}{}, false
	}
	return o.MaxDevices, true
}

// HasMaxDevices returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasMaxDevices() bool {
	if o != nil && !isNil(o.MaxDevices) {
		return true
	}

	return false
}

// SetMaxDevices gets a reference to the given map[string]interface{} and assigns it to the MaxDevices field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetMaxDevices(v map[string]interface{}) {
	o.MaxDevices = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMembers() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.Members) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMembersOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasMembers() bool {
	if o != nil && !isNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Members field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetMembers(v []FindBatchById200ResponseDevicesInner) {
	o.Members = v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMemberships() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.Memberships) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetMembershipsOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasMemberships() bool {
	if o != nil && !isNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Memberships field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetMemberships(v []FindBatchById200ResponseDevicesInner) {
	o.Memberships = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetName(v string) {
	o.Name = &v
}

// GetNetworkStatus returns the NetworkStatus field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetNetworkStatus() map[string]interface{} {
	if o == nil || isNil(o.NetworkStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.NetworkStatus
}

// GetNetworkStatusOk returns a tuple with the NetworkStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetNetworkStatusOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.NetworkStatus) {
		return map[string]interface{}{}, false
	}
	return o.NetworkStatus, true
}

// HasNetworkStatus returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasNetworkStatus() bool {
	if o != nil && !isNil(o.NetworkStatus) {
		return true
	}

	return false
}

// SetNetworkStatus gets a reference to the given map[string]interface{} and assigns it to the NetworkStatus field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetNetworkStatus(v map[string]interface{}) {
	o.NetworkStatus = v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetPaymentMethod() FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.PaymentMethod) {
		var ret FindBatchById200ResponseDevicesInner
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetPaymentMethodOk() (*FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasPaymentMethod() bool {
	if o != nil && !isNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given FindBatchById200ResponseDevicesInner and assigns it to the PaymentMethod field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetPaymentMethod(v FindBatchById200ResponseDevicesInner) {
	o.PaymentMethod = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetSshKeys() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.SshKeys) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetSshKeysOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.SshKeys) {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasSshKeys() bool {
	if o != nil && !isNil(o.SshKeys) {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the SshKeys field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetSshKeys(v []FindBatchById200ResponseDevicesInner) {
	o.SshKeys = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetVolumes() []FindBatchById200ResponseDevicesInner {
	if o == nil || isNil(o.Volumes) {
		var ret []FindBatchById200ResponseDevicesInner
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) GetVolumesOk() ([]FindBatchById200ResponseDevicesInner, bool) {
	if o == nil || isNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) HasVolumes() bool {
	if o != nil && !isNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []FindBatchById200ResponseDevicesInner and assigns it to the Volumes field.
func (o *FindProjectAPIKeys200ResponseApiKeysInnerProject) SetVolumes(v []FindBatchById200ResponseDevicesInner) {
	o.Volumes = v
}

func (o FindProjectAPIKeys200ResponseApiKeysInnerProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BgpConfig) {
		toSerialize["bgp_config"] = o.BgpConfig
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.Customdata) {
		toSerialize["customdata"] = o.Customdata
	}
	if !isNil(o.Devices) {
		toSerialize["devices"] = o.Devices
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Invitations) {
		toSerialize["invitations"] = o.Invitations
	}
	if !isNil(o.MaxDevices) {
		toSerialize["max_devices"] = o.MaxDevices
	}
	if !isNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !isNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.NetworkStatus) {
		toSerialize["network_status"] = o.NetworkStatus
	}
	if !isNil(o.PaymentMethod) {
		toSerialize["payment_method"] = o.PaymentMethod
	}
	if !isNil(o.SshKeys) {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return json.Marshal(toSerialize)
}

type NullableFindProjectAPIKeys200ResponseApiKeysInnerProject struct {
	value *FindProjectAPIKeys200ResponseApiKeysInnerProject
	isSet bool
}

func (v NullableFindProjectAPIKeys200ResponseApiKeysInnerProject) Get() *FindProjectAPIKeys200ResponseApiKeysInnerProject {
	return v.value
}

func (v *NullableFindProjectAPIKeys200ResponseApiKeysInnerProject) Set(val *FindProjectAPIKeys200ResponseApiKeysInnerProject) {
	v.value = val
	v.isSet = true
}

func (v NullableFindProjectAPIKeys200ResponseApiKeysInnerProject) IsSet() bool {
	return v.isSet
}

func (v *NullableFindProjectAPIKeys200ResponseApiKeysInnerProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindProjectAPIKeys200ResponseApiKeysInnerProject(val *FindProjectAPIKeys200ResponseApiKeysInnerProject) *NullableFindProjectAPIKeys200ResponseApiKeysInnerProject {
	return &NullableFindProjectAPIKeys200ResponseApiKeysInnerProject{value: val, isSet: true}
}

func (v NullableFindProjectAPIKeys200ResponseApiKeysInnerProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindProjectAPIKeys200ResponseApiKeysInnerProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
