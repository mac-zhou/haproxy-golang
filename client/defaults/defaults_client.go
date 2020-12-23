// Code generated by go-swagger; DO NOT EDIT.

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new defaults API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for defaults API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetDefaults(params *GetDefaultsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultsOK, error)

	ReplaceDefaults(params *ReplaceDefaultsParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceDefaultsOK, *ReplaceDefaultsAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDefaults returns defaults part of configuration

  Returns defaults part of configuration.
*/
func (a *Client) GetDefaults(params *GetDefaultsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDefaults",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDefaultsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDefaultsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefaultsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceDefaults replaces defaults

  Replace defaults part of config
*/
func (a *Client) ReplaceDefaults(params *ReplaceDefaultsParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceDefaultsOK, *ReplaceDefaultsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceDefaultsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceDefaults",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/defaults",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceDefaultsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceDefaultsOK:
		return value, nil, nil
	case *ReplaceDefaultsAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceDefaultsDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
