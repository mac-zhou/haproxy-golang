// Code generated by go-swagger; DO NOT EDIT.

package maps

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

// CreateRuntimeMapReader is a Reader for the CreateRuntimeMap structure.
type CreateRuntimeMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRuntimeMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRuntimeMapCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateRuntimeMapAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRuntimeMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateRuntimeMapConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateRuntimeMapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRuntimeMapCreated creates a CreateRuntimeMapCreated with default headers values
func NewCreateRuntimeMapCreated() *CreateRuntimeMapCreated {
	return &CreateRuntimeMapCreated{}
}

/*CreateRuntimeMapCreated handles this case with default header values.

Map file created with its entries
*/
type CreateRuntimeMapCreated struct {
	Payload *models.Map
}

func (o *CreateRuntimeMapCreated) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps][%d] createRuntimeMapCreated  %+v", 201, o.Payload)
}

func (o *CreateRuntimeMapCreated) GetPayload() *models.Map {
	return o.Payload
}

func (o *CreateRuntimeMapCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Map)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuntimeMapAccepted creates a CreateRuntimeMapAccepted with default headers values
func NewCreateRuntimeMapAccepted() *CreateRuntimeMapAccepted {
	return &CreateRuntimeMapAccepted{}
}

/*CreateRuntimeMapAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type CreateRuntimeMapAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string

	Payload *models.Map
}

func (o *CreateRuntimeMapAccepted) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps][%d] createRuntimeMapAccepted  %+v", 202, o.Payload)
}

func (o *CreateRuntimeMapAccepted) GetPayload() *models.Map {
	return o.Payload
}

func (o *CreateRuntimeMapAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	o.Payload = new(models.Map)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuntimeMapBadRequest creates a CreateRuntimeMapBadRequest with default headers values
func NewCreateRuntimeMapBadRequest() *CreateRuntimeMapBadRequest {
	return &CreateRuntimeMapBadRequest{
		ConfigurationVersion: 0,
	}
}

/*CreateRuntimeMapBadRequest handles this case with default header values.

Bad request
*/
type CreateRuntimeMapBadRequest struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateRuntimeMapBadRequest) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps][%d] createRuntimeMapBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRuntimeMapBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRuntimeMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRuntimeMapConflict creates a CreateRuntimeMapConflict with default headers values
func NewCreateRuntimeMapConflict() *CreateRuntimeMapConflict {
	return &CreateRuntimeMapConflict{
		ConfigurationVersion: 0,
	}
}

/*CreateRuntimeMapConflict handles this case with default header values.

The specified resource already exists
*/
type CreateRuntimeMapConflict struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *CreateRuntimeMapConflict) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps][%d] createRuntimeMapConflict  %+v", 409, o.Payload)
}

func (o *CreateRuntimeMapConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRuntimeMapConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRuntimeMapDefault creates a CreateRuntimeMapDefault with default headers values
func NewCreateRuntimeMapDefault(code int) *CreateRuntimeMapDefault {
	return &CreateRuntimeMapDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*CreateRuntimeMapDefault handles this case with default header values.

General Error
*/
type CreateRuntimeMapDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the create runtime map default response
func (o *CreateRuntimeMapDefault) Code() int {
	return o._statusCode
}

func (o *CreateRuntimeMapDefault) Error() string {
	return fmt.Sprintf("[POST /services/haproxy/runtime/maps][%d] createRuntimeMap default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRuntimeMapDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRuntimeMapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
