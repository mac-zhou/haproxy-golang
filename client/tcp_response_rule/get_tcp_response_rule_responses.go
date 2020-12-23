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

	"github.com/mac-zhou/haproxy-golang/models"
)

// GetTCPResponseRuleReader is a Reader for the GetTCPResponseRule structure.
type GetTCPResponseRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPResponseRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPResponseRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTCPResponseRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTCPResponseRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTCPResponseRuleOK creates a GetTCPResponseRuleOK with default headers values
func NewGetTCPResponseRuleOK() *GetTCPResponseRuleOK {
	return &GetTCPResponseRuleOK{}
}

/*GetTCPResponseRuleOK handles this case with default header values.

Successful operation
*/
type GetTCPResponseRuleOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetTCPResponseRuleOKBody
}

func (o *GetTCPResponseRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules/{index}][%d] getTcpResponseRuleOK  %+v", 200, o.Payload)
}

func (o *GetTCPResponseRuleOK) GetPayload() *GetTCPResponseRuleOKBody {
	return o.Payload
}

func (o *GetTCPResponseRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetTCPResponseRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPResponseRuleNotFound creates a GetTCPResponseRuleNotFound with default headers values
func NewGetTCPResponseRuleNotFound() *GetTCPResponseRuleNotFound {
	return &GetTCPResponseRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetTCPResponseRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type GetTCPResponseRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetTCPResponseRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules/{index}][%d] getTcpResponseRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetTCPResponseRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPResponseRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTCPResponseRuleDefault creates a GetTCPResponseRuleDefault with default headers values
func NewGetTCPResponseRuleDefault(code int) *GetTCPResponseRuleDefault {
	return &GetTCPResponseRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetTCPResponseRuleDefault handles this case with default header values.

General Error
*/
type GetTCPResponseRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get TCP response rule default response
func (o *GetTCPResponseRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetTCPResponseRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_response_rules/{index}][%d] getTCPResponseRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetTCPResponseRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPResponseRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetTCPResponseRuleOKBody get TCP response rule o k body
swagger:model GetTCPResponseRuleOKBody
*/
type GetTCPResponseRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.TCPResponseRule `json:"data,omitempty"`
}

// Validate validates this get TCP response rule o k body
func (o *GetTCPResponseRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPResponseRuleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTcpResponseRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPResponseRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPResponseRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPResponseRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
