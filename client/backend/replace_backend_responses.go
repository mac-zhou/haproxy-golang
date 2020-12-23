// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/mac-zhou/haproxy-golang/models"
)

// ReplaceBackendReader is a Reader for the ReplaceBackend structure.
type ReplaceBackendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceBackendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceBackendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceBackendAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceBackendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceBackendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceBackendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceBackendOK creates a ReplaceBackendOK with default headers values
func NewReplaceBackendOK() *ReplaceBackendOK {
	return &ReplaceBackendOK{}
}

/*ReplaceBackendOK handles this case with default header values.

Backend replaced
*/
type ReplaceBackendOK struct {
	Payload *models.Backend
}

func (o *ReplaceBackendOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backends/{name}][%d] replaceBackendOK  %+v", 200, o.Payload)
}

func (o *ReplaceBackendOK) GetPayload() *models.Backend {
	return o.Payload
}

func (o *ReplaceBackendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Backend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendAccepted creates a ReplaceBackendAccepted with default headers values
func NewReplaceBackendAccepted() *ReplaceBackendAccepted {
	return &ReplaceBackendAccepted{}
}

/*ReplaceBackendAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceBackendAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Backend
}

func (o *ReplaceBackendAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backends/{name}][%d] replaceBackendAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceBackendAccepted) GetPayload() *models.Backend {
	return o.Payload
}

func (o *ReplaceBackendAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Backend)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendBadRequest creates a ReplaceBackendBadRequest with default headers values
func NewReplaceBackendBadRequest() *ReplaceBackendBadRequest {
	return &ReplaceBackendBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceBackendBadRequest handles this case with default header values.

Bad request
*/
type ReplaceBackendBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceBackendBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backends/{name}][%d] replaceBackendBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceBackendBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceBackendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceBackendNotFound creates a ReplaceBackendNotFound with default headers values
func NewReplaceBackendNotFound() *ReplaceBackendNotFound {
	return &ReplaceBackendNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceBackendNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceBackendNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceBackendNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backends/{name}][%d] replaceBackendNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceBackendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceBackendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceBackendDefault creates a ReplaceBackendDefault with default headers values
func NewReplaceBackendDefault(code int) *ReplaceBackendDefault {
	return &ReplaceBackendDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceBackendDefault handles this case with default header values.

General Error
*/
type ReplaceBackendDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace backend default response
func (o *ReplaceBackendDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceBackendDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backends/{name}][%d] replaceBackend default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceBackendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceBackendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
