# VirtualCircuitListVirtualCircuitsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bill** | **bool** | True if the Virtual Circuit is being billed. Currently, only Virtual Circuits of Fabric VCs (Metal Billed) will be billed. Usage will start the first time the Virtual Circuit becomes active, and will not stop until it is deleted from Metal. | [default to false]
**Description** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**NniVlan** | **int32** |  | 
**Port** | [**Href**](Href.md) |  | 
**Project** | [**Href**](Href.md) |  | 
**Speed** | Pointer to **int32** | integer representing bps speed | [optional] 
**Status** | **string** |  | 
**Tags** | **[]string** |  | 
**VirtualNetwork** | [**Href**](Href.md) |  | 
**Vnid** | **int32** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomerIp** | Pointer to **string** | An IP address from the subnet that will be used on the Customer side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Metal IP. By default, the last usable IP address in the subnet will be used. | [optional] 
**Md5** | Pointer to **string** | The MD5 password for the BGP peering in plaintext (not a checksum). | [optional] 
**MetalIp** | Pointer to **string** | An IP address from the subnet that will be used on the Metal side. This parameter is optional, but if supplied, we will use the other usable IP address in the subnet as the Customer IP. By default, the first usable IP address in the subnet will be used. | [optional] 
**PeerAsn** | Pointer to **int32** | The peer ASN that will be used with the VRF on the Virtual Circuit. | [optional] 
**Subnet** | Pointer to **string** | The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. This subnet does not have to be an existing VRF IP reservation, as we will create the VRF IP reservation on creation if it does not exist. The Metal IP and Customer IP must be IPs from this subnet. For /30 subnets, the network and broadcast IPs cannot be used as the Metal or Customer IP. | [optional] 
**Vrf** | Pointer to [**Vrf**](Vrf.md) |  | [optional] 

## Methods

### NewVirtualCircuitListVirtualCircuitsInner

`func NewVirtualCircuitListVirtualCircuitsInner(bill bool, description string, id string, name string, nniVlan int32, port Href, project Href, status string, tags []string, virtualNetwork Href, vnid int32, ) *VirtualCircuitListVirtualCircuitsInner`

NewVirtualCircuitListVirtualCircuitsInner instantiates a new VirtualCircuitListVirtualCircuitsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCircuitListVirtualCircuitsInnerWithDefaults

`func NewVirtualCircuitListVirtualCircuitsInnerWithDefaults() *VirtualCircuitListVirtualCircuitsInner`

NewVirtualCircuitListVirtualCircuitsInnerWithDefaults instantiates a new VirtualCircuitListVirtualCircuitsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBill

`func (o *VirtualCircuitListVirtualCircuitsInner) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *VirtualCircuitListVirtualCircuitsInner) SetBill(v bool)`

SetBill sets Bill field to given value.


### GetDescription

`func (o *VirtualCircuitListVirtualCircuitsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCircuitListVirtualCircuitsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *VirtualCircuitListVirtualCircuitsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCircuitListVirtualCircuitsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VirtualCircuitListVirtualCircuitsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCircuitListVirtualCircuitsInner) SetName(v string)`

SetName sets Name field to given value.


### GetNniVlan

`func (o *VirtualCircuitListVirtualCircuitsInner) GetNniVlan() int32`

GetNniVlan returns the NniVlan field if non-nil, zero value otherwise.

### GetNniVlanOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetNniVlanOk() (*int32, bool)`

GetNniVlanOk returns a tuple with the NniVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNniVlan

`func (o *VirtualCircuitListVirtualCircuitsInner) SetNniVlan(v int32)`

SetNniVlan sets NniVlan field to given value.


### GetPort

`func (o *VirtualCircuitListVirtualCircuitsInner) GetPort() Href`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetPortOk() (*Href, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualCircuitListVirtualCircuitsInner) SetPort(v Href)`

SetPort sets Port field to given value.


### GetProject

`func (o *VirtualCircuitListVirtualCircuitsInner) GetProject() Href`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetProjectOk() (*Href, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VirtualCircuitListVirtualCircuitsInner) SetProject(v Href)`

SetProject sets Project field to given value.


### GetSpeed

