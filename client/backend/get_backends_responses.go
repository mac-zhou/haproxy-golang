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
	"github.com/go-openapi/validate"

	"mac_zhou.exe/go/pkg/mod/github.com/mac-zhou/haproxy-golang/models"
)

// GetBackendsReader is a Reader for the GetBackends structure.
type GetBackendsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackendsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackendsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackendsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackendsOK creates a GetBackendsOK with default headers values
func NewGetBackendsOK() *GetBackendsOK {
	return &GetBackendsOK{}
}

/*GetBackendsOK handles this case with default header values.

Successful operation
*/
type GetBackendsOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetBackendsOKBody
}

func (o *GetBackendsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backends][%d] getBackendsOK  %+v", 200, o.Payload)
}

func (o *GetBackendsOK) GetPayload() *GetBackendsOKBody {
	return o.Payload
}

func (o *GetBackendsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetBackendsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackendsDefault creates a GetBackendsDefault with default headers values
func NewGetBackendsDefault(code int) *GetBackendsDefault {
	return &GetBackendsDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetBackendsDefault handles this case with default header values.

General Error
*/
type GetBackendsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get backends default response
func (o *GetBackendsDefault) Code() int {
	return o._statusCode
}

func (o *GetBackendsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/backends][%d] getBackends default  %+v", o._statusCode, o.Payload)
}

func (o *GetBackendsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetBackendsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetBackendsOKBody get backends o k body
swagger:model GetBackendsOKBody
*/
type GetBackendsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Backends `json:"data"`
}

// Validate validates this get backends o k body
func (o *GetBackendsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetBackendsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getBackendsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getBackendsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetBackendsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetBackendsOKBody) UnmarshalBinary(b []byte) error {
	var res GetBackendsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
