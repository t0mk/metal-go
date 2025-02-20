# InstancesBatchCreateInputBatchesInnerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Quantity** | Pointer to **int32** | The number of devices to create in this batch. The hostname may contain an &#x60;{{index}}&#x60; placeholder, which will be replaced with the index of the device in the batch. For example, if the hostname is &#x60;device-{{index}}&#x60;, the first device in the batch will have the hostname &#x60;device-01&#x60;, the second device will have the hostname &#x60;device-02&#x60;, and so on. | [optional] 

## Methods

### NewInstancesBatchCreateInputBatchesInnerAllOf

`func NewInstancesBatchCreateInputBatchesInnerAllOf() *InstancesBatchCreateInputBatchesInnerAllOf`

NewInstancesBatchCreateInputBatchesInnerAllOf instantiates a new InstancesBatchCreateInputBatchesInnerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesBatchCreateInputBatchesInnerAllOfWithDefaults

`func NewInstancesBatchCreateInputBatchesInnerAllOfWithDefaults() *InstancesBatchCreateInputBatchesInnerAllOf`

NewInstancesBatchCreateInputBatchesInnerAllOfWithDefaults instantiates a new InstancesBatchCreateInputBatchesInnerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostnames

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetQuantity

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InstancesBatchCreateInputBatchesInnerAllOf) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


