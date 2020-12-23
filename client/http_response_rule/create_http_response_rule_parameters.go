// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

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

// NewCreateHTTPResponseRuleParams creates a new CreateHTTPResponseRuleParams object
// with the default values initialized.
func NewCreateHTTPResponseRuleParams() *CreateHTTPResponseRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateHTTPResponseRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHTTPResponseRuleParamsWithTimeout creates a new CreateHTTPResponseRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateHTTPResponseRuleParamsWithTimeout(timeout time.Duration) *CreateHTTPResponseRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateHTTPResponseRuleParams{
		ForceReload: &forceReloadDefault,

		timeout: timeout,
	}
}

// NewCreateHTTPResponseRuleParamsWithContext creates a new CreateHTTPResponseRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateHTTPResponseRuleParamsWithContext(ctx context.Context) *CreateHTTPResponseRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateHTTPResponseRuleParams{
		ForceReload: &forceReloadDefault,

		Context: ctx,
	}
}

// NewCreateHTTPResponseRuleParamsWithHTTPClient creates a new CreateHTTPResponseRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateHTTPResponseRuleParamsWithHTTPClient(client *http.Client) *CreateHTTPResponseRuleParams {
	var (
		forceReloadDefault = bool(false)
	)
	return &CreateHTTPResponseRuleParams{
		ForceReload: &forceReloadDefault,
		HTTPClient:  client,
	}
}

/*CreateHTTPResponseRuleParams contains all the parameters to send to the API endpoint
for the create HTTP response rule operation typically these are written to a http.Request
*/
type CreateHTTPResponseRuleParams struct {

	/*Data*/
	Data *models.HTTPResponseRule
	/*ForceReload
	  If set, do a force reload, do not wait for the configured reload-delay. Cannot be used when transaction is specified, as changes in transaction are not applied directly to configuration.

	*/
	ForceReload *bool
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

// WithTimeout adds the timeout to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithTimeout(timeout time.Duration) *CreateHTTPResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithContext(ctx context.Context) *CreateHTTPResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithHTTPClient(client *http.Client) *CreateHTTPResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithData(data *models.HTTPResponseRule) *CreateHTTPResponseRuleParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetData(data *models.HTTPResponseRule) {
	o.Data = data
}

// WithForceReload adds the forceReload to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithForceReload(forceReload *bool) *CreateHTTPResponseRuleParams {
	o.SetForceReload(forceReload)
	return o
}

// SetForceReload adds the forceReload to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetForceReload(forceReload *bool) {
	o.ForceReload = forceReload
}

// WithParentName adds the parentName to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithParentName(parentName string) *CreateHTTPResponseRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithParentType(parentType string) *CreateHTTPResponseRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithTransactionID(transactionID *string) *CreateHTTPResponseRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WithVersion adds the version to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) WithVersion(version *int64) *CreateHTTPResponseRuleParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create HTTP response rule params
func (o *CreateHTTPResponseRuleParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHTTPResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
