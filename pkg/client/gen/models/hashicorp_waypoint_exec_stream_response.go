// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpWaypointExecStreamResponse hashicorp waypoint exec stream response
//
// swagger:model hashicorp.waypoint.ExecStreamResponse
type HashicorpWaypointExecStreamResponse struct {

	// exit
	Exit *HashicorpWaypointExecStreamResponseExit `json:"exit,omitempty"`

	// Open is always sent first no matter what (unless there is an error
	// in which case the stream will exit). This should be used to validate
	// that the exec process started properly.
	Open HashicorpWaypointExecStreamResponseOpen `json:"open,omitempty"`

	// output
	Output *HashicorpWaypointExecStreamResponseOutput `json:"output,omitempty"`
}

// Validate validates this hashicorp waypoint exec stream response
func (m *HashicorpWaypointExecStreamResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointExecStreamResponse) validateExit(formats strfmt.Registry) error {
	if swag.IsZero(m.Exit) { // not required
		return nil
	}

	if m.Exit != nil {
		if err := m.Exit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exit")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointExecStreamResponse) validateOutput(formats strfmt.Registry) error {
	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint exec stream response based on the context it is used
func (m *HashicorpWaypointExecStreamResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutput(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointExecStreamResponse) contextValidateExit(ctx context.Context, formats strfmt.Registry) error {

	if m.Exit != nil {
		if err := m.Exit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("exit")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpWaypointExecStreamResponse) contextValidateOutput(ctx context.Context, formats strfmt.Registry) error {

	if m.Output != nil {
		if err := m.Output.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointExecStreamResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointExecStreamResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointExecStreamResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}