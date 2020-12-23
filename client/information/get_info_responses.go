// Code generated by go-swagger; DO NOT EDIT.

package information

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

// GetInfoReader is a Reader for the GetInfo structure.
type GetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInfoOK creates a GetInfoOK with default headers values
func NewGetInfoOK() *GetInfoOK {
	return &GetInfoOK{}
}

/*GetInfoOK handles this case with default header values.

Success
*/
type GetInfoOK struct {
	Payload *models.Info
}

func (o *GetInfoOK) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfoOK  %+v", 200, o.Payload)
}

func (o *GetInfoOK) GetPayload() *models.Info {
	return o.Payload
}

func (o *GetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Info)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInfoDefault creates a GetInfoDefault with default headers values
func NewGetInfoDefault(code int) *GetInfoDefault {
	return &GetInfoDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetInfoDefault handles this case with default header values.

General Error
*/
type GetInfoDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get info default response
func (o *GetInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetInfoDefault) Error() string {
	return fmt.Sprintf("[GET /info][%d] getInfo default  %+v", o._statusCode, o.Payload)
}

func (o *GetInfoDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
