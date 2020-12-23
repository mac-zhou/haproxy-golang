// Code generated by go-swagger; DO NOT EDIT.

package peer_entry

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

// GetPeerEntriesReader is a Reader for the GetPeerEntries structure.
type GetPeerEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeerEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeerEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPeerEntriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPeerEntriesOK creates a GetPeerEntriesOK with default headers values
func NewGetPeerEntriesOK() *GetPeerEntriesOK {
	return &GetPeerEntriesOK{}
}

/*GetPeerEntriesOK handles this case with default header values.

Successful operation
*/
type GetPeerEntriesOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetPeerEntriesOKBody
}

func (o *GetPeerEntriesOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries][%d] getPeerEntriesOK  %+v", 200, o.Payload)
}

func (o *GetPeerEntriesOK) GetPayload() *GetPeerEntriesOKBody {
	return o.Payload
}

func (o *GetPeerEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetPeerEntriesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeerEntriesDefault creates a GetPeerEntriesDefault with default headers values
func NewGetPeerEntriesDefault(code int) *GetPeerEntriesDefault {
	return &GetPeerEntriesDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetPeerEntriesDefault handles this case with default header values.

General Error
*/
type GetPeerEntriesDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get peer entries default response
func (o *GetPeerEntriesDefault) Code() int {
	return o._statusCode
}

func (o *GetPeerEntriesDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_entries][%d] getPeerEntries default  %+v", o._statusCode, o.Payload)
}

func (o *GetPeerEntriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerEntriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPeerEntriesOKBody get peer entries o k body
swagger:model GetPeerEntriesOKBody
*/
type GetPeerEntriesOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.PeerEntries `json:"data"`
}

// Validate validates this get peer entries o k body
func (o *GetPeerEntriesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerEntriesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPeerEntriesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerEntriesOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerEntriesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerEntriesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerEntriesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}