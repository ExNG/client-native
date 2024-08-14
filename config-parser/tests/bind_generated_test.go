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

func TestBind(t *testing.T) {
	tests := map[string]bool{
		"bind :80,:443":                      true,
		"bind 10.0.0.1:10080,10.0.0.1:10443": true,
		"bind /var/run/ssl-frontend.sock user root mode 600 accept-proxy": true,
		"bind :80": true,
		"bind :443 ssl crt /etc/haproxy/site.pem":                         true,
		"bind :443 ssl crt /etc/haproxy/site.pem alpn h2,http/1.1":        true,
		"bind :::443 v4v6 ssl crt /etc/haproxy/site.pem alpn h2,http/1.1": true,
		"bind ipv6@:80": true,
		"bind ipv4@public_ssl:443 ssl crt /etc/haproxy/site.pem":               true,
		"bind unix@ssl-frontend.sock user root mode 600 accept-proxy":          true,
		"bind :443 accept-netscaler-cip 1234":                                  true,
		"bind :443 accept-proxy":                                               true,
		"bind :443 allow-0rtt":                                                 true,
		"bind :443 alpn h2":                                                    true,
		"bind :443 alpn http/1.1":                                              true,
		"bind :443 alpn h2,http/1.1":                                           true,
		"bind :443 backlog test":                                               true,
		"bind :443 curves ECDH_ECDSA,ECDHE_ECDSA,ECDH_RSA,ECDHE_RSA,ECDH_anon": true,
		"bind :443 ecdhe ECDH_ECDSA,ECDHE_ECDSA,ECDH_RSA,ECDHE_RSA,ECDH_anon":  true,
		"bind :443 ca-file file.pem":                                           true,
		"bind :443 ca-ignore-err all":                                          true,
		"bind :443 ca-ignore-err 1234":                                         true,
		"bind :443 ca-sign-file file.test":                                     true,
		"bind :443 ca-sign-pass passphrase":                                    true,
		"bind :443 ca-verify-file file.test":                                   true,
		"bind :443 ciphers ECDHE+aRSA+AES256+GCM+SHA384:ECDHE+aRSA+AES128+GCM+SHA256:ECDHE+aRSA+AES256+SHA384:ECDHE+aRSA+AES128+SHA256:ECDHE+aRSA+RC4+SHA:ECDHE+aRSA+AES256+SHA:ECDHE+aRSA+AES128+SHA:AES256+GCM+SHA384:AES128+GCM+SHA256:AES128+SHA256:AES256+SHA256:DHE+aRSA+AES128+SHA:RC4+SHA:HIGH:!aNULL:!eNULL:!LOW:!3DES:!MD5:!EXP:!PSK:!SRP:!DSS": true,
		"bind :443 ciphersuites TODO":                    true,
		"bind :443 client-sigalgs value":                 true,
		"bind :443 crl-file file.test":                   true,
		"bind :443 crt example.pem":                      true,
		"bind :443 crt-ignore-err all":                   true,
		"bind :443 crt-ignore-err 404,410":               true,
		"bind :443 crt-list cert1.pem":                   true,
		"bind :443 defer-accept":                         true,
		"bind :443 expose-fd listeners":                  true,
		"bind :443 force-sslv3":                          true,
		"bind :443 force-tlsv10":                         true,
		"bind :443 force-tlsv11":                         true,
		"bind :443 force-tlsv12":                         true,
		"bind :443 force-tlsv13":                         true,
		"bind :443 generate-certificates":                true,
		"bind :443 gid users":                            true,
		"bind :443 group group":                          true,
		"bind :443 id 1":                                 true,
		"bind :443 interface eth0":                       true,
		"bind :443 interface eth1":                       true,
		"bind :443 interface pppoe-wan":                  true,
		"bind :443 level user":                           true,
		"bind :443 level opeerator":                      true,
		"bind :443 level admin":                          true,
		"bind :443 severity-output none":                 true,
		"bind :443 severity-output number":               true,
		"bind :443 severity-output string":               true,
		"bind :443 maxconn 1024":                         true,
		"bind :443 mode TODO":                            true,
		"bind :443 mss 1460":                             true,
		"bind :443 mss -1460":                            true,
		"bind :443 name sockets":                         true,
		"bind :443 namespace example":                    true,
		"bind :443 nice 0":                               true,
		"bind :443 nice 1024":                            true,
		"bind :443 nice -1024":                           true,
		"bind :443 no-alpn":                              true,
		"bind :443 no-ca-names":                          true,
		"bind :443 no-sslv3":                             true,
		"bind :443 no-tls-tickets":                       true,
		"bind :443 no-tlsv10":                            true,
		"bind :443 no-tlsv11":                            true,
		"bind :443 no-tlsv12":                            true,
		"bind :443 no-tlsv13":                            true,
		"bind :443 npn http/1.0":                         true,
		"bind :443 npn http/1.1":                         true,
		"bind :443 npn http/1.0,http/1.1":                true,
		"bind :443 ocsp-update off":                      true,
		"bind :443 ocsp-update on":                       true,
		"bind :443 prefer-client-ciphers":                true,
		"bind :443 process all":                          true,
		"bind :443 process odd":                          true,
		"bind :443 process even":                         true,
		"bind :443 process 1-4":                          true,
		"bind :443 proto h2":                             true,
		"bind :443 sigalgs value":                        true,
		"bind :443 ssl":                                  true,
		"bind :443 ssl-max-ver SSLv3":                    true,
		"bind :443 ssl-max-ver TLSv1.0":                  true,
		"bind :443 ssl-max-ver TLSv1.1":                  true,
		"bind :443 ssl-max-ver TLSv1.2":                  true,
		"bind :443 ssl-max-ver TLSv1.3":                  true,
		"bind :443 ssl-min-ver SSLv3":                    true,
		"bind :443 ssl-min-ver TLSv1.0":                  true,
		"bind :443 ssl-min-ver TLSv1.1":                  true,
		"bind :443 ssl-min-ver TLSv1.2":                  true,
		"bind :443 ssl-min-ver TLSv1.3":                  true,
		"bind :443 strict-sni":                           true,
		"bind :443 tcp-ut 30s":                           true,
		"bind :443 tfo":                                  true,
		"bind :443 thread all":                           true,
		"bind :443 thread odd":                           true,
		"bind :443 thread even":                          true,
		"bind :443 thread 1":                             true,
		"bind :443 thread 1-1":                           true,
		"bind :443 thread 1/all":                         true,
		"bind :443 thread 1/odd":                         true,
		"bind :443 thread 1/even":                        true,
		"bind :443 thread 1/1":                           true,
		"bind :443 thread 1/1-1":                         true,
		"bind :443 tls-ticket-keys /tmp/tls_ticket_keys": true,
		"bind :443 transparent":                          true,
		"bind :443 v4v6":                                 true,
		"bind :443 v6only":                               true,
		"bind :443 uid 65534":                            true,
		"bind :443 user web1":                            true,
		"bind :443 verify none":                          true,
		"bind :443 verify optional":                      true,
		"bind :443 verify required":                      true,
		"bind :443 quic-cc-algo cubic":                   true,
		"bind :443 quic-cc-algo newreno":                 true,
		"bind :443 quic-force-retry":                     true,
		"bind :443 quic-socket connection":               true,
		"bind :443 quic-socket listener":                 true,
		"bind :443 nbconn 1":                             true,
		"bind :443 nbconn +2":                            true,
		"bind :443 guid-prefix guid-example":             true,
		"bind :443 default-crt foobar.pem.rsa default-crt foobar.pem.ecdsa": true,
		"bind":                             false,
		"bind :443 quic-cc-algo something": false,
		"bind :443 quic-socket something":  false,
		"---":                              false,
		"--- ---":                          false,
	}
	parser := &parsers.Bind{}
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