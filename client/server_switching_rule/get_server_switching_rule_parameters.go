// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

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

// NewGetServerSwitchingRuleParams creates a new GetServerSwitchingRuleParams object
// with the default values initialized.
func NewGetServerSwitchingRuleParams() *GetServerSwitchingRuleParams {
	var ()
	return &GetServerSwitchingRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServerSwitchingRuleParamsWithTimeout creates a new GetServerSwitchingRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServerSwitchingRuleParamsWithTimeout(timeout time.Duration) *GetServerSwitchingRuleParams {
	var ()
	return &GetServerSwitchingRuleParams{

		timeout: timeout,
	}
}

// NewGetServerSwitchingRuleParamsWithContext creates a new GetServerSwitchingRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServerSwitchingRuleParamsWithContext(ctx context.Context) *GetServerSwitchingRuleParams {
	var ()
	return &GetServerSwitchingRuleParams{

		Context: ctx,
	}
}

// NewGetServerSwitchingRuleParamsWithHTTPClient creates a new GetServerSwitchingRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServerSwitchingRuleParamsWithHTTPClient(client *http.Client) *GetServerSwitchingRuleParams {
	var ()
	return &GetServerSwitchingRuleParams{
		HTTPClient: client,
	}
}

/*GetServerSwitchingRuleParams contains all the parameters to send to the API endpoint
for the get server switching rule operation typically these are written to a http.Request
*/
type GetServerSwitchingRuleParams struct {

	/*Backend
	  Backend name

	*/
	Backend string
	/*Index
	  Switching Rule Index

	*/
	Index int64
	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get server switching rule params
func (o *GetServerSwitchingRuleParams) WithTimeout(timeout time.Duration) *GetServerSwitchingRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get server switching rule params
func (o *GetServerSwitchingRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get server switching rule params
func (o *GetServerSwitchingRuleParams) WithContext(ctx context.Context) *GetServerSwitchingRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get server switching rule params
func (o *GetServerSwitchingRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get server switching rule params
func (o *GetServerSwitchingRuleParams) WithHTTPClient(client *http.Client) *GetServerSwitchingRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get server switching rule params
func (o *GetServerSwitchingRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackend adds the backend to the get server switching rule params
func (o *GetServerSwitchingRuleParams) WithBackend(backend string) *GetServerSwitchingRuleParams {
	o.SetBackend(backend)
	return o
}

// SetBackend adds the backend to the get server switching rule params
func (o *GetServerSwitchingRuleParams) SetBackend(backend string) {
	o.Backend = backend
}

// WithIndex adds the index to the get server switching rule params
func (o *GetServerSwitchingRuleParams) WithIndex(index int64) *GetServerSwitchingRuleParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the get server switching rule params
func (o *GetServerSwitchingRuleParams) SetIndex(index int64) {
	o.Index = index
}

// WithTransactionID adds the transactionID to the get server switching rule params
func (o *GetServerSwitchingRuleParams) WithTransactionID(transactionID *string) *GetServerSwitchingRuleParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get server switching rule params
func (o *GetServerSwitchingRuleParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServerSwitchingRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param backend
	qrBackend := o.Backend
	qBackend := qrBackend
	if qBackend != "" {
		if err := r.SetQueryParam("backend", qBackend); err != nil {
			return err
		}
	}

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.Index)); err != nil {
		return err
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
