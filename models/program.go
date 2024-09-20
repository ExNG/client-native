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

// Program Program
//
// HAProxy program configuration
// Example: {"command":"spoa-mirror --runtime 0 --mirror-url http://test.local","group":"mygroupname","name":"mirror","start-on-reload":"enabled","user":"myusername"}
//
// swagger:model program
type Program struct {

	// The command to be run, with flags and options.
	// Required: true
	Command *string `json:"command"`

	// The group to run the command as, if different than the HAProxy group.
	Group string `json:"group,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9-_.:]+$`
	Name string `json:"name"`

	// HAProxy stops and recreates child programs at reload.
	// Enum: ["enabled","disabled"]
	// +kubebuilder:validation:Enum="enabled","disabled";
	StartOnReload string `json:"start-on-reload,omitempty"`

	// The user to run the command as, if different than the HAProxy user.
	User string `json:"user,omitempty"`
}

// Validate validates this program
func (m *Program) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartOnReload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Program) validateCommand(formats strfmt.Registry) error {

	if err := validate.Required("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

func (m *Program) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

var programTypeStartOnReloadPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		programTypeStartOnReloadPropEnum = append(programTypeStartOnReloadPropEnum, v)
	}
}

const (

	// ProgramStartOnReloadEnabled captures enum value "enabled"
	ProgramStartOnReloadEnabled string = "enabled"

	// ProgramStartOnReloadDisabled captures enum value "disabled"
	ProgramStartOnReloadDisabled string = "disabled"
)

// prop value enum
func (m *Program) validateStartOnReloadEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, programTypeStartOnReloadPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Program) validateStartOnReload(formats strfmt.Registry) error {
	if swag.IsZero(m.StartOnReload) { // not required
		return nil
	}

	// value enum
	if err := m.validateStartOnReloadEnum("start-on-reload", "body", m.StartOnReload); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this program based on context it is used
func (m *Program) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Program) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Program) UnmarshalBinary(b []byte) error {
	var res Program
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
