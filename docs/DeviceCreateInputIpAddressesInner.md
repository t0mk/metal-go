# DeviceCreateInputIpAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **float32** | Address Family for IP Address | [optional] 
**Cidr** | Pointer to **float32** | Cidr Size for the IP Block created. Valid values depends on the operating system being provisioned. (28..32 for IPv4 addresses, 124..127 for IPv6 addresses) | [optional] 
**IpReservations** | Pointer to **[]string** | UUIDs of any IP reservations to use when assigning IPs | [optional] 
**Public** | Pointer to **bool** | Address Type for IP Address | [optional] [default to true]

## Methods

### NewDeviceCreateInputIpAddressesInner

`func NewDeviceCreateInputIpAddressesInner() *DeviceCreateInputIpAddressesInner`

NewDeviceCreateInputIpAddressesInner instantiates a new DeviceCreateInputIpAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateInputIpAddressesInnerWithDefaults

`func NewDeviceCreateInputIpAddressesInnerWithDefaults() *DeviceCreateInputIpAddressesInner`

NewDeviceCreateInputIpAddressesInnerWithDefaults instantiates a new DeviceCreateInputIpAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *DeviceCreateInputIpAddressesInner) GetAddressFamily() float32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *DeviceCreateInputIpAddressesInner) GetAddressFamilyOk() (*float32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *DeviceCreateInputIpAddressesInner) SetAddressFamily(v float32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *DeviceCreateInputIpAddressesInner) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetCidr

`func (o *DeviceCreateInputIpAddressesInner) GetCidr() float32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *DeviceCreateInputIpAddressesInner) GetCidrOk() (*float32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *DeviceCreateInputIpAddressesInner) SetCidr(v float32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *DeviceCreateInputIpAddressesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetIpReservations

`func (o *DeviceCreateInputIpAddressesInner) GetIpReservations() []string`

GetIpReservations returns the IpReservations field if non-nil, zero value otherwise.

### GetIpReservationsOk

`func (o *DeviceCreateInputIpAddressesInner) GetIpReservationsOk() (*[]string, bool)`

GetIpReservationsOk returns a tuple with the IpReservations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpReservations

`func (o *DeviceCreateInputIpAddressesInner) SetIpReservations(v []string)`

SetIpReservations sets IpReservations field to given value.

### HasIpReservations

`func (o *DeviceCreateInputIpAddressesInner) HasIpReservations() bool`

HasIpReservations returns a boolean if a field has been set.

### GetPublic

`func (o *DeviceCreateInputIpAddressesInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *DeviceCreateInputIpAddressesInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *DeviceCreateInputIpAddressesInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *DeviceCreateInputIpAddressesInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


