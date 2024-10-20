// Copyright 2024 CloudWeGo Authors.
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

package nacos

import (
	"os"
	"strings"

	"github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoskitex/v2/nacos"
)

const (
	NACOS_ENV_SERVER_ADDR     = "serverAddr"
	NACOS_ENV_PORT            = "serverPort"
	NACOS_ENV_TAGS            = "KITEX_NACOS_ENV_TAGS"
	NACOS_ENV_NAMESPACE_ID    = "namespace"
	NACOS_DEFAULT_SERVER_ADDR = "127.0.0.1"
	NACOS_DEFAULT_PORT        = 8848
	NACOS_DEFAULT_REGIONID    = "cn-hangzhou"
)

// Tags providers the default tags to inject nacos.
var Tags map[string]string

func init() {
	Tags = parseTags(os.Getenv(NACOS_ENV_TAGS))
}

func parseTags(tags string) map[string]string {
	if len(tags) == 0 {
		return nil
	}
	out := map[string]string{}
	parts := strings.Split(tags, ",")
	for _, part := range parts {
		tag := strings.Split(part, "=")
		if len(tag) == 2 {
			out[tag[0]] = tag[1]
		}
	}
	return out
}

// NacosPort Get Nacos port from environment variables
func NacosPort() int64 {
	return nacos.NacosPort()
}

// NacosAddr Get Nacos addr from environment variables
func NacosAddr() string {
	return nacos.NacosAddr()
}

// NacosNameSpaceId Get Nacos namespace id from environment variables
func NacosNameSpaceId() string {
	return nacos.NacosNameSpaceId()
}
