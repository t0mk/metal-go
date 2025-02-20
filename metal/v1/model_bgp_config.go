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

// checks if the BgpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpConfig{}

// BgpConfig struct for BgpConfig
type BgpConfig struct {
	// Autonomous System Number. ASN is required with Global BGP. With Local BGP the private ASN, 65000, is assigned.
	Asn       *int32     `json:"asn,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// In a Local BGP deployment, a customer uses an internal ASN to control routes within a single Equinix Metal datacenter. This means that the routes are never advertised to the global Internet. Global BGP, on the other hand, requires a customer to have a registered ASN and IP space.
	DeploymentType *string `json:"deployment_type,omitempty"`
	Href           *string `json:"href,omitempty"`
	Id             *string `json:"id,omitempty"`
	// The maximum number of route filters allowed per server
	MaxPrefix *int32 `json:"max_prefix,omitempty"`
	// (Optional) Password for BGP session in plaintext (not a checksum)
	Md5     NullableString `json:"md5,omitempty"`
	Project *Href          `json:"project,omitempty"`
	// The IP block ranges associated to the ASN (Populated in Global BGP only)
	Ranges      []GlobalBgpRange `json:"ranges,omitempty"`
	RequestedAt *time.Time       `json:"requested_at,omitempty"`
	// Specifies AS-MACRO (aka AS-SET) to use when building client route filters
	RouteObject *string `json:"route_object,omitempty"`
	// The direct connections between neighboring routers that want to exchange routing information.
	Sessions []BgpSession `json:"sessions,omitempty"`
	// Status of the BGP Config. Status \"requested\" is valid only with the \"global\" deployment_type.
	Status *string `json:"status,omitempty"`
}

// NewBgpConfig instantiates a new BgpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfig() *BgpConfig {
	this := BgpConfig{}
	var maxPrefix int32 = 10
	this.MaxPrefix = &maxPrefix
	return &this
}

// NewBgpConfigWithDefaults instantiates a new BgpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigWithDefaults() *BgpConfig {
	this := BgpConfig{}
	var maxPrefix int32 = 10
	this.MaxPrefix = &maxPrefix
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *BgpConfig) GetAsn() int32 {
	if o == nil || isNil(o.Asn) {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetAsnOk() (*int32, bool) {
	if o == nil || isNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *BgpConfig) HasAsn() bool {
	if o != nil && !isNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *BgpConfig) SetAsn(v int32) {
	o.Asn = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BgpConfig) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BgpConfig) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BgpConfig) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *BgpConfig) GetDeploymentType() string {
	if o == nil || isNil(o.DeploymentType) {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || isNil(o.DeploymentType) {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *BgpConfig) HasDeploymentType() bool {
	if o != nil && !isNil(o.DeploymentType) {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *BgpConfig) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BgpConfig) GetHref() string {
	if o == nil || isNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetHrefOk() (*string, bool) {
	if o == nil || isNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BgpConfig) HasHref() bool {
	if o != nil && !isNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BgpConfig) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BgpConfig) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BgpConfig) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BgpConfig) SetId(v string) {
	o.Id = &v
}

// GetMaxPrefix returns the MaxPrefix field value if set, zero value otherwise.
func (o *BgpConfig) GetMaxPrefix() int32 {
	if o == nil || isNil(o.MaxPrefix) {
		var ret int32
		return ret
	}
	return *o.MaxPrefix
}

// GetMaxPrefixOk returns a tuple with the MaxPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetMaxPrefixOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPrefix) {
		return nil, false
	}
	return o.MaxPrefix, true
}

// HasMaxPrefix returns a boolean if a field has been set.
func (o *BgpConfig) HasMaxPrefix() bool {
	if o != nil && !isNil(o.MaxPrefix) {
		return true
	}

	return false
}

// SetMaxPrefix gets a reference to the given int32 and assigns it to the MaxPrefix field.
func (o *BgpConfig) SetMaxPrefix(v int32) {
	o.MaxPrefix = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BgpConfig) GetMd5() string {
	if o == nil || isNil(o.Md5.Get()) {
		var ret string
		return ret
	}
	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BgpConfig) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// HasMd5 returns a boolean if a field has been set.
func (o *BgpConfig) HasMd5() bool {
	if o != nil && o.Md5.IsSet() {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given NullableString and assigns it to the Md5 field.
func (o *BgpConfig) SetMd5(v string) {
	o.Md5.Set(&v)
}

// SetMd5Nil sets the value for Md5 to be an explicit nil
func (o *BgpConfig) SetMd5Nil() {
	o.Md5.Set(nil)
}

// UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
func (o *BgpConfig) UnsetMd5() {
	o.Md5.Unset()
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *BgpConfig) GetProject() Href {
	if o == nil || isNil(o.Project) {
		var ret Href
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetProjectOk() (*Href, bool) {
	if o == nil || isNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *BgpConfig) HasProject() bool {
	if o != nil && !isNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Href and assigns it to the Project field.
func (o *BgpConfig) SetProject(v Href) {
	o.Project = &v
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *BgpConfig) GetRanges() []GlobalBgpRange {
	if o == nil || isNil(o.Ranges) {
		var ret []GlobalBgpRange
		return ret
	}
	return o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetRangesOk() ([]GlobalBgpRange, bool) {
	if o == nil || isNil(o.Ranges) {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *BgpConfig) HasRanges() bool {
	if o != nil && !isNil(o.Ranges) {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []GlobalBgpRange and assigns it to the Ranges field.
func (o *BgpConfig) SetRanges(v []GlobalBgpRange) {
	o.Ranges = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BgpConfig) GetRequestedAt() time.Time {
	if o == nil || isNil(o.RequestedAt) {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestedAt) {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BgpConfig) HasRequestedAt() bool {
	if o != nil && !isNil(o.RequestedAt) {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BgpConfig) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetRouteObject returns the RouteObject field value if set, zero value otherwise.
func (o *BgpConfig) GetRouteObject() string {
	if o == nil || isNil(o.RouteObject) {
		var ret string
		return ret
	}
	return *o.RouteObject
}

// GetRouteObjectOk returns a tuple with the RouteObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetRouteObjectOk() (*string, bool) {
	if o == nil || isNil(o.RouteObject) {
		return nil, false
	}
	return o.RouteObject, true
}

// HasRouteObject returns a boolean if a field has been set.
func (o *BgpConfig) HasRouteObject() bool {
	if o != nil && !isNil(o.RouteObject) {
		return true
	}

	return false
}

// SetRouteObject gets a reference to the given string and assigns it to the RouteObject field.
func (o *BgpConfig) SetRouteObject(v string) {
	o.RouteObject = &v
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *BgpConfig) GetSessions() []BgpSession {
	if o == nil || isNil(o.Sessions) {
		var ret []BgpSession
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetSessionsOk() ([]BgpSession, bool) {
	if o == nil || isNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *BgpConfig) HasSessions() bool {
	if o != nil && !isNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []BgpSession and assigns it to the Sessions field.
func (o *BgpConfig) SetSessions(v []BgpSession) {
	o.Sessions = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BgpConfig) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BgpConfig) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BgpConfig) SetStatus(v string) {
	o.Status = &v
}

func (o BgpConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.DeploymentType) {
		toSerialize["deployment_type"] = o.DeploymentType
	}
	if !isNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.MaxPrefix) {
		toSerialize["max_prefix"] = o.MaxPrefix
	}
	if o.Md5.IsSet() {
		toSerialize["md5"] = o.Md5.Get()
	}
	if !isNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !isNil(o.Ranges) {
		toSerialize["ranges"] = o.Ranges
	}
	if !isNil(o.RequestedAt) {
		toSerialize["requested_at"] = o.RequestedAt
	}
	if !isNil(o.RouteObject) {
		toSerialize["route_object"] = o.RouteObject
	}
	if !isNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBgpConfig struct {
	value *BgpConfig
	isSet bool
}

func (v NullableBgpConfig) Get() *BgpConfig {
	return v.value
}

func (v *NullableBgpConfig) Set(val *BgpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfig(val *BgpConfig) *NullableBgpConfig {
	return &NullableBgpConfig{value: val, isSet: true}
}

func (v NullableBgpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
