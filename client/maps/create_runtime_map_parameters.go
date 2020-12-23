// Code generated by go-swagger; DO NOT EDIT.

package maps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCreateRuntimeMapParams creates a new CreateRuntimeMapParams object
// with the default values initialized.
func NewCreateRuntimeMapParams() *CreateRuntimeMapParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateRuntimeMapParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRuntimeMapParamsWithTimeout creates a new CreateRuntimeMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRuntimeMapParamsWithTimeout(timeout time.Duration) *CreateRuntimeMapParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateRuntimeMapParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewCreateRuntimeMapParamsWithContext creates a new CreateRuntimeMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRuntimeMapParamsWithContext(ctx context.Context) *CreateRuntimeMapParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateRuntimeMapParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewCreateRuntimeMapParamsWithHTTPClient creates a new CreateRuntimeMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRuntimeMapParamsWithHTTPClient(client *http.Client) *CreateRuntimeMapParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateRuntimeMapParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*CreateRuntimeMapParams contains all the parameters to send to the API endpoint
for the create runtime map operation typically these are written to a http.Request
*/
type CreateRuntimeMapParams struct {

	/*FileUpload
	  The map file to upload

	*/
	FileUpload runtime.NamedReadCloser
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create runtime map params
func (o *CreateRuntimeMapParams) WithTimeout(timeout time.Duration) *CreateRuntimeMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create runtime map params
func (o *CreateRuntimeMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create runtime map params
func (o *CreateRuntimeMapParams) WithContext(ctx context.Context) *CreateRuntimeMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create runtime map params
func (o *CreateRuntimeMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create runtime map params
func (o *CreateRuntimeMapParams) WithHTTPClient(client *http.Client) *CreateRuntimeMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create runtime map params
func (o *CreateRuntimeMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFileUpload adds the fileUpload to the create runtime map params
func (o *CreateRuntimeMapParams) WithFileUpload(fileUpload runtime.NamedReadCloser) *CreateRuntimeMapParams {
	o.SetFileUpload(fileUpload)
	return o
}

// SetFileUpload adds the fileUpload to the create runtime map params
func (o *CreateRuntimeMapParams) SetFileUpload(fileUpload runtime.NamedReadCloser) {
	o.FileUpload = fileUpload
}

// WithForceReload adds the forceReload to the create runtime map params
func (o *CreateRuntimeMapParams) WithForceReload(forceReload *bool) *CreateRuntimeMapParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create runtime map params
func (o *CreateRuntimeMapParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRuntimeMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FileUpload != nil {

		if o.FileUpload != nil {

			// form file param fileUpload
			if err := r.SetFileParam("fileUpload", o.FileUpload); err != nil {
				return err
			}

		}

	}

	if o.ForceReload != nil {

		// query param force_reload
		var qrForceReload bool
		if o.ForceReload != nil {
			qrForceReload = *o.ForceReload
		}
		qForceReload := swag.FormatBool(qrForceReload)
		if qForceReload != "" {
			if err := r.SetQueryParam("force_reload", qForceReload); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}