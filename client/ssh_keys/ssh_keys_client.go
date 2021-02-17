// Code generated by go-swagger; DO NOT EDIT.

package ssh_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ssh keys API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ssh keys API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSSHKey(params *CreateSSHKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSSHKeyCreated, error)

	DeleteSSHKey(params *DeleteSSHKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSSHKeyNoContent, error)

	FindSSHKeyByID(params *FindSSHKeyByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSSHKeyByIDOK, error)

	FindSSHKeys(params *FindSSHKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSSHKeysOK, error)

	UpdateSSHKey(params *UpdateSSHKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSSHKeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSSHKey creates a ssh key for the current user

  Creates a ssh key.
*/
func (a *Client) CreateSSHKey(params *CreateSSHKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSSHKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSSHKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSSHKey",
		Method:             "POST",
		PathPattern:        "/ssh-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSSHKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSSHKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createSSHKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteSSHKey deletes the ssh key

  Deletes the ssh key.
*/
func (a *Client) DeleteSSHKey(params *DeleteSSHKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSSHKeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSSHKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSSHKey",
		Method:             "DELETE",
		PathPattern:        "/ssh-keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSSHKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSSHKeyNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteSSHKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindSSHKeyByID retrieves a ssh key

  Returns a single ssh key if the user has access
*/
func (a *Client) FindSSHKeyByID(params *FindSSHKeyByIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSSHKeyByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSSHKeyByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSSHKeyById",
		Method:             "GET",
		PathPattern:        "/ssh-keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSSHKeyByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSSHKeyByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSSHKeyById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  FindSSHKeys retrieves all ssh keys

  Returns a collection of the user’s ssh keys.
*/
func (a *Client) FindSSHKeys(params *FindSSHKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSSHKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSSHKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSSHKeys",
		Method:             "GET",
		PathPattern:        "/ssh-keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSSHKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSSHKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for findSSHKeys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSSHKey updates the ssh key

  Updates the ssh key.
*/
func (a *Client) UpdateSSHKey(params *UpdateSSHKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSSHKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSSHKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSSHKey",
		Method:             "PUT",
		PathPattern:        "/ssh-keys/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSSHKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSSHKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSSHKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
