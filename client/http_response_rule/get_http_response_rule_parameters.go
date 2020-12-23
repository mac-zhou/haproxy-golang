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
)

// NewGetHTTPResponseRuleParams creates a new GetHTTPResponseRuleParams object
// with the default values initialized.
func NewGetHTTPResponseRuleParams() *GetHTTPResponseRuleParams {
	var ()
	return &GetHTTPResponseRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHTTPResponseRuleParamsWithTimeout creates a new GetHTTPResponseRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHTTPResponseRuleParamsWithTimeout(timeout time.Duration) *GetHTTPResponseRuleParams {
	var ()
	return &GetHTTPResponseRuleParams{

		timeout: timeout,
	}
}

// NewGetHTTPResponseRuleParamsWithContext creates a new GetHTTPResponseRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHTTPResponseRuleParamsWithContext(ctx context.Context) *GetHTTPResponseRuleParams {
	var ()
	return &GetHTTPResponseRuleParams{

		Context: ctx,
	}
}

// NewGetHTTPResponseRuleParamsWithHTTPClient creates a new GetHTTPResponseRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHTTPResponseRuleParamsWithHTTPClient(client *http.Client) *GetHTTPResponseRuleParams {
	var ()
	return &GetHTTPResponseRuleParams{
		HTTPClient: client,
	}
}

/*GetHTTPResponseRuleParams contains all the parameters to send to the API endpoint
for the get HTTP response rule operation typically these are written to a http.Request
*/
type GetHTTPResponseRuleParams struct {

	/*Index
	  HTTP Response Rule Index

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithTimeout(timeout time.Duration) *GetHTTPResponseRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithContext(ctx context.Context) *GetHTTPResponseRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithHTTPClient(client *http.Client) *GetHTTPResponseRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithIndex(index int64) *GetHTTPResponseRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithParentName adds the parentName to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithParentName(parentName string) *GetHTTPResponseRuleParams {
	o.SetParentName(parentName)
	return o
}

// SetParentName adds the parentName to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetParentName(parentName string) {
	o.ParentName = parentName
}

// WithParentType adds the parentType to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithParentType(parentType string) *GetHTTPResponseRuleParams {
	o.SetParentType(parentType)
	return o
}

// SetParentType adds the parentType to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetParentType(parentType string) {
	o.ParentType = parentType
}

// WithTransactionID adds the transactionID to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) WithTransactionID(transactionID *string) *GetHTTPResponseRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get HTTP response rule params
func (o *GetHTTPResponseRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetHTTPResponseRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}