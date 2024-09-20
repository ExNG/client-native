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

// PerformanceOptions performance options
//
// swagger:model performance_options
type PerformanceOptions struct {

	// busy polling
	BusyPolling bool `json:"busy_polling,omitempty"`

	// max spread checks
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	MaxSpreadChecks *int64 `json:"max_spread_checks,omitempty"`

	// maxcompcpuusage
	Maxcompcpuusage int64 `json:"maxcompcpuusage,omitempty"`

	// maxcomprate
	Maxcomprate int64 `json:"maxcomprate,omitempty"`

	// maxconn
	Maxconn int64 `json:"maxconn,omitempty"`

	// maxconnrate
	Maxconnrate int64 `json:"maxconnrate,omitempty"`

	// maxpipes
	Maxpipes int64 `json:"maxpipes,omitempty"`

	// maxsessrate
	Maxsessrate int64 `json:"maxsessrate,omitempty"`

	// maxzlibmem
	Maxzlibmem int64 `json:"maxzlibmem,omitempty"`

	// noepoll
	Noepoll bool `json:"noepoll,omitempty"`

	// noevports
	Noevports bool `json:"noevports,omitempty"`

	// nogetaddrinfo
	Nogetaddrinfo bool `json:"nogetaddrinfo,omitempty"`

	// nokqueue
	Nokqueue bool `json:"nokqueue,omitempty"`

	// nopoll
	Nopoll bool `json:"nopoll,omitempty"`

	// noreuseport
	Noreuseport bool `json:"noreuseport,omitempty"`

	// nosplice
	Nosplice bool `json:"nosplice,omitempty"`

	// profiling memory
	// Enum: ["enabled","disabled"]
	// +kubebuilder:validation:Enum="enabled","disabled";
	ProfilingMemory string `json:"profiling_memory,omitempty"`

	// profiling tasks
	// Enum: ["auto","enabled","disabled"]
	// +kubebuilder:validation:Enum="auto","enabled","disabled";
	ProfilingTasks string `json:"profiling_tasks,omitempty"`

	// server state base
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	ServerStateBase string `json:"server_state_base,omitempty"`

	// server state file
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	ServerStateFile string `json:"server_state_file,omitempty"`

	// spread checks
	SpreadChecks int64 `json:"spread_checks,omitempty"`

	// thread hard limit
	ThreadHardLimit *int64 `json:"thread_hard_limit,omitempty"`
}

// Validate validates this performance options
func (m *PerformanceOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMaxSpreadChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfilingMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfilingTasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStateBase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerStateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PerformanceOptions) validateMaxSpreadChecks(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxSpreadChecks) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_spread_checks", "body", *m.MaxSpreadChecks, 0, false); err != nil {
		return err
	}

	return nil
}

var performanceOptionsTypeProfilingMemoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceOptionsTypeProfilingMemoryPropEnum = append(performanceOptionsTypeProfilingMemoryPropEnum, v)
	}
}

const (

	// PerformanceOptionsProfilingMemoryEnabled captures enum value "enabled"
	PerformanceOptionsProfilingMemoryEnabled string = "enabled"

	// PerformanceOptionsProfilingMemoryDisabled captures enum value "disabled"
	PerformanceOptionsProfilingMemoryDisabled string = "disabled"
)

// prop value enum
func (m *PerformanceOptions) validateProfilingMemoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceOptionsTypeProfilingMemoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceOptions) validateProfilingMemory(formats strfmt.Registry) error {
	if swag.IsZero(m.ProfilingMemory) { // not required
		return nil
	}

	// value enum
	if err := m.validateProfilingMemoryEnum("profiling_memory", "body", m.ProfilingMemory); err != nil {
		return err
	}

	return nil
}

var performanceOptionsTypeProfilingTasksPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","enabled","disabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		performanceOptionsTypeProfilingTasksPropEnum = append(performanceOptionsTypeProfilingTasksPropEnum, v)
	}
}

const (

	// PerformanceOptionsProfilingTasksAuto captures enum value "auto"
	PerformanceOptionsProfilingTasksAuto string = "auto"

	// PerformanceOptionsProfilingTasksEnabled captures enum value "enabled"
	PerformanceOptionsProfilingTasksEnabled string = "enabled"

	// PerformanceOptionsProfilingTasksDisabled captures enum value "disabled"
	PerformanceOptionsProfilingTasksDisabled string = "disabled"
)

// prop value enum
func (m *PerformanceOptions) validateProfilingTasksEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, performanceOptionsTypeProfilingTasksPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PerformanceOptions) validateProfilingTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.ProfilingTasks) { // not required
		return nil
	}

	// value enum
	if err := m.validateProfilingTasksEnum("profiling_tasks", "body", m.ProfilingTasks); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceOptions) validateServerStateBase(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerStateBase) { // not required
		return nil
	}

	if err := validate.Pattern("server_state_base", "body", m.ServerStateBase, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *PerformanceOptions) validateServerStateFile(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerStateFile) { // not required
		return nil
	}

	if err := validate.Pattern("server_state_file", "body", m.ServerStateFile, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this performance options based on context it is used
func (m *PerformanceOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PerformanceOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PerformanceOptions) UnmarshalBinary(b []byte) error {
	var res PerformanceOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
