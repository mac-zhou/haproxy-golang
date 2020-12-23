// Code generated by go-swagger; DO NOT EDIT.

package frontend

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

// GetFrontendReader is a Reader for the GetFrontend structure.
type GetFrontendReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFrontendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFrontendOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetFrontendNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetFrontendDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFrontendOK creates a GetFrontendOK with default headers values
func NewGetFrontendOK() *GetFrontendOK {
	return &GetFrontendOK{}
}

/*GetFrontendOK handles this case with default header values.

Successful operation
*/
type GetFrontendOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetFrontendOKBody
}

func (o *GetFrontendOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/frontends/{name}][%d] getFrontendOK  %+v", 200, o.Payload)
}

func (o *GetFrontendOK) GetPayload() *GetFrontendOKBody {
	return o.Payload
}

func (o *GetFrontendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetFrontendOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFrontendNotFound creates a GetFrontendNotFound with default headers values
func NewGetFrontendNotFound() *GetFrontendNotFound {
	return &GetFrontendNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetFrontendNotFound handles this case with default header values.

The specified resource was not found
*/
type GetFrontendNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetFrontendNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/frontends/{name}][%d] getFrontendNotFound  %+v", 404, o.Payload)
}

func (o *GetFrontendNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFrontendNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFrontendDefault creates a GetFrontendDefault with default headers values
func NewGetFrontendDefault(code int) *GetFrontendDefault {
	return &GetFrontendDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetFrontendDefault handles this case with default header values.

General Error
*/
type GetFrontendDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get frontend default response
func (o *GetFrontendDefault) Code() int {
	return o._statusCode
}

func (o *GetFrontendDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/frontends/{name}][%d] getFrontend default  %+v", o._statusCode, o.Payload)
}

func (o *GetFrontendDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFrontendDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetFrontendOKBody get frontend o k body
swagger:model GetFrontendOKBody
*/
type GetFrontendOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.Frontend `json:"data,omitempty"`
}

// Validate validates this get frontend o k body
func (o *GetFrontendOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFrontendOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getFrontendOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetFrontendOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetFrontendOKBody) UnmarshalBinary(b []byte) error {
	var res GetFrontendOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
