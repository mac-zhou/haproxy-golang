// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PgsqlCheckParams pgsql check params
//
// swagger:model pgsql_check_params
type PgsqlCheckParams struct {

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this pgsql check params
func (m *PgsqlCheckParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PgsqlCheckParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PgsqlCheckParams) UnmarshalBinary(b []byte) error {
	var res PgsqlCheckParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
