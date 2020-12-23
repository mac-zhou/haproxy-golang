// Code generated by go-swagger; DO NOT EDIT.

package tcp_response_rule

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

// GetTCPResponseRulesReader is a Reader for the GetTCPResponseRules structure.
type GetTCPResponseRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPResponseRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPResponseRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTCPResponseRulesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTCPResponseRulesOK creates a GetTCPResponseRulesOK with default headers values
func NewGetTCPResponseRulesOK() *GetTCPResponseRulesOK {
	return &GetTCPResponseRulesOK{}
}

/*GetTCPResponseRulesOK handles this case with default header values.

Successful operation
*/
type GetTCPResponseRulesOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetTCPResponseRulesOKBody
}

func (o *GetTCPResponseRulesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules][%d] getTcpResponseRulesOK  %+v", 200, o.Payload)
}

func (o *GetTCPResponseRulesOK) GetPayload() *GetTCPResponseRulesOKBody {
	return o.Payload
}

func (o *GetTCPResponseRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetTCPResponseRulesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPResponseRulesDefault creates a GetTCPResponseRulesDefault with default headers values
func NewGetTCPResponseRulesDefault(code int) *GetTCPResponseRulesDefault {
	return &GetTCPResponseRulesDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetTCPResponseRulesDefault handles this case with default header values.

General Error
*/
type GetTCPResponseRulesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get TCP response rules default response
func (o *GetTCPResponseRulesDefault) Code() int {
	return o._statusCode
}

func (o *GetTCPResponseRulesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules][%d] getTCPResponseRules default  %+v", o._statusCode, o.Payload)
}

func (o *GetTCPResponseRulesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPResponseRulesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetTCPResponseRulesOKBody get TCP response rules o k body
swagger:model GetTCPResponseRulesOKBody
*/
type GetTCPResponseRulesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.TCPResponseRules `json:"data"`
}

// Validate validates this get TCP response rules o k body
func (o *GetTCPResponseRulesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPResponseRulesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getTcpResponseRulesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getTcpResponseRulesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPResponseRulesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPResponseRulesOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPResponseRulesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
