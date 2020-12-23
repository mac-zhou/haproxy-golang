// Code generated by go-swagger; DO NOT EDIT.

package acl

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

// GetAclsReader is a Reader for the GetAcls structure.
type GetAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAclsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAclsOK creates a GetAclsOK with default headers values
func NewGetAclsOK() *GetAclsOK {
	return &GetAclsOK{}
}

/*GetAclsOK handles this case with default header values.

Successful operation
*/
type GetAclsOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetAclsOKBody
}

func (o *GetAclsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/acls][%d] getAclsOK  %+v", 200, o.Payload)
}

func (o *GetAclsOK) GetPayload() *GetAclsOKBody {
	return o.Payload
}

func (o *GetAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetAclsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAclsDefault creates a GetAclsDefault with default headers values
func NewGetAclsDefault(code int) *GetAclsDefault {
	return &GetAclsDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetAclsDefault handles this case with default header values.

General Error
*/
type GetAclsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get acls default response
func (o *GetAclsDefault) Code() int {
	return o._statusCode
}

func (o *GetAclsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/acls][%d] getAcls default  %+v", o._statusCode, o.Payload)
}

func (o *GetAclsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAclsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetAclsOKBody get acls o k body
swagger:model GetAclsOKBody
*/
type GetAclsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.Acls `json:"data"`
}

// Validate validates this get acls o k body
func (o *GetAclsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAclsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getAclsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getAclsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAclsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAclsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAclsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
