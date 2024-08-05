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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/client-native/v6/config-parser/parsers"
)

func TestLogStdErr(t *testing.T) {
	tests := map[string]bool{
		"log-stderr 127.0.0.1:1515 len 8192 format rfc5424 sample 1,2-5:6 local2 info debug": true,
		"log-stderr 127.0.0.1:1515 len 8192 format rfc5424 sample 1,2-5:6 local2 info":       true,
		"log-stderr 127.0.0.1:1515 local2":                                                   true,
		"log-stderr global":                                                                  true,
		"log-stderr":                                                                         false,
		"log-stderr 127.0.0.1:1515":                                                          false,
		"log-stderr 127.0.0.1:1515 len 8192 format rfc5424 sample 1,2-5 local2 info debug":   false,
		"log-stderr 127.0.0.1:1515 len 8192 format sample 1,2-5:6 local2 info debug":         false,
		"log-stderr 127.0.0.1:1515 len format rfc5424 sample 1,2-5:6 local2 info debug":      false,
		"---":     false,
		"--- ---": false,
	}
	parser := &parsers.LogStdErr{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
			line := strings.TrimSpace(command)
			lines := strings.SplitN(line, "\n", -1)
			var err error
			parser.Init()
			if len(lines) > 1 {
				for _, line = range lines {
					line = strings.TrimSpace(line)
					if err = ProcessLine(line, parser); err != nil {
						break
					}
				}
			} else {
				err = ProcessLine(line, parser)
			}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
