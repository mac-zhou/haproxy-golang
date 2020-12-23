// Code generated by go-swagger; DO NOT EDIT.

package log_target

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

// NewDeleteLogTargetParams creates a new DeleteLogTargetParams object
// with the default values initialized.
func NewDeleteLogTargetParams() *DeleteLogTargetParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteLogTargetParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLogTargetParamsWithTimeout creates a new DeleteLogTargetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLogTargetParamsWithTimeout(timeout time.Duration) *DeleteLogTargetParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteLogTargetParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewDeleteLogTargetParamsWithContext creates a new DeleteLogTargetParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLogTargetParamsWithContext(ctx context.Context) *DeleteLogTargetParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteLogTargetParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewDeleteLogTargetParamsWithHTTPClient creates a new DeleteLogTargetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLogTargetParamsWithHTTPClient(client *http.Client) *DeleteLogTargetParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &DeleteLogTargetParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*DeleteLogTargetParams contains all the parameters to send to the API endpoint
for the delete log target operation typically these are written to a http.Request
*/
type DeleteLogTargetParams struct {

	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
	/*Index
	  Log Target Index

	*/
	Index int64
	/*ParentName
	  Parent name

	*/
	ParentName string
	/*ParentType
	  Parent type

	*/
	ParentType string
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

// WithTimeout adds the timeout to the delete log target params
func (o *DeleteLogTargetParams) WithTimeout(timeout time.Duration) *DeleteLogTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete log target params
func (o *DeleteLogTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete log target params
func (o *DeleteLogTargetParams) WithContext(ctx context.Context) *DeleteLogTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete log target params
func (o *DeleteLogTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete log target params
func (o *DeleteLogTargetParams) WithHTTPClient(client *http.Client) *DeleteLogTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete log target params
func (o *DeleteLogTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceReload adds the forceReload to the delete log target params
func (o *DeleteLogTargetParams) WithForceReload(forceReload *bool) *DeleteLogTargetParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the delete log target params
func (o *DeleteLogTargetParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithIndex adds the index to the delete log target params
func (o *DeleteLogTargetParams) WithIndex(index int64) *DeleteLogTargetParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the delete log target params
func (o *DeleteLogTargetParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the delete log target params
func (o *DeleteLogTargetParams) WithParentName(parentName string) *DeleteLogTargetParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the delete log target params
func (o *DeleteLogTargetParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the delete log target params
func (o *DeleteLogTargetParams) WithParentType(parentType string) *DeleteLogTargetParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the delete log target params
func (o *DeleteLogTargetParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the delete log target params
func (o *DeleteLogTargetParams) WithTransactionID(transactionID *string) *DeleteLogTargetParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete log target params
func (o *DeleteLogTargetParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the delete log target params
func (o *DeleteLogTargetParams) WithVersion(version *int64) *DeleteLogTargetParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete log target params
func (o *DeleteLogTargetParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLogTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
	}

	// query param parent_name
	qrParentName := o.ParentName
	qParentName := qrParentName
	if qParentName != "" {
		if err := r.SetQueryParam("parent_name", qParentName); err != nil {
			return err
		}
	}

	// query param parent_type
	qrParentType := o.ParentType
	qParentType := qrParentType
	if qParentType != "" {
		if err := r.SetQueryParam("parent_type", qParentType); err != nil {
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
