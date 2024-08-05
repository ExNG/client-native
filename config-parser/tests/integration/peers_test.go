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

package integration_test

import (
	"bytes"
	"testing"

	parser "github.com/haproxytech/client-native/v6/config-parser"
	"github.com/haproxytech/client-native/v6/config-parser/options"
)

func TestWholeConfigsSectionsPeers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		Name, Config string
	}{
		{"peers_defaultbindtlsticketkeystmptlsti", peers_defaultbindtlsticketkeystmptlsti},
		{"peers_defaultbinduserrootmode600accept", peers_defaultbinduserrootmode600accept},
		{"peers_defaultbindv4v6sslcrtetchaproxys", peers_defaultbindv4v6sslcrtetchaproxys},
		{"peers_defaultserveraddr1", peers_defaultserveraddr1},
		{"peers_defaultserveraddr127001", peers_defaultserveraddr127001},
		{"peers_defaultserveragentaddr127001", peers_defaultserveragentaddr127001},
		{"peers_defaultserveragentaddrsitecom", peers_defaultserveragentaddrsitecom},
		{"peers_defaultserveragentcheck", peers_defaultserveragentcheck},
		{"peers_defaultserveragentinter1000ms", peers_defaultserveragentinter1000ms},
		{"peers_defaultserveragentport1", peers_defaultserveragentport1},
		{"peers_defaultserveragentport65535", peers_defaultserveragentport65535},
		{"peers_defaultserveragentsendname", peers_defaultserveragentsendname},
		{"peers_defaultserverallow0rtt", peers_defaultserverallow0rtt},
		{"peers_defaultserveralpnh2", peers_defaultserveralpnh2},
		{"peers_defaultserveralpnh2http11", peers_defaultserveralpnh2http11},
		{"peers_defaultserveralpnhttp11", peers_defaultserveralpnhttp11},
		{"peers_defaultserverbackup", peers_defaultserverbackup},
		{"peers_defaultservercafilecertcrt", peers_defaultservercafilecertcrt},
		{"peers_defaultservercheck", peers_defaultservercheck},
		{"peers_defaultservercheckalpnhttp10", peers_defaultservercheckalpnhttp10},
		{"peers_defaultservercheckalpnhttp11http", peers_defaultservercheckalpnhttp11http},
		{"peers_defaultservercheckprotoh2", peers_defaultservercheckprotoh2},
		{"peers_defaultserverchecksendproxy", peers_defaultserverchecksendproxy},
		{"peers_defaultservercheckssl", peers_defaultservercheckssl},
		{"peers_defaultservercheckviasocks4", peers_defaultservercheckviasocks4},
		{"peers_defaultserverciphersECDHEECDSACH", peers_defaultserverciphersECDHEECDSACH},
		{"peers_defaultserverciphersECDHERSAAES1", peers_defaultserverciphersECDHERSAAES1},
		{"peers_defaultserverciphersuitesECDHEEC", peers_defaultserverciphersuitesECDHEEC},
		{"peers_defaultservercookievalue", peers_defaultservercookievalue},
		{"peers_defaultservercrlfilefilepem", peers_defaultservercrlfilefilepem},
		{"peers_defaultservercrtcertpem", peers_defaultservercrtcertpem},
		{"peers_defaultserverdisabled", peers_defaultserverdisabled},
		{"peers_defaultserverdowninter3500ms", peers_defaultserverdowninter3500ms},
		{"peers_defaultserverenabled", peers_defaultserverenabled},
		{"peers_defaultservererrorlimit50", peers_defaultservererrorlimit50},
		{"peers_defaultserverfall1rise2inter3spo", peers_defaultserverfall1rise2inter3spo},
		{"peers_defaultserverfall30", peers_defaultserverfall30},
		{"peers_defaultserverfastinter2500ms", peers_defaultserverfastinter2500ms},
		{"peers_defaultserverfastinterunknown", peers_defaultserverfastinterunknown},
		{"peers_defaultserverforcesslv3", peers_defaultserverforcesslv3},
		{"peers_defaultserverforcetlsv10", peers_defaultserverforcetlsv10},
		{"peers_defaultserverforcetlsv11", peers_defaultserverforcetlsv11},
		{"peers_defaultserverforcetlsv12", peers_defaultserverforcetlsv12},
		{"peers_defaultserverforcetlsv13", peers_defaultserverforcetlsv13},
		{"peers_defaultserverinitaddrlastlibcnon", peers_defaultserverinitaddrlastlibcnon},
		{"peers_defaultserverinitaddrlastlibcnon_", peers_defaultserverinitaddrlastlibcnon_},
		{"peers_defaultserverinter1000weight13", peers_defaultserverinter1000weight13},
		{"peers_defaultserverinter1500ms", peers_defaultserverinter1500ms},
		{"peers_defaultserverlogbufsize10", peers_defaultserverlogbufsize10},
		{"peers_defaultserverlogprotolegacy", peers_defaultserverlogprotolegacy},
		{"peers_defaultserverlogprotooctetcount", peers_defaultserverlogprotooctetcount},
		{"peers_defaultservermaxconn1", peers_defaultservermaxconn1},
		{"peers_defaultservermaxconn50", peers_defaultservermaxconn50},
		{"peers_defaultservermaxqueue0", peers_defaultservermaxqueue0},
		{"peers_defaultservermaxqueue1000", peers_defaultservermaxqueue1000},
		{"peers_defaultservermaxreuse0", peers_defaultservermaxreuse0},
		{"peers_defaultservermaxreuse1", peers_defaultservermaxreuse1},
		{"peers_defaultservermaxreuse1_", peers_defaultservermaxreuse1_},
		{"peers_defaultserverminconn1", peers_defaultserverminconn1},
		{"peers_defaultserverminconn50", peers_defaultserverminconn50},
		{"peers_defaultservernamespacetest", peers_defaultservernamespacetest},
		{"peers_defaultservernoagentcheck", peers_defaultservernoagentcheck},
		{"peers_defaultservernobackup", peers_defaultservernobackup},
		{"peers_defaultservernocheck", peers_defaultservernocheck},
		{"peers_defaultservernocheckssl", peers_defaultservernocheckssl},
		{"peers_defaultservernonstick", peers_defaultservernonstick},
		{"peers_defaultservernosendproxyv2", peers_defaultservernosendproxyv2},
		{"peers_defaultservernosendproxyv2ssl", peers_defaultservernosendproxyv2ssl},
		{"peers_defaultservernosendproxyv2sslcn", peers_defaultservernosendproxyv2sslcn},
		{"peers_defaultservernossl", peers_defaultservernossl},
		{"peers_defaultservernosslreuse", peers_defaultservernosslreuse},
		{"peers_defaultservernosslv3", peers_defaultservernosslv3},
		{"peers_defaultservernotfo", peers_defaultservernotfo},
		{"peers_defaultservernotlstickets", peers_defaultservernotlstickets},
		{"peers_defaultservernotlsv10", peers_defaultservernotlsv10},
		{"peers_defaultservernotlsv11", peers_defaultservernotlsv11},
		{"peers_defaultservernotlsv12", peers_defaultservernotlsv12},
		{"peers_defaultservernotlsv13", peers_defaultservernotlsv13},
		{"peers_defaultservernoverifyhost", peers_defaultservernoverifyhost},
		{"peers_defaultservernpnhttp11http10", peers_defaultservernpnhttp11http10},
		{"peers_defaultserverobservelayer4", peers_defaultserverobservelayer4},
		{"peers_defaultserverobservelayer7", peers_defaultserverobservelayer7},
		{"peers_defaultserveronerrorfailcheck", peers_defaultserveronerrorfailcheck},
		{"peers_defaultserveronerrorfastinter", peers_defaultserveronerrorfastinter},
		{"peers_defaultserveronerrormarkdown", peers_defaultserveronerrormarkdown},
		{"peers_defaultserveronerrorsuddendeath", peers_defaultserveronerrorsuddendeath},
		{"peers_defaultserveronmarkeddownshutdow", peers_defaultserveronmarkeddownshutdow},
		{"peers_defaultserveronmarkedupshutdownb", peers_defaultserveronmarkedupshutdownb},
		{"peers_defaultserverpoollowconn384", peers_defaultserverpoollowconn384},
		{"peers_defaultserverpoolmaxconn0", peers_defaultserverpoolmaxconn0},
		{"peers_defaultserverpoolmaxconn1", peers_defaultserverpoolmaxconn1},
		{"peers_defaultserverpoolmaxconn100", peers_defaultserverpoolmaxconn100},
		{"peers_defaultserverpoolpurgedelay0", peers_defaultserverpoolpurgedelay0},
		{"peers_defaultserverpoolpurgedelay5", peers_defaultserverpoolpurgedelay5},
		{"peers_defaultserverpoolpurgedelay500", peers_defaultserverpoolpurgedelay500},
		{"peers_defaultserverport27015", peers_defaultserverport27015},
		{"peers_defaultserverport27016", peers_defaultserverport27016},
		{"peers_defaultserverprotoh2", peers_defaultserverprotoh2},
		{"peers_defaultserverproxyv2optionsssl", peers_defaultserverproxyv2optionsssl},
		{"peers_defaultserverproxyv2optionssslce", peers_defaultserverproxyv2optionssslce},
		{"peers_defaultserverproxyv2optionssslce_", peers_defaultserverproxyv2optionssslce_},
		{"peers_defaultserverredirhttpimage1mydo", peers_defaultserverredirhttpimage1mydo},
		{"peers_defaultserverredirhttpsimage1myd", peers_defaultserverredirhttpsimage1myd},
		{"peers_defaultserverresolvenet100008", peers_defaultserverresolvenet100008},
		{"peers_defaultserverresolvenet100008100", peers_defaultserverresolvenet100008100},
		{"peers_defaultserverresolveoptsallowdup", peers_defaultserverresolveoptsallowdup},
		{"peers_defaultserverresolveoptsallowdup_", peers_defaultserverresolveoptsallowdup_},
		{"peers_defaultserverresolveoptsignorewe", peers_defaultserverresolveoptsignorewe},
		{"peers_defaultserverresolveoptspreventd", peers_defaultserverresolveoptspreventd},
		{"peers_defaultserverresolvepreferipv4", peers_defaultserverresolvepreferipv4},
		{"peers_defaultserverresolvepreferipv6", peers_defaultserverresolvepreferipv6},
		{"peers_defaultserverresolversmydns", peers_defaultserverresolversmydns},
		{"peers_defaultserverrise2", peers_defaultserverrise2},
		{"peers_defaultserverrise200", peers_defaultserverrise200},
		{"peers_defaultserversendproxy", peers_defaultserversendproxy},
		{"peers_defaultserversendproxyv2", peers_defaultserversendproxyv2},
		{"peers_defaultserversendproxyv2ssl", peers_defaultserversendproxyv2ssl},
		{"peers_defaultserversendproxyv2sslcn", peers_defaultserversendproxyv2sslcn},
		{"peers_defaultserverslowstart2000ms", peers_defaultserverslowstart2000ms},
		{"peers_defaultserversniTODO", peers_defaultserversniTODO},
		{"peers_defaultserversocks412700181", peers_defaultserversocks412700181},
		{"peers_defaultserversourceTODO", peers_defaultserversourceTODO},
		{"peers_defaultserverssl", peers_defaultserverssl},
		{"peers_defaultserversslmaxverSSLv3", peers_defaultserversslmaxverSSLv3},
		{"peers_defaultserversslmaxverTLSv10", peers_defaultserversslmaxverTLSv10},
		{"peers_defaultserversslmaxverTLSv11", peers_defaultserversslmaxverTLSv11},
		{"peers_defaultserversslmaxverTLSv12", peers_defaultserversslmaxverTLSv12},
		{"peers_defaultserversslmaxverTLSv13", peers_defaultserversslmaxverTLSv13},
		{"peers_defaultserversslminverSSLv3", peers_defaultserversslminverSSLv3},
		{"peers_defaultserversslminverTLSv10", peers_defaultserversslminverTLSv10},
		{"peers_defaultserversslminverTLSv11", peers_defaultserversslminverTLSv11},
		{"peers_defaultserversslminverTLSv12", peers_defaultserversslminverTLSv12},
		{"peers_defaultserversslminverTLSv13", peers_defaultserversslminverTLSv13},
		{"peers_defaultserversslreuse", peers_defaultserversslreuse},
		{"peers_defaultserverstick", peers_defaultserverstick},
		{"peers_defaultservertcput20ms", peers_defaultservertcput20ms},
		{"peers_defaultservertfo", peers_defaultservertfo},
		{"peers_defaultservertlstickets", peers_defaultservertlstickets},
		{"peers_defaultservertrackTODO", peers_defaultservertrackTODO},
		{"peers_defaultserververifyhostsitecom", peers_defaultserververifyhostsitecom},
		{"peers_defaultserververifynone", peers_defaultserververifynone},
		{"peers_defaultserververifyrequired", peers_defaultserververifyrequired},
		{"peers_defaultserverweight1", peers_defaultserverweight1},
		{"peers_defaultserverweight128", peers_defaultserverweight128},
		{"peers_defaultserverweight256", peers_defaultserverweight256},
		{"peers_defaultserverwsauto", peers_defaultserverwsauto},
		{"peers_defaultserverwsh1", peers_defaultserverwsh1},
		{"peers_defaultserverwsh2", peers_defaultserverwsh2},
		{"peers_peername1270018080", peers_peername1270018080},
		{"peers_peername1270018080shard1", peers_peername1270018080shard1},
	}
	for _, config := range tests {
		t.Run(config.Name, func(t *testing.T) {
			t.Parallel()
			var buffer bytes.Buffer
			buffer.WriteString(config.Config)
			p, err := parser.New(options.Reader(&buffer))
			if err != nil {
				t.Fatalf(err.Error())
			}
			result := p.String()
			if result != config.Config {
				compare(t, config.Config, result)
				t.Error("======== ORIGINAL =========")
				t.Error(config.Config)
				t.Error("======== RESULT ===========")
				t.Error(result)
				t.Error("===========================")
				t.Fatalf("configurations does not match")
			}
		})
	}
}
