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

// ConfigStickTable config stick table
//
// swagger:model config_stick_table
type ConfigStickTable struct {

	// expire
	Expire *int64 `json:"expire,omitempty"`

	// keylen
	Keylen *int64 `json:"keylen,omitempty"`

	// nopurge
	Nopurge bool `json:"nopurge,omitempty"`

	// peers
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	Peers string `json:"peers,omitempty"`

	// size
	Size *int64 `json:"size,omitempty"`

	// srvkey
	// Enum: [addr name]
	// +kubebuilder:validation:Enum=addr;name;
	Srvkey *string `json:"srvkey,omitempty"`

	// store
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	Store string `json:"store,omitempty"`

	// type
	// Enum: [ip ipv6 integer string binary]
	// +kubebuilder:validation:Enum=ip;ipv6;integer;string;binary;
	Type string `json:"type,omitempty"`

	// write to
	// Pattern: ^[^\s]+$
	// +kubebuilder:validation:Pattern=`^[^\s]+$`
	WriteTo *string `json:"write_to,omitempty"`
}

// Validate validates this config stick table
func (m *ConfigStickTable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSrvkey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWriteTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigStickTable) validatePeers(formats strfmt.Registry) error {
	if swag.IsZero(m.Peers) { // not required
		return nil
	}

	if err := validate.Pattern("peers", "body", m.Peers, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var configStickTableTypeSrvkeyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["addr","name"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configStickTableTypeSrvkeyPropEnum = append(configStickTableTypeSrvkeyPropEnum, v)
	}
}

const (

	// ConfigStickTableSrvkeyAddr captures enum value "addr"
	ConfigStickTableSrvkeyAddr string = "addr"

	// ConfigStickTableSrvkeyName captures enum value "name"
	ConfigStickTableSrvkeyName string = "name"
)

// prop value enum
func (m *ConfigStickTable) validateSrvkeyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configStickTableTypeSrvkeyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConfigStickTable) validateSrvkey(formats strfmt.Registry) error {
	if swag.IsZero(m.Srvkey) { // not required
		return nil
	}

	// value enum
	if err := m.validateSrvkeyEnum("srvkey", "body", *m.Srvkey); err != nil {
		return err
	}

	return nil
}

func (m *ConfigStickTable) validateStore(formats strfmt.Registry) error {
	if swag.IsZero(m.Store) { // not required
		return nil
	}

	if err := validate.Pattern("store", "body", m.Store, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

var configStickTableTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ip","ipv6","integer","string","binary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configStickTableTypeTypePropEnum = append(configStickTableTypeTypePropEnum, v)
	}
}

const (

	// ConfigStickTableTypeIP captures enum value "ip"
	ConfigStickTableTypeIP string = "ip"

	// ConfigStickTableTypeIPV6 captures enum value "ipv6"
	ConfigStickTableTypeIPV6 string = "ipv6"

	// ConfigStickTableTypeInteger captures enum value "integer"
	ConfigStickTableTypeInteger string = "integer"

	// ConfigStickTableTypeString captures enum value "string"
	ConfigStickTableTypeString string = "string"

	// ConfigStickTableTypeBinary captures enum value "binary"
	ConfigStickTableTypeBinary string = "binary"
)

// prop value enum
func (m *ConfigStickTable) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configStickTableTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConfigStickTable) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ConfigStickTable) validateWriteTo(formats strfmt.Registry) error {
	if swag.IsZero(m.WriteTo) { // not required
		return nil
	}

	if err := validate.Pattern("write_to", "body", *m.WriteTo, `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this config stick table based on context it is used
func (m *ConfigStickTable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigStickTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigStickTable) UnmarshalBinary(b []byte) error {
	var res ConfigStickTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
