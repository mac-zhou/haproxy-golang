// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	InitiateCertificateRefresh(params *InitiateCertificateRefreshParams, authInfo runtime.ClientAuthInfoWriter) (*InitiateCertificateRefreshOK, error)

	PostCluster(params *PostClusterParams, authInfo runtime.ClientAuthInfoWriter) (*PostClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InitiateCertificateRefresh initiates a certificate refresh

  Initiates a certificate refresh
*/
func (a *Client) InitiateCertificateRefresh(params *InitiateCertificateRefreshParams, authInfo runtime.ClientAuthInfoWriter) (*InitiateCertificateRefreshOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInitiateCertificateRefreshParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "initiateCertificateRefresh",
		Method:             "POST",
		PathPattern:        "/cluster/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InitiateCertificateRefreshReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InitiateCertificateRefreshOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*InitiateCertificateRefreshDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostCluster posts cluster settings

  Post cluster settings
*/
func (a *Client) PostCluster(params *PostClusterParams, authInfo runtime.ClientAuthInfoWriter) (*PostClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCluster",
		Method:             "POST",
		PathPattern:        "/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
