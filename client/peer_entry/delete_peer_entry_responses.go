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

	"mac_zhou.exe/go/pkg/mod/github.com/mac-zhou/haproxy-golang/models"
)

// DeletePeerEntryReader is a Reader for the DeletePeerEntry structure.
type DeletePeerEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePeerEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeletePeerEntryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeletePeerEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeletePeerEntryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeletePeerEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePeerEntryAccepted creates a DeletePeerEntryAccepted with default headers values
func NewDeletePeerEntryAccepted() *DeletePeerEntryAccepted {
	return &DeletePeerEntryAccepted{}
}

/*DeletePeerEntryAccepted handles this case with default header values.

Configuration change accepted and reload requested
*/
type DeletePeerEntryAccepted struct {
	/*ID of the requested reload
	 */
	ReloadID string
}

func (o *DeletePeerEntryAccepted) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_entries/{name}][%d] deletePeerEntryAccepted ", 202)
}

func (o *DeletePeerEntryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Reload-ID
	o.ReloadID = response.GetHeader("Reload-ID")

	return nil
}

// NewDeletePeerEntryNoContent creates a DeletePeerEntryNoContent with default headers values
func NewDeletePeerEntryNoContent() *DeletePeerEntryNoContent {
	return &DeletePeerEntryNoContent{}
}

/*DeletePeerEntryNoContent handles this case with default header values.

PeerEntry deleted
*/
type DeletePeerEntryNoContent struct {
}

func (o *DeletePeerEntryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_entries/{name}][%d] deletePeerEntryNoContent ", 204)
}

func (o *DeletePeerEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePeerEntryNotFound creates a DeletePeerEntryNotFound with default headers values
func NewDeletePeerEntryNotFound() *DeletePeerEntryNotFound {
	return &DeletePeerEntryNotFound{
		ConfigurationVersion: 0,
	}
}

/*DeletePeerEntryNotFound handles this case with default header values.

The specified resource was not found
*/
type DeletePeerEntryNotFound struct {
	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

func (o *DeletePeerEntryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_entries/{name}][%d] deletePeerEntryNotFound  %+v", 404, o.Payload)
}

func (o *DeletePeerEntryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerEntryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeletePeerEntryDefault creates a DeletePeerEntryDefault with default headers values
func NewDeletePeerEntryDefault(code int) *DeletePeerEntryDefault {
	return &DeletePeerEntryDefault{
		_statusCode:          code,
		ConfigurationVersion: 0,
	}
}

/*DeletePeerEntryDefault handles this case with default header values.

General Error
*/
type DeletePeerEntryDefault struct {
	_statusCode int

	/*Configuration file version
	 */
	ConfigurationVersion int64

	Payload *models.Error
}

// Code gets the status code for the delete peer entry default response
func (o *DeletePeerEntryDefault) Code() int {
	return o._statusCode
}

func (o *DeletePeerEntryDefault) Error() string {
	return fmt.Sprintf("[DELETE /services/haproxy/configuration/peer_entries/{name}][%d] deletePeerEntry default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePeerEntryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeletePeerEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
