// Code generated by go-swagger; DO NOT EDIT.

package http_response_rule

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

// DeleteHTTPResponseRuleReader is a Reader for the DeleteHTTPResponseRule structure.
type DeleteHTTPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHTTPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteHTTPResponseRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteHTTPResponseRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteHTTPResponseRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteHTTPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteHTTPResponseRuleAccepted creates a DeleteHTTPResponseRuleAccepted with default headers values
func NewDeleteHTTPResponseRuleAccepted() *DeleteHTTPResponseRuleAccepted {
	return &DeleteHTTPResponseRuleAccepted{}
}

/*DeleteHTTPResponseRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteHTTPResponseRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteHTTPResponseRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_response_rules/{index}][%d] deleteHttpResponseRuleAccepted ", 202)
}

func (o *DeleteHTTPResponseRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteHTTPResponseRuleNoContent creates a DeleteHTTPResponseRuleNoContent with default headers values
func NewDeleteHTTPResponseRuleNoContent() *DeleteHTTPResponseRuleNoContent {
	return &DeleteHTTPResponseRuleNoContent{}
}

/*DeleteHTTPResponseRuleNoContent handles this case with default header values.

HTTP Response Rule deleted
*/
type DeleteHTTPResponseRuleNoContent struct {
}

func (o *DeleteHTTPResponseRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_response_rules/{index}][%d] deleteHttpResponseRuleNoContent ", 204)
}

func (o *DeleteHTTPResponseRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHTTPResponseRuleNotFound creates a DeleteHTTPResponseRuleNotFound with default headers values
func NewDeleteHTTPResponseRuleNotFound() *DeleteHTTPResponseRuleNotFound {
	return &DeleteHTTPResponseRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteHTTPResponseRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteHTTPResponseRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteHTTPResponseRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_response_rules/{index}][%d] deleteHttpResponseRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteHTTPResponseRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHTTPResponseRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteHTTPResponseRuleDefault creates a DeleteHTTPResponseRuleDefault with default headers values
func NewDeleteHTTPResponseRuleDefault(code int) *DeleteHTTPResponseRuleDefault {
	return &DeleteHTTPResponseRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteHTTPResponseRuleDefault handles this case with default header values.

General Error
*/
type DeleteHTTPResponseRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete HTTP response rule default response
func (o *DeleteHTTPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteHTTPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/http_response_rules/{index}][%d] deleteHTTPResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteHTTPResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteHTTPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
