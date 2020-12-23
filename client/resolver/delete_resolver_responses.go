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

	"mac_zhou.exe/go/pkg/mod/github.com/mac-zhou/haproxy-golang/models"
)

// DeleteResolverReader is a Reader for the DeleteResolver structure.
type DeleteResolverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResolverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteResolverAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteResolverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteResolverNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteResolverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteResolverAccepted creates a DeleteResolverAccepted with default headers values
func NewDeleteResolverAccepted() *DeleteResolverAccepted {
	return &DeleteResolverAccepted{}
}

/*DeleteResolverAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteResolverAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteResolverAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverAccepted ", 202)
}

func (o *DeleteResolverAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteResolverNoContent creates a DeleteResolverNoContent with default headers values
func NewDeleteResolverNoContent() *DeleteResolverNoContent {
	return &DeleteResolverNoContent{}
}

/*DeleteResolverNoContent handles this case with default header values.

Resolver deleted
*/
type DeleteResolverNoContent struct {
}

func (o *DeleteResolverNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverNoContent ", 204)
}

func (o *DeleteResolverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResolverNotFound creates a DeleteResolverNotFound with default headers values
func NewDeleteResolverNotFound() *DeleteResolverNotFound {
	return &DeleteResolverNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteResolverNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteResolverNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteResolverNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolverNotFound  %+v", 404, o.Payload)
}

func (o *DeleteResolverNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteResolverNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteResolverDefault creates a DeleteResolverDefault with default headers values
func NewDeleteResolverDefault(code int) *DeleteResolverDefault {
	return &DeleteResolverDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteResolverDefault handles this case with default header values.

General Error
*/
type DeleteResolverDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete resolver default response
func (o *DeleteResolverDefault) Code() int {
	return o._statusCode
}

func (o *DeleteResolverDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/resolvers/{name}][%d] deleteResolver default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteResolverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteResolverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
