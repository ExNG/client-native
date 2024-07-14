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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resolver Resolver
//
// # Runtime DNS configuration
//
// swagger:model resolver
type Resolver struct {
	// accepted payload size
	// Maximum: 8192
	// Minimum: 512
	// +kubebuilder:validation:Maximum=8192
	// +kubebuilder:validation:Minimum=512
	AcceptedPayloadSize int64 `json:"accepted_payload_size,omitempty"`

	// hold nx
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	HoldNx *int64 `json:"hold_nx,omitempty"`

	// hold obsolete
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	HoldObsolete *int64 `json:"hold_obsolete,omitempty"`

	// hold other
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	HoldOther *int64 `json:"hold_other,omitempty"`

	// hold refused
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	HoldRefused *int64 `json:"hold_refused,omitempty"`

	// hold timeout
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	HoldTimeout *int64 `json:"hold_timeout,omitempty"`

	// hold valid
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	HoldValid *int64 `json:"hold_valid,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9-_.:]+$`
	Name string `json:"name"`

	// parse resolv conf
	ParseResolvConf bool `json:"parse-resolv-conf,omitempty"`

	// resolve retries
	// Minimum: 1
	// +kubebuilder:validation:Minimum=1
	ResolveRetries int64 `json:"resolve_retries,omitempty"`

	// timeout resolve
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	TimeoutResolve int64 `json:"timeout_resolve,omitempty"`

	// timeout retry
	// Minimum: 0
	// +kubebuilder:validation:Minimum=0
	TimeoutRetry int64 `json:"timeout_retry,omitempty"`
}

// Validate validates this resolver
func (m *Resolver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedPayloadSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldNx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldObsolete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldOther(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldRefused(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHoldValid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolveRetries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeoutResolve(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeoutRetry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Resolver) validateAcceptedPayloadSize(formats strfmt.Registry) error {
	if swag.IsZero(m.AcceptedPayloadSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("accepted_payload_size", "body", m.AcceptedPayloadSize, 512, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("accepted_payload_size", "body", m.AcceptedPayloadSize, 8192, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateHoldNx(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldNx) { // not required
		return nil
	}

	if err := validate.MinimumInt("hold_nx", "body", *m.HoldNx, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateHoldObsolete(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldObsolete) { // not required
		return nil
	}

	if err := validate.MinimumInt("hold_obsolete", "body", *m.HoldObsolete, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateHoldOther(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldOther) { // not required
		return nil
	}

	if err := validate.MinimumInt("hold_other", "body", *m.HoldOther, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateHoldRefused(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldRefused) { // not required
		return nil
	}

	if err := validate.MinimumInt("hold_refused", "body", *m.HoldRefused, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateHoldTimeout(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldTimeout) { // not required
		return nil
	}

	if err := validate.MinimumInt("hold_timeout", "body", *m.HoldTimeout, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateHoldValid(formats strfmt.Registry) error {
	if swag.IsZero(m.HoldValid) { // not required
		return nil
	}

	if err := validate.MinimumInt("hold_valid", "body", *m.HoldValid, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", m.Name, `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateResolveRetries(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolveRetries) { // not required
		return nil
	}

	if err := validate.MinimumInt("resolve_retries", "body", m.ResolveRetries, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateTimeoutResolve(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeoutResolve) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout_resolve", "body", m.TimeoutResolve, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Resolver) validateTimeoutRetry(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeoutRetry) { // not required
		return nil
	}

	if err := validate.MinimumInt("timeout_retry", "body", m.TimeoutRetry, 0, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this resolver based on context it is used
func (m *Resolver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Resolver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resolver) UnmarshalBinary(b []byte) error {
	var res Resolver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