`func (o *VirtualCircuitListVirtualCircuitsInner) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VirtualCircuitListVirtualCircuitsInner) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VirtualCircuitListVirtualCircuitsInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualCircuitListVirtualCircuitsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCircuitListVirtualCircuitsInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTags

`func (o *VirtualCircuitListVirtualCircuitsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualCircuitListVirtualCircuitsInner) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetVirtualNetwork

`func (o *VirtualCircuitListVirtualCircuitsInner) GetVirtualNetwork() Href`

GetVirtualNetwork returns the VirtualNetwork field if non-nil, zero value otherwise.

### GetVirtualNetworkOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetVirtualNetworkOk() (*Href, bool)`

GetVirtualNetworkOk returns a tuple with the VirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetwork

`func (o *VirtualCircuitListVirtualCircuitsInner) SetVirtualNetwork(v Href)`

SetVirtualNetwork sets VirtualNetwork field to given value.


### GetVnid

`func (o *VirtualCircuitListVirtualCircuitsInner) GetVnid() int32`

GetVnid returns the Vnid field if non-nil, zero value otherwise.

### GetVnidOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetVnidOk() (*int32, bool)`

GetVnidOk returns a tuple with the Vnid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnid

`func (o *VirtualCircuitListVirtualCircuitsInner) SetVnid(v int32)`

SetVnid sets Vnid field to given value.


### GetCreatedAt

`func (o *VirtualCircuitListVirtualCircuitsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualCircuitListVirtualCircuitsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualCircuitListVirtualCircuitsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VirtualCircuitListVirtualCircuitsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualCircuitListVirtualCircuitsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VirtualCircuitListVirtualCircuitsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCustomerIp

`func (o *VirtualCircuitListVirtualCircuitsInner) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *VirtualCircuitListVirtualCircuitsInner) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *VirtualCircuitListVirtualCircuitsInner) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetMd5

`func (o *VirtualCircuitListVirtualCircuitsInner) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *VirtualCircuitListVirtualCircuitsInner) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *VirtualCircuitListVirtualCircuitsInner) SetMd5(v string)`

SetMd5 sets Md5 field to given value.

### HasMd5

`func (o *VirtualCircuitListVirtualCircuitsInner) HasMd5() bool`

HasMd5 returns a boolean if a field has been set.

### GetMetalIp

`func (o *VirtualCircuitListVirtualCircuitsInner) GetMetalIp() string`

GetMetalIp returns the MetalIp field if non-nil, zero value otherwise.

### GetMetalIpOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetMetalIpOk() (*string, bool)`

GetMetalIpOk returns a tuple with the MetalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalIp

`func (o *VirtualCircuitListVirtualCircuitsInner) SetMetalIp(v string)`

SetMetalIp sets MetalIp field to given value.

### HasMetalIp

`func (o *VirtualCircuitListVirtualCircuitsInner) HasMetalIp() bool`

HasMetalIp returns a boolean if a field has been set.

### GetPeerAsn

`func (o *VirtualCircuitListVirtualCircuitsInner) GetPeerAsn() int32`

GetPeerAsn returns the PeerAsn field if non-nil, zero value otherwise.

### GetPeerAsnOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetPeerAsnOk() (*int32, bool)`

GetPeerAsnOk returns a tuple with the PeerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerAsn

`func (o *VirtualCircuitListVirtualCircuitsInner) SetPeerAsn(v int32)`

SetPeerAsn sets PeerAsn field to given value.

### HasPeerAsn

`func (o *VirtualCircuitListVirtualCircuitsInner) HasPeerAsn() bool`

HasPeerAsn returns a boolean if a field has been set.

### GetSubnet

`func (o *VirtualCircuitListVirtualCircuitsInner) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *VirtualCircuitListVirtualCircuitsInner) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *VirtualCircuitListVirtualCircuitsInner) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVrf

`func (o *VirtualCircuitListVirtualCircuitsInner) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VirtualCircuitListVirtualCircuitsInner) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VirtualCircuitListVirtualCircuitsInner) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.

### HasVrf

`func (o *VirtualCircuitListVirtualCircuitsInner) HasVrf() bool`

HasVrf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


