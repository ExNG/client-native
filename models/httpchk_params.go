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

// HttpchkParams httpchk params
//
// swagger:model httpchk_params
type HttpchkParams struct {
	// method
	// Enum: ["HEAD","PUT","POST","GET","TRACE","PATCH","DELETE","CONNECT","OPTIONS"]
	// +kubebuilder:validation:Enum="HEAD","PUT","POST","GET","TRACE","PATCH","DELETE","CONNECT","OPTIONS";
	Method string `json:"method,omitempty"`

	// uri
	// Pattern: ^[^ ]*$
	// +kubebuilder:validation:Pattern=`^[^ ]*$`
	URI string `json:"uri,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this httpchk params
func (m *HttpchkParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var httpchkParamsTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEAD","PUT","POST","GET","TRACE","PATCH","DELETE","CONNECT","OPTIONS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpchkParamsTypeMethodPropEnum = append(httpchkParamsTypeMethodPropEnum, v)
	}
}

const (

	// HttpchkParamsMethodHEAD captures enum value "HEAD"
	HttpchkParamsMethodHEAD string = "HEAD"

	// HttpchkParamsMethodPUT captures enum value "PUT"
	HttpchkParamsMethodPUT string = "PUT"

	// HttpchkParamsMethodPOST captures enum value "POST"
	HttpchkParamsMethodPOST string = "POST"

	// HttpchkParamsMethodGET captures enum value "GET"
	HttpchkParamsMethodGET string = "GET"

	// HttpchkParamsMethodTRACE captures enum value "TRACE"
	HttpchkParamsMethodTRACE string = "TRACE"

	// HttpchkParamsMethodPATCH captures enum value "PATCH"
	HttpchkParamsMethodPATCH string = "PATCH"

	// HttpchkParamsMethodDELETE captures enum value "DELETE"
	HttpchkParamsMethodDELETE string = "DELETE"

	// HttpchkParamsMethodCONNECT captures enum value "CONNECT"
	HttpchkParamsMethodCONNECT string = "CONNECT"

	// HttpchkParamsMethodOPTIONS captures enum value "OPTIONS"
	HttpchkParamsMethodOPTIONS string = "OPTIONS"
)

// prop value enum
func (m *HttpchkParams) validateMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, httpchkParamsTypeMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HttpchkParams) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", m.Method); err != nil {
		return err
	}

	return nil
}

func (m *HttpchkParams) validateURI(formats strfmt.Registry) error {
	if swag.IsZero(m.URI) { // not required
		return nil
	}

	if err := validate.Pattern("uri", "body", m.URI, `^[^ ]*$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this httpchk params based on context it is used
func (m *HttpchkParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HttpchkParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HttpchkParams) UnmarshalBinary(b []byte) error {
	var res HttpchkParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
