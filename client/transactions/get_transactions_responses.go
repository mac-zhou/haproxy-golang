// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"mac_zhou.exe/go/pkg/mod/github.com/mac-zhou/haproxy-golang/models"
)

// GetTransactionsReader is a Reader for the GetTransactions structure.
type GetTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTransactionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTransactionsOK creates a GetTransactionsOK with default headers values
func NewGetTransactionsOK() *GetTransactionsOK {
	return &GetTransactionsOK{}
}

/*GetTransactionsOK handles this case with default header values.

Success
*/
type GetTransactionsOK struct {
	Payload models.Transactions
}

func (o *GetTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/transactions][%d] getTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetTransactionsOK) GetPayload() models.Transactions {
	return o.Payload
}

func (o *GetTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionsDefault creates a GetTransactionsDefault with default headers values
func NewGetTransactionsDefault(code int) *GetTransactionsDefault {
	return &GetTransactionsDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetTransactionsDefault handles this case with default header values.

General Error
*/
type GetTransactionsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get transactions default response
func (o *GetTransactionsDefault) Code() int {
	return o._statusCode
}

func (o *GetTransactionsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/transactions][%d] getTransactions default  %+v", o._statusCode, o.Payload)
}

func (o *GetTransactionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTransactionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
