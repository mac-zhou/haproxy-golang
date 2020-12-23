// Code generated by go-swagger; DO NOT EDIT.

package global

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
)

// NewGetGlobalParams creates a new GetGlobalParams object
// with the default values initialized.
func NewGetGlobalParams() *GetGlobalParams {
	var ()
	return &GetGlobalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalParamsWithTimeout creates a new GetGlobalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGlobalParamsWithTimeout(timeout time.Duration) *GetGlobalParams {
	var ()
	return &GetGlobalParams{

		timeout: timeout,
	}
}

// NewGetGlobalParamsWithContext creates a new GetGlobalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGlobalParamsWithContext(ctx context.Context) *GetGlobalParams {
	var ()
	return &GetGlobalParams{

		Context: ctx,
	}
}

// NewGetGlobalParamsWithHTTPClient creates a new GetGlobalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGlobalParamsWithHTTPClient(client *http.Client) *GetGlobalParams {
	var ()
	return &GetGlobalParams{
		HTTPClient: client,
	}
}

/*GetGlobalParams contains all the parameters to send to the API endpoint
for the get global operation typically these are written to a http.Request
*/
type GetGlobalParams struct {

	/*TransactionID
	  ID of the transaction where we want to add the operation. Cannot be used when version is specified.

	*/
	TransactionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get global params
func (o *GetGlobalParams) WithTimeout(timeout time.Duration) *GetGlobalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global params
func (o *GetGlobalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global params
func (o *GetGlobalParams) WithContext(ctx context.Context) *GetGlobalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global params
func (o *GetGlobalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global params
func (o *GetGlobalParams) WithHTTPClient(client *http.Client) *GetGlobalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global params
func (o *GetGlobalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTransactionID adds the transactionID to the get global params
func (o *GetGlobalParams) WithTransactionID(transactionID *string) *GetGlobalParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get global params
func (o *GetGlobalParams) SetTransactionID(transactionID *string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
