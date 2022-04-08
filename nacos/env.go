// Copyright 2021 CloudWeGo Authors.
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
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	NACOS_ENV_SERVICE_ADDR     = "serviceAddr"
	NACOS_ENV_PORT             = "servicePort"
	NACOS_ENV_NAMESPACE_ID     = "namespace"
	NACOS_DEFAULT_SERVICE_ADDR = "127.0.0.1"
	NACOS_DEFAULT_PORT         = 8848
	NACOS_DEFAULT_REGIONID     = "cn-hangzhou"
)

// NacosPort Get Nacos port from environment variables
func NacosPort() int64 {
	portText := os.Getenv(NACOS_ENV_PORT)
	if len(portText) == 0 {
		return NACOS_DEFAULT_PORT
	}
	port, err := strconv.ParseInt(portText, 10, 64)
	if err != nil {
		klog.Errorf("ParseInt failed,err:%+v", err)
		return NACOS_DEFAULT_PORT
	}
	return port
}

// NacosAddr Get Nacos addr from environment variables
func NacosAddr() string {
	addr := os.Getenv(NACOS_ENV_SERVICE_ADDR)
	if len(addr) == 0 {
		return NACOS_DEFAULT_SERVICE_ADDR
	}
	return addr
}

// NacosNameSpaceId Get Nacos namespace id from environment variables
func NacosNameSpaceId() string {
	return os.Getenv(NACOS_ENV_NAMESPACE_ID)
}
