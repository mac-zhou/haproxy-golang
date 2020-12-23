// Code generated by go-swagger; DO NOT EDIT.

package nameserver

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

	"mac_zhou.exe/go/pkg/mod/github.com/mac-zhou/haproxy-golang/models"
)

// NewReplaceNameserverParams creates a new ReplaceNameserverParams object
// with the default values initialized.
func NewReplaceNameserverParams() *ReplaceNameserverParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceNameserverParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceNameserverParamsWithTimeout creates a new ReplaceNameserverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceNameserverParamsWithTimeout(timeout time.Duration) *ReplaceNameserverParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceNameserverParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewReplaceNameserverParamsWithContext creates a new ReplaceNameserverParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceNameserverParamsWithContext(ctx context.Context) *ReplaceNameserverParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceNameserverParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewReplaceNameserverParamsWithHTTPClient creates a new ReplaceNameserverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceNameserverParamsWithHTTPClient(client *http.Client) *ReplaceNameserverParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &ReplaceNameserverParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*ReplaceNameserverParams contains all the parameters to send to the API endpoint
for the replace nameserver operation typically these are written to a http.Request
*/
type ReplaceNameserverParams struct {

	/*Data*/
	Data *models.Nameserver
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Name
	  Nameserver name

	*/
	Name string
	/*Resolver
	  Parent resolver name

	*/
	Resolver string
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string
	/*Version
	  Version used for checking configuration version. Cannot be used when transaction is specified, transaction has it's own version.

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace nameserver params
func (o *ReplaceNameserverParams) WithTimeout(timeout time.Duration) *ReplaceNameserverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace nameserver params
func (o *ReplaceNameserverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace nameserver params
func (o *ReplaceNameserverParams) WithContext(ctx context.Context) *ReplaceNameserverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace nameserver params
func (o *ReplaceNameserverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace nameserver params
func (o *ReplaceNameserverParams) WithHTTPClient(client *http.Client) *ReplaceNameserverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace nameserver params
func (o *ReplaceNameserverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the replace nameserver params
func (o *ReplaceNameserverParams) WithData(data *models.Nameserver) *ReplaceNameserverParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the replace nameserver params
func (o *ReplaceNameserverParams) SetData(data *models.Nameserver) {
	o.Data = data
}

// WithForceReload adds the forceReload to the replace nameserver params
func (o *ReplaceNameserverParams) WithForceReload(forceReload *bool) *ReplaceNameserverParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the replace nameserver params
func (o *ReplaceNameserverParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithName adds the name to the replace nameserver params
func (o *ReplaceNameserverParams) WithName(name string) *ReplaceNameserverParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the replace nameserver params
func (o *ReplaceNameserverParams) SetName(name string) {
	o.Name = name
}

// WithResolver adds the resolver to the replace nameserver params
func (o *ReplaceNameserverParams) WithResolver(resolver string) *ReplaceNameserverParams {
	o.SetResolver(resolver)
	return o
}

// SetResolver adds the resolver to the replace nameserver params
func (o *ReplaceNameserverParams) SetResolver(resolver string) {
	o.Resolver = resolver
}

// WithTransactionID adds the transactionID to the replace nameserver params
func (o *ReplaceNameserverParams) WithTransactionID(transactionID *string) *ReplaceNameserverParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the replace nameserver params
func (o *ReplaceNameserverParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the replace nameserver params
func (o *ReplaceNameserverParams) WithVersion(version *int64) *ReplaceNameserverParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the replace nameserver params
func (o *ReplaceNameserverParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceNameserverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// query param resolver
	qrResolver := o.Resolver
	qResolver := qrResolver
	if qResolver != "" {
		if err := r.SetQueryParam("resolver", qResolver); err != nil {
			return err
		}
	}

	if o.TransactionID != nil {

		// query param transaction_id
		var qrTransactionID string
		if o.TransactionID != nil {
			qrTransactionID = *o.TransactionID
		}
		qTransactionID := qrTransactionID
		if qTransactionID != "" {
			if err := r.SetQueryParam("transaction_id", qTransactionID); err != nil {
				return err
			}
		}

	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
