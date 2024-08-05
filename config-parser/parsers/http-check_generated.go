// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package parsers

import (
	"github.com/haproxytech/client-native/v6/config-parser/common"
	"github.com/haproxytech/client-native/v6/config-parser/errors"
	"github.com/haproxytech/client-native/v6/config-parser/types"
)

func (p *HTTPCheckV2) Init() {
	p.data = []types.HTTPCheckV2{}
	p.preComments = []string{}
}

func (p *HTTPCheckV2) GetParserName() string {
	return "http-check"
}

func (p *HTTPCheckV2) Get(createIfNotExist bool) (common.ParserData, error) {
	if len(p.data) == 0 && !createIfNotExist {
		return nil, errors.ErrFetch
	}
	return p.data, nil
}

func (p *HTTPCheckV2) GetPreComments() ([]string, error) {
	return p.preComments, nil
}

func (p *HTTPCheckV2) SetPreComments(preComments []string) {
	p.preComments = preComments
}

func (p *HTTPCheckV2) GetOne(index int) (common.ParserData, error) {
	if index < 0 || index >= len(p.data) {
		return nil, errors.ErrFetch
	}
	return p.data[index], nil
}

func (p *HTTPCheckV2) Delete(index int) error {
	if index < 0 || index >= len(p.data) {
		return errors.ErrFetch
	}
	copy(p.data[index:], p.data[index+1:])
	p.data[len(p.data)-1] = types.HTTPCheckV2{}
	p.data = p.data[:len(p.data)-1]
	return nil
}

func (p *HTTPCheckV2) Insert(data common.ParserData, index int) error {
	if data == nil {
		return errors.ErrInvalidData
	}
	switch newValue := data.(type) {
	case []types.HTTPCheckV2:
		p.data = newValue
	case *types.HTTPCheckV2:
		if index > -1 {
			if index > len(p.data) {
				return errors.ErrIndexOutOfRange
			}
			p.data = append(p.data, types.HTTPCheckV2{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = *newValue
		} else {
			p.data = append(p.data, *newValue)
		}
	case types.HTTPCheckV2:
		if index > -1 {
			if index > len(p.data) {
				return errors.ErrIndexOutOfRange
			}
			p.data = append(p.data, types.HTTPCheckV2{})
			copy(p.data[index+1:], p.data[index:])
			p.data[index] = newValue
		} else {
			p.data = append(p.data, newValue)
		}
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *HTTPCheckV2) Set(data common.ParserData, index int) error {
	if data == nil {
		p.Init()
		return nil
	}
	switch newValue := data.(type) {
	case []types.HTTPCheckV2:
		p.data = newValue
	case *types.HTTPCheckV2:
		if index > -1 && index < len(p.data) {
			p.data[index] = *newValue
		} else if index == -1 {
			p.data = append(p.data, *newValue)
		} else {
			return errors.ErrIndexOutOfRange
		}
	case types.HTTPCheckV2:
		if index > -1 && index < len(p.data) {
			p.data[index] = newValue
		} else if index == -1 {
			p.data = append(p.data, newValue)
		} else {
			return errors.ErrIndexOutOfRange
		}
	default:
		return errors.ErrInvalidData
	}
	return nil
}

func (p *HTTPCheckV2) PreParse(line string, parts []string, preComments []string, comment string) (string, error) {
	changeState, err := p.Parse(line, parts, comment)
	if err == nil && preComments != nil {
		p.preComments = append(p.preComments, preComments...)
	}
	return changeState, err
}

func (p *HTTPCheckV2) Parse(line string, parts []string, comment string) (string, error) {
	if parts[0] == "http-check" {
		data, err := p.parse(line, parts, comment)
		if err != nil {
			if _, ok := err.(*errors.ParseError); ok {
				return "", err
			}
			return "", &errors.ParseError{Parser: "HTTPCheckV2", Line: line}
		}
		p.data = append(p.data, *data)
		return "", nil
	}
	return "", &errors.ParseError{Parser: "HTTPCheckV2", Line: line}
}

func (p *HTTPCheckV2) ResultAll() ([]common.ReturnResultLine, []string, error) {
	res, err := p.Result()
	return res, p.preComments, err
}
