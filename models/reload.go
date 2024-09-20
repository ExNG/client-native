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

// Reload HAProxy reload
//
// HAProxy reload
// Example: {"id":"2019-01-03-44","status":"in_progress"}
//
// swagger:model reload
type Reload struct {
	// id
	// Pattern: ^\d{4}-\d{2}-\d{2}-\d+$
	// +kubebuilder:validation:Pattern=`^\d{4}-\d{2}-\d{2}-\d+$`
	ID string `json:"id,omitempty"`

	// reload timestamp
	ReloadTimestamp int64 `json:"reload_timestamp,omitempty"`

	// response
	Response string `json:"response,omitempty"`

	// status
	// Enum: ["failed","in_progress","succeeded"]
	// +kubebuilder:validation:Enum="failed","in_progress","succeeded";
	Status string `json:"status,omitempty"`
}

// Validate validates this reload
func (m *Reload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Reload) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `^\d{4}-\d{2}-\d{2}-\d+$`); err != nil {
		return err
	}

	return nil
}

var reloadTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["failed","in_progress","succeeded"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reloadTypeStatusPropEnum = append(reloadTypeStatusPropEnum, v)
	}
}

const (

	// ReloadStatusFailed captures enum value "failed"
	ReloadStatusFailed string = "failed"

	// ReloadStatusInProgress captures enum value "in_progress"
	ReloadStatusInProgress string = "in_progress"

	// ReloadStatusSucceeded captures enum value "succeeded"
	ReloadStatusSucceeded string = "succeeded"
)

// prop value enum
func (m *Reload) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reloadTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Reload) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this reload based on context it is used
func (m *Reload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Reload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reload) UnmarshalBinary(b []byte) error {
	var res Reload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
