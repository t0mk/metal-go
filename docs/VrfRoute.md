# VrfRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the newly-created resource | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the route. Potential values are \&quot;pending\&quot;, \&quot;active\&quot;, \&quot;deleting\&quot;, and \&quot;error\&quot;, representing various lifecycle states of the route and whether or not it has been successfully configured on the network | [optional] [readonly] 
**Prefix** | Pointer to **string** | The IPv4 prefix for the route, in CIDR-style notation | [optional] 
**NextHop** | Pointer to **string** | The next-hop IPv4 address for the route | [optional] 
**Type** | Pointer to **string** | VRF route type, like &#39;bgp&#39;, &#39;connected&#39;, and &#39;static&#39;. Currently, only static routes are supported | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**MetalGateway** | Pointer to [**VrfRouteMetalGateway**](VrfRouteMetalGateway.md) |  | [optional] 
**VirtualNetwork** | Pointer to [**VrfRouteVirtualNetwork**](VrfRouteVirtualNetwork.md) |  | [optional] 
**Vrf** | Pointer to [**VrfRouteVrf**](VrfRouteVrf.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewVrfRoute

`func NewVrfRoute() *VrfRoute`

NewVrfRoute instantiates a new VrfRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfRouteWithDefaults

`func NewVrfRouteWithDefaults() *VrfRoute`

NewVrfRouteWithDefaults instantiates a new VrfRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VrfRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfRoute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *VrfRoute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VrfRoute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VrfRoute) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VrfRoute) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrefix

`func (o *VrfRoute) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VrfRoute) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VrfRoute) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VrfRoute) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetNextHop

`func (o *VrfRoute) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *VrfRoute) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *VrfRoute) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *VrfRoute) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetType

`func (o *VrfRoute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VrfRoute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VrfRoute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VrfRoute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VrfRoute) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VrfRoute) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VrfRoute) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VrfRoute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VrfRoute) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VrfRoute) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VrfRoute) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VrfRoute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMetalGateway

`func (o *VrfRoute) GetMetalGateway() VrfRouteMetalGateway`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *VrfRoute) GetMetalGatewayOk() (*VrfRouteMetalGateway, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *VrfRoute) SetMetalGateway(v VrfRouteMetalGateway)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *VrfRoute) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetVirtualNetwork

`func (o *VrfRoute) GetVirtualNetwork() VrfRouteVirtualNetwork`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VrfRoute) GetVirtualNetworkOk() (*VrfRouteVirtualNetwork, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VrfRoute) SetVirtualNetwork(v VrfRouteVirtualNetwork)`

SetVirtualNetwork sets VirtualNetwork field to given value.

### HasVirtualNetwork

`func (o *VrfRoute) HasVirtualNetwork() bool`

HasVirtualNetwork returns a boolean if a field has been set.

### GetVrf

`func (o *VrfRoute) GetVrf() VrfRouteVrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VrfRoute) GetVrfOk() (*VrfRouteVrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VrfRoute) SetVrf(v VrfRouteVrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *VrfRoute) HasVrf() bool`

HasVrf returns a boolean if a field has been set.

### GetHref

`func (o *VrfRoute) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VrfRoute) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VrfRoute) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VrfRoute) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


