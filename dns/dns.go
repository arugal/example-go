// Copyright 2020 arugal, zhangwei24@apache.com
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

package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

func main() {
	// 读取当前运行环境的 /etc/resolv.conf，获得 name server 的配置
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

	// 构造发起 DNS 请求的客户端
	c := new(dns.Client)

	// 构造 DNS 报文
	m := new(dns.Msg)

	m.SetQuestion(dns.Fqdn("www.baidu.com"), dns.TypeA)
	m.RecursionDesired = true

	// client 发起 DNS 请求，其中 c 为上文创建的 client，m 为构造的 DNS 报文
	// config 为从 /etc/resolv.conf 构造出来的配置
	r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
	if r == nil {
		log.Fatalf("*** error: %s\n", err.Error())
	}

	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("*** invalid answer name %s after MX query for %s\n", os.Args[1], os.Args[1])
	}

	// 如果 DNS 查询成功
	for _, a := range r.Answer {
		fmt.Printf("%v\n", a)
	}
}
