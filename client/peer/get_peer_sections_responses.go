// Code generated by go-swagger; DO NOT EDIT.

package peer

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

	"github.com/mac-zhou/haproxy-golang/models"
)

// GetPeerSectionsReader is a Reader for the GetPeerSections structure.
type GetPeerSectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPeerSectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPeerSectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetPeerSectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPeerSectionsOK creates a GetPeerSectionsOK with default headers values
func NewGetPeerSectionsOK() *GetPeerSectionsOK {
	return &GetPeerSectionsOK{}
}

/*GetPeerSectionsOK handles this case with default header values.

Successful operation
*/
type GetPeerSectionsOK struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *GetPeerSectionsOKBody
}

func (o *GetPeerSectionsOK) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section][%d] getPeerSectionsOK  %+v", 200, o.Payload)
}

func (o *GetPeerSectionsOK) GetPayload() *GetPeerSectionsOKBody {
	return o.Payload
}

func (o *GetPeerSectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Configuration-Version
	configurationVersion, err := swag.ConvertInt64(response.GetHeader("Configuration-Version"))
	if err != nil {
		return errors.InvalidType("Configuration-Version", "header", "int64", response.GetHeader("Configuration-Version"))
	}
	o.ConfigurationVersion = configurationVersion

	o.Payload = new(GetPeerSectionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeerSectionsDefault creates a GetPeerSectionsDefault with default headers values
func NewGetPeerSectionsDefault(code int) *GetPeerSectionsDefault {
	return &GetPeerSectionsDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*GetPeerSectionsDefault handles this case with default header values.

General Error
*/
type GetPeerSectionsDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the get peer sections default response
func (o *GetPeerSectionsDefault) Code() int {
	return o._statusCode
}

func (o *GetPeerSectionsDefault) Error() string {
	return fmt.Sprintf("[GET /services/haproxy/configuration/peer_section][%d] getPeerSections default  %+v", o._statusCode, o.Payload)
}

func (o *GetPeerSectionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetPeerSectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetPeerSectionsOKBody get peer sections o k body
swagger:model GetPeerSectionsOKBody
*/
type GetPeerSectionsOKBody struct {

	// version
	Version int64 `json:"_version,omitempty"`

	// data
	// Required: true
	Data models.PeerSections `json:"data"`
}

// Validate validates this get peer sections o k body
func (o *GetPeerSectionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPeerSectionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getPeerSectionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	if err := o.Data.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("getPeerSectionsOK" + "." + "data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPeerSectionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPeerSectionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetPeerSectionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
