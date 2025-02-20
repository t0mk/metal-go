# \InterconnectionsApi

All URIs are relative to *https://api.equinix.com/metal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterconnectionPortVirtualCircuit**](InterconnectionsApi.md#CreateInterconnectionPortVirtualCircuit) | **Post** /connections/{connection_id}/ports/{port_id}/virtual-circuits | Create a new Virtual Circuit
[**CreateOrganizationInterconnection**](InterconnectionsApi.md#CreateOrganizationInterconnection) | **Post** /organizations/{organization_id}/connections | Request a new interconnection for the organization
[**CreateProjectInterconnection**](InterconnectionsApi.md#CreateProjectInterconnection) | **Post** /projects/{project_id}/connections | Request a new interconnection for the project&#39;s organization
[**DeleteInterconnection**](InterconnectionsApi.md#DeleteInterconnection) | **Delete** /connections/{connection_id} | Delete interconnection
[**DeleteVirtualCircuit**](InterconnectionsApi.md#DeleteVirtualCircuit) | **Delete** /virtual-circuits/{id} | Delete a virtual circuit
[**GetInterconnection**](InterconnectionsApi.md#GetInterconnection) | **Get** /connections/{connection_id} | Get interconnection
[**GetInterconnectionPort**](InterconnectionsApi.md#GetInterconnectionPort) | **Get** /connections/{connection_id}/ports/{id} | Get a interconnection port
[**GetVirtualCircuit**](InterconnectionsApi.md#GetVirtualCircuit) | **Get** /virtual-circuits/{id} | Get a virtual circuit
[**ListInterconnectionPortVirtualCircuits**](InterconnectionsApi.md#ListInterconnectionPortVirtualCircuits) | **Get** /connections/{connection_id}/ports/{port_id}/virtual-circuits | List a interconnection port&#39;s virtual circuits
[**ListInterconnectionPorts**](InterconnectionsApi.md#ListInterconnectionPorts) | **Get** /connections/{connection_id}/ports | List a interconnection&#39;s ports
[**ListInterconnectionVirtualCircuits**](InterconnectionsApi.md#ListInterconnectionVirtualCircuits) | **Get** /connections/{connection_id}/virtual-circuits | List a interconnection&#39;s virtual circuits
[**OrganizationListInterconnections**](InterconnectionsApi.md#OrganizationListInterconnections) | **Get** /organizations/{organization_id}/connections | List organization connections
[**ProjectListInterconnections**](InterconnectionsApi.md#ProjectListInterconnections) | **Get** /projects/{project_id}/connections | List project connections
[**UpdateInterconnection**](InterconnectionsApi.md#UpdateInterconnection) | **Put** /connections/{connection_id} | Update interconnection
[**UpdateVirtualCircuit**](InterconnectionsApi.md#UpdateVirtualCircuit) | **Put** /virtual-circuits/{id} | Update a virtual circuit



## CreateInterconnectionPortVirtualCircuit

> CreateInterconnectionPortVirtualCircuit201Response CreateInterconnectionPortVirtualCircuit(ctx, connectionId, portId).CreateInterconnectionPortVirtualCircuitRequest(createInterconnectionPortVirtualCircuitRequest).Execute()

Create a new Virtual Circuit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection
    portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection port
    createInterconnectionPortVirtualCircuitRequest := openapiclient.createInterconnectionPortVirtualCircuit_request{VirtualCircuitCreateInput: openapiclient.NewVirtualCircuitCreateInput("ProjectId_example")} // CreateInterconnectionPortVirtualCircuitRequest | Virtual Circuit details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.CreateInterconnectionPortVirtualCircuit(context.Background(), connectionId, portId).CreateInterconnectionPortVirtualCircuitRequest(createInterconnectionPortVirtualCircuitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.CreateInterconnectionPortVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInterconnectionPortVirtualCircuit`: CreateInterconnectionPortVirtualCircuit201Response
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.CreateInterconnectionPortVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the interconnection | 
**portId** | **string** | UUID of the interconnection port | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterconnectionPortVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createInterconnectionPortVirtualCircuitRequest** | [**CreateInterconnectionPortVirtualCircuitRequest**](CreateInterconnectionPortVirtualCircuitRequest.md) | Virtual Circuit details | 

### Return type

[**CreateInterconnectionPortVirtualCircuit201Response**](CreateInterconnectionPortVirtualCircuit201Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInterconnection

> Interconnection CreateOrganizationInterconnection(ctx, organizationId).InterconnectionCreateInput(interconnectionCreateInput).Execute()

Request a new interconnection for the organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization
    interconnectionCreateInput := *openapiclient.NewInterconnectionCreateInput("Metro_example", "Name_example", "Redundancy_example", "Type_example") // InterconnectionCreateInput | Interconnection details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.CreateOrganizationInterconnection(context.Background(), organizationId).InterconnectionCreateInput(interconnectionCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.CreateOrganizationInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.CreateOrganizationInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | UUID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interconnectionCreateInput** | [**InterconnectionCreateInput**](InterconnectionCreateInput.md) | Interconnection details | 

### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProjectInterconnection

> Interconnection CreateProjectInterconnection(ctx, projectId).InterconnectionCreateInput(interconnectionCreateInput).Execute()

Request a new interconnection for the project's organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the project
    interconnectionCreateInput := *openapiclient.NewInterconnectionCreateInput("Metro_example", "Name_example", "Redundancy_example", "Type_example") // InterconnectionCreateInput | Interconnection details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.CreateProjectInterconnection(context.Background(), projectId).InterconnectionCreateInput(interconnectionCreateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.CreateProjectInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProjectInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.CreateProjectInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | UUID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interconnectionCreateInput** | [**InterconnectionCreateInput**](InterconnectionCreateInput.md) | Interconnection details | 

### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterconnection

> Interconnection DeleteInterconnection(ctx, connectionId).Execute()

Delete interconnection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Interconnection UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.DeleteInterconnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.DeleteInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.DeleteInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Interconnection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVirtualCircuit

> CreateInterconnectionPortVirtualCircuit201Response DeleteVirtualCircuit(ctx, id).Execute()

Delete a virtual circuit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.DeleteVirtualCircuit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.DeleteVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualCircuit`: CreateInterconnectionPortVirtualCircuit201Response
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.DeleteVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateInterconnectionPortVirtualCircuit201Response**](CreateInterconnectionPortVirtualCircuit201Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnection

> Interconnection GetInterconnection(ctx, connectionId).Execute()

Get interconnection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Interconnection UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.GetInterconnection(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.GetInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.GetInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Interconnection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterconnectionPort

> InterconnectionPort GetInterconnectionPort(ctx, connectionId, id).Execute()

Get a interconnection port



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Port UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.GetInterconnectionPort(context.Background(), connectionId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.GetInterconnectionPort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterconnectionPort`: InterconnectionPort
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.GetInterconnectionPort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the interconnection | 
**id** | **string** | Port UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterconnectionPortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InterconnectionPort**](InterconnectionPort.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualCircuit

> CreateInterconnectionPortVirtualCircuit201Response GetVirtualCircuit(ctx, id).Execute()

Get a virtual circuit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.GetVirtualCircuit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.GetVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualCircuit`: CreateInterconnectionPortVirtualCircuit201Response
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.GetVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateInterconnectionPortVirtualCircuit201Response**](CreateInterconnectionPortVirtualCircuit201Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterconnectionPortVirtualCircuits

> VirtualCircuitList ListInterconnectionPortVirtualCircuits(ctx, connectionId, portId).Execute()

List a interconnection port's virtual circuits



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection
    portId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection port

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.ListInterconnectionPortVirtualCircuits(context.Background(), connectionId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.ListInterconnectionPortVirtualCircuits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInterconnectionPortVirtualCircuits`: VirtualCircuitList
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.ListInterconnectionPortVirtualCircuits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the interconnection | 
**portId** | **string** | UUID of the interconnection port | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterconnectionPortVirtualCircuitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VirtualCircuitList**](VirtualCircuitList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterconnectionPorts

> InterconnectionPortList ListInterconnectionPorts(ctx, connectionId).Execute()

List a interconnection's ports



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.ListInterconnectionPorts(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.ListInterconnectionPorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInterconnectionPorts`: InterconnectionPortList
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.ListInterconnectionPorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the interconnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterconnectionPortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterconnectionPortList**](InterconnectionPortList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterconnectionVirtualCircuits

> VirtualCircuitList ListInterconnectionVirtualCircuits(ctx, connectionId).Execute()

List a interconnection's virtual circuits



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the interconnection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.ListInterconnectionVirtualCircuits(context.Background(), connectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.ListInterconnectionVirtualCircuits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInterconnectionVirtualCircuits`: VirtualCircuitList
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.ListInterconnectionVirtualCircuits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | UUID of the interconnection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterconnectionVirtualCircuitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualCircuitList**](VirtualCircuitList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationListInterconnections

> InterconnectionList OrganizationListInterconnections(ctx, organizationId).Execute()

List organization connections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the organization

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.OrganizationListInterconnections(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.OrganizationListInterconnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrganizationListInterconnections`: InterconnectionList
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.OrganizationListInterconnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | UUID of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationListInterconnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterconnectionList**](InterconnectionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectListInterconnections

> InterconnectionList ProjectListInterconnections(ctx, projectId).Execute()

List project connections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.ProjectListInterconnections(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.ProjectListInterconnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectListInterconnections`: InterconnectionList
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.ProjectListInterconnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | UUID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectListInterconnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterconnectionList**](InterconnectionList.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterconnection

> Interconnection UpdateInterconnection(ctx, connectionId).InterconnectionUpdateInput(interconnectionUpdateInput).Execute()

Update interconnection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Interconnection UUID
    interconnectionUpdateInput := *openapiclient.NewInterconnectionUpdateInput() // InterconnectionUpdateInput | Updated interconnection details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.UpdateInterconnection(context.Background(), connectionId).InterconnectionUpdateInput(interconnectionUpdateInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.UpdateInterconnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterconnection`: Interconnection
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.UpdateInterconnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Interconnection UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterconnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **interconnectionUpdateInput** | [**InterconnectionUpdateInput**](InterconnectionUpdateInput.md) | Updated interconnection details | 

### Return type

[**Interconnection**](Interconnection.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVirtualCircuit

> CreateInterconnectionPortVirtualCircuit201Response UpdateVirtualCircuit(ctx, id).UpdateVirtualCircuitRequest(updateVirtualCircuitRequest).Execute()

Update a virtual circuit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Virtual Circuit UUID
    updateVirtualCircuitRequest := openapiclient.updateVirtualCircuit_request{VirtualCircuitUpdateInput: openapiclient.NewVirtualCircuitUpdateInput()} // UpdateVirtualCircuitRequest | Updated Virtual Circuit details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterconnectionsApi.UpdateVirtualCircuit(context.Background(), id).UpdateVirtualCircuitRequest(updateVirtualCircuitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterconnectionsApi.UpdateVirtualCircuit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualCircuit`: CreateInterconnectionPortVirtualCircuit201Response
    fmt.Fprintf(os.Stdout, "Response from `InterconnectionsApi.UpdateVirtualCircuit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Virtual Circuit UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVirtualCircuitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVirtualCircuitRequest** | [**UpdateVirtualCircuitRequest**](UpdateVirtualCircuitRequest.md) | Updated Virtual Circuit details | 

### Return type

[**CreateInterconnectionPortVirtualCircuit201Response**](CreateInterconnectionPortVirtualCircuit201Response.md)

### Authorization

[x_auth_token](../README.md#x_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

