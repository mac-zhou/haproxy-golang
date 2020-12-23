// Code generated by go-swagger; DO NOT EDIT.

package sites

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sites API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sites API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSite(params *CreateSiteParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSiteCreated, *CreateSiteAccepted, error)

	DeleteSite(params *DeleteSiteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSiteAccepted, *DeleteSiteNoContent, error)

	GetSite(params *GetSiteParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteOK, error)

	GetSites(params *GetSitesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSitesOK, error)

	ReplaceSite(params *ReplaceSiteParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceSiteOK, *ReplaceSiteAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSite adds a site

  Adds a new site to the configuration file.
*/
func (a *Client) CreateSite(params *CreateSiteParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSiteCreated, *CreateSiteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSite",
		Method:             "POST",
		PathPattern:        "/services/haproxy/sites",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateSiteCreated:
		return value, nil, nil
	case *CreateSiteAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSiteDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSite deletes a site

  Deletes a site from the configuration by it's name.
*/
func (a *Client) DeleteSite(params *DeleteSiteParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSiteAccepted, *DeleteSiteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSite",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/sites/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteSiteAccepted:
		return value, nil, nil
	case *DeleteSiteNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSiteDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSite returns a site

  Returns one site configuration by it's name.
*/
func (a *Client) GetSite(params *GetSiteParams, authInfo runtime.ClientAuthInfoWriter) (*GetSiteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSite",
		Method:             "GET",
		PathPattern:        "/services/haproxy/sites/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSiteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSiteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSites returns an array of sites

  Returns an array of all configured sites.
*/
func (a *Client) GetSites(params *GetSitesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSitesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSitesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSites",
		Method:             "GET",
		PathPattern:        "/services/haproxy/sites",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSitesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSitesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSitesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceSite replaces a site

  Replaces a site configuration by it's name.
*/
func (a *Client) ReplaceSite(params *ReplaceSiteParams, authInfo runtime.ClientAuthInfoWriter) (*ReplaceSiteOK, *ReplaceSiteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceSiteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "replaceSite",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/sites/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceSiteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceSiteOK:
		return value, nil, nil
	case *ReplaceSiteAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceSiteDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
