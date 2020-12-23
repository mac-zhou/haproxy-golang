// Code generated by go-swagger; DO NOT EDIT.

package server_switching_rule

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

// DeleteServerSwitchingRuleReader is a Reader for the DeleteServerSwitchingRule structure.
type DeleteServerSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteServerSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteServerSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteServerSwitchingRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteServerSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteServerSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteServerSwitchingRuleAccepted creates a DeleteServerSwitchingRuleAccepted with default headers values
func NewDeleteServerSwitchingRuleAccepted() *DeleteServerSwitchingRuleAccepted {
	return &DeleteServerSwitchingRuleAccepted{}
}

/*DeleteServerSwitchingRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeleteServerSwitchingRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeleteServerSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_switching_rules/{index}][%d] deleteServerSwitchingRuleAccepted ", 202)
}

func (o *DeleteServerSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeleteServerSwitchingRuleNoContent creates a DeleteServerSwitchingRuleNoContent with default headers values
func NewDeleteServerSwitchingRuleNoContent() *DeleteServerSwitchingRuleNoContent {
	return &DeleteServerSwitchingRuleNoContent{}
}

/*DeleteServerSwitchingRuleNoContent handles this case with default header values.

Server Switching Rule deleted
*/
type DeleteServerSwitchingRuleNoContent struct {
}

func (o *DeleteServerSwitchingRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_switching_rules/{index}][%d] deleteServerSwitchingRuleNoContent ", 204)
}

func (o *DeleteServerSwitchingRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteServerSwitchingRuleNotFound creates a DeleteServerSwitchingRuleNotFound with default headers values
func NewDeleteServerSwitchingRuleNotFound() *DeleteServerSwitchingRuleNotFound {
	return &DeleteServerSwitchingRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeleteServerSwitchingRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type DeleteServerSwitchingRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeleteServerSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_switching_rules/{index}][%d] deleteServerSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteServerSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteServerSwitchingRuleDefault creates a DeleteServerSwitchingRuleDefault with default headers values
func NewDeleteServerSwitchingRuleDefault(code int) *DeleteServerSwitchingRuleDefault {
	return &DeleteServerSwitchingRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeleteServerSwitchingRuleDefault handles this case with default header values.

General Error
*/
type DeleteServerSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete server switching rule default response
func (o *DeleteServerSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *DeleteServerSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/server_switching_rules/{index}][%d] deleteServerSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteServerSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteServerSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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