// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

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

// GetTCPRequestRuleReader is a Reader for the GetTCPRequestRule structure.
type GetTCPRequestRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTCPRequestRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTCPRequestRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTCPRequestRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetTCPRequestRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTCPRequestRuleOK creates a GetTCPRequestRuleOK with default headers values
func NewGetTCPRequestRuleOK() *GetTCPRequestRuleOK {
	return &GetTCPRequestRuleOK{}
}

/*GetTCPRequestRuleOK handles this case with default header values.

Successful operation
*/
type GetTCPRequestRuleOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetTCPRequestRuleOKBody
}

func (o *GetTCPRequestRuleOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTcpRequestRuleOK  %+v", 200, o.Payload)
}

func (o *GetTCPRequestRuleOK) GetPayload() *GetTCPRequestRuleOKBody {
	return o.Payload
}

func (o *GetTCPRequestRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetTCPRequestRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTCPRequestRuleNotFound creates a GetTCPRequestRuleNotFound with default headers values
func NewGetTCPRequestRuleNotFound() *GetTCPRequestRuleNotFound {
	return &GetTCPRequestRuleNotFound{
		ConfigurationVersion: 0,
	}
}

/*GetTCPRequestRuleNotFound handles this case with default header values.

The specified resource was not found
*/
type GetTCPRequestRuleNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *GetTCPRequestRuleNotFound) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTcpRequestRuleNotFound  %+v", 404, o.Payload)
}

func (o *GetTCPRequestRuleNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPRequestRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetTCPRequestRuleDefault creates a GetTCPRequestRuleDefault with default headers values
func NewGetTCPRequestRuleDefault(code int) *GetTCPRequestRuleDefault {
	return &GetTCPRequestRuleDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetTCPRequestRuleDefault handles this case with default header values.

General Error
*/
type GetTCPRequestRuleDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get TCP request rule default response
func (o *GetTCPRequestRuleDefault) Code() int {
	return o._statusCode
}

func (o *GetTCPRequestRuleDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/tcp_request_rules/{index}][%d] getTCPRequestRule default  %+v", o._statusCode, o.Payload)
}

func (o *GetTCPRequestRuleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTCPRequestRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetTCPRequestRuleOKBody get TCP request rule o k body
swagger:model GetTCPRequestRuleOKBody
*/
type GetTCPRequestRuleOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	Data *models.TCPRequestRule `json:"data,omitempty"`
}

// Validate validates this get TCP request rule o k body
func (o *GetTCPRequestRuleOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTCPRequestRuleOKBody) validateData(formats strfmt.Registry) error {

	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getTcpRequestRuleOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTCPRequestRuleOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTCPRequestRuleOKBody) UnmarshalBinary(b []byte) error {
	var res GetTCPRequestRuleOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}