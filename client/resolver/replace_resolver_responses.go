// Code generated by go-swagger; DO NOT EDIT.

package resolver

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

// ReplaceResolverReader is a Reader for the ReplaceResolver structure.
type ReplaceResolverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceResolverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceResolverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceResolverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceResolverBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceResolverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceResolverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceResolverOK creates a ReplaceResolverOK with default headers values
func NewReplaceResolverOK() *ReplaceResolverOK {
	return &ReplaceResolverOK{}
}

/*ReplaceResolverOK handles this case with default header values.

Resolver replaced
*/
type ReplaceResolverOK struct {
	Payload *models.Resolver
}

func (o *ReplaceResolverOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/resolvers/{name}][%d] replaceResolverOK  %+v", 200, o.Payload)
}

func (o *ReplaceResolverOK) GetPayload() *models.Resolver {
	return o.Payload
}

func (o *ReplaceResolverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Resolver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceResolverAccepted creates a ReplaceResolverAccepted with default headers values
func NewReplaceResolverAccepted() *ReplaceResolverAccepted {
	return &ReplaceResolverAccepted{}
}

/*ReplaceResolverAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceResolverAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Resolver
}

func (o *ReplaceResolverAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/resolvers/{name}][%d] replaceResolverAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceResolverAccepted) GetPayload() *models.Resolver {
	return o.Payload
}

func (o *ReplaceResolverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Resolver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceResolverBadRequest creates a ReplaceResolverBadRequest with default headers values
func NewReplaceResolverBadRequest() *ReplaceResolverBadRequest {
	return &ReplaceResolverBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceResolverBadRequest handles this case with default header values.

Bad request
*/
type ReplaceResolverBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceResolverBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/resolvers/{name}][%d] replaceResolverBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceResolverBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceResolverBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceResolverNotFound creates a ReplaceResolverNotFound with default headers values
func NewReplaceResolverNotFound() *ReplaceResolverNotFound {
	return &ReplaceResolverNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceResolverNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceResolverNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceResolverNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/resolvers/{name}][%d] replaceResolverNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceResolverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceResolverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceResolverDefault creates a ReplaceResolverDefault with default headers values
func NewReplaceResolverDefault(code int) *ReplaceResolverDefault {
	return &ReplaceResolverDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceResolverDefault handles this case with default header values.

General Error
*/
type ReplaceResolverDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace resolver default response
func (o *ReplaceResolverDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceResolverDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/resolvers/{name}][%d] replaceResolver default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceResolverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceResolverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
