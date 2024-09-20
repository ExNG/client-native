// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FCGISetParam Sets a FastCGI parameter to pass to this application.
// Its value, defined by <format> can take a formatted string, the same as the log directive.
// Optionally, you can follow it with an ACL-based condition, in which case the FastCGI application evaluates it only if the condition is true.
//
// swagger:model fcgiSetParam
type FCGISetParam struct {
	// cond
	// Enum: ["if","unless"]
	// +kubebuilder:validation:Enum="if","unless";
	Cond string `json:"cond,omitempty"`

	// cond test
	CondTest string `json:"cond_test,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this fcgi set param
func (m *FCGISetParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCond(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fcgiSetParamTypeCondPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["if","unless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcgiSetParamTypeCondPropEnum = append(fcgiSetParamTypeCondPropEnum, v)
	}
}

const (

	// FCGISetParamCondIf captures enum value "if"
	FCGISetParamCondIf string = "if"

	// FCGISetParamCondUnless captures enum value "unless"
	FCGISetParamCondUnless string = "unless"
)

// prop value enum
func (m *FCGISetParam) validateCondEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcgiSetParamTypeCondPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FCGISetParam) validateCond(formats strfmt.Registry) error {
	if swag.IsZero(m.Cond) { // not required
		return nil
	}

	// value enum
	if err := m.validateCondEnum("cond", "body", m.Cond); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this fcgi set param based on context it is used
func (m *FCGISetParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FCGISetParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FCGISetParam) UnmarshalBinary(b []byte) error {
	var res FCGISetParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
