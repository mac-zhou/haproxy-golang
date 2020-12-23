// Code generated by go-swagger; DO NOT EDIT.

package backend_switching_rule

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

// ReplaceBackendSwitchingRuleReader is a Reader for the ReplaceBackendSwitchingRule structure.
type ReplaceBackendSwitchingRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceBackendSwitchingRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceBackendSwitchingRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewReplaceBackendSwitchingRuleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReplaceBackendSwitchingRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceBackendSwitchingRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewReplaceBackendSwitchingRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceBackendSwitchingRuleOK creates a ReplaceBackendSwitchingRuleOK with default headers values
func NewReplaceBackendSwitchingRuleOK() *ReplaceBackendSwitchingRuleOK {
	return &ReplaceBackendSwitchingRuleOK{}
}

/*ReplaceBackendSwitchingRuleOK handles this case with default header values.

Backend Switching Rule replaced
*/
type ReplaceBackendSwitchingRuleOK struct {
	Payload *models.BackendSwitchingRule
}

func (o *ReplaceBackendSwitchingRuleOK) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleOK  %+v", 200, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleOK) GetPayload() *models.BackendSwitchingRule {
	return o.Payload
}

func (o *ReplaceBackendSwitchingRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackendSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendSwitchingRuleAccepted creates a ReplaceBackendSwitchingRuleAccepted with default headers values
func NewReplaceBackendSwitchingRuleAccepted() *ReplaceBackendSwitchingRuleAccepted {
	return &ReplaceBackendSwitchingRuleAccepted{}
}

/*ReplaceBackendSwitchingRuleAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type ReplaceBackendSwitchingRuleAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.BackendSwitchingRule
}

func (o *ReplaceBackendSwitchingRuleAccepted) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleAccepted  %+v", 202, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleAccepted) GetPayload() *models.BackendSwitchingRule {
	return o.Payload
}

func (o *ReplaceBackendSwitchingRuleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.BackendSwitchingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBackendSwitchingRuleBadRequest creates a ReplaceBackendSwitchingRuleBadRequest with default headers values
func NewReplaceBackendSwitchingRuleBadRequest() *ReplaceBackendSwitchingRuleBadRequest {
	return &ReplaceBackendSwitchingRuleBadRequest{
		ConfigurationVersion: 0,
	}
}

/*ReplaceBackendSwitchingRuleBadRequest handles this case with default header values.

Bad request
*/
type ReplaceBackendSwitchingRuleBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceBackendSwitchingRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleBadRequest  %+v", 400, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceBackendSwitchingRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceBackendSwitchingRuleNotFound creates a ReplaceBackendSwitchingRuleNotFound with default headers values
func NewReplaceBackendSwitchingRuleNotFound() *ReplaceBackendSwitchingRuleNotFound {
	return &ReplaceBackendSwitchingRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*ReplaceBackendSwitchingRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type ReplaceBackendSwitchingRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *ReplaceBackendSwitchingRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRuleNotFound  %+v", 404, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceBackendSwitchingRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewReplaceBackendSwitchingRuleDefault creates a ReplaceBackendSwitchingRuleDefault with default headers values
func NewReplaceBackendSwitchingRuleDefault(code int) *ReplaceBackendSwitchingRuleDefault {
	return &ReplaceBackendSwitchingRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*ReplaceBackendSwitchingRuleDefault handles this case with default header values.

General Error
*/
type ReplaceBackendSwitchingRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the replace backend switching rule default response
func (o *ReplaceBackendSwitchingRuleDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceBackendSwitchingRuleDefault) Error() string {
	return fmt.Sprintf("[PUT /services/haproxy/configuration/backend_switching_rules/{index}][%d] replaceBackendSwitchingRule default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceBackendSwitchingRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReplaceBackendSwitchingRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
