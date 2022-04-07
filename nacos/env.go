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
	NacosEnvServiceAddr     = "serviceAddr"
	NacosEnvPort            = "servicePort"
	NacosEnvNameSpaceId     = "namespace"
	NacosDefaultServiceAddr = "127.0.0.1"
	NacosDefaultPort        = 8848
	NacosDefaultRegionId    = "cn-hangzhou"
)

// NacosPort Get Nacos port from environment variables
func NacosPort() int64 {
	portText := os.Getenv(NacosEnvPort)
	if len(portText) == 0 {
		return NacosDefaultPort
	}
	port, err := strconv.ParseInt(portText, 10, 64)
	if err != nil {
		klog.Errorf("ParseInt failed,err=%+v", err)
		return NacosDefaultPort
	}
	return port
}

// NacosAddr Get Nacos addr from environment variables
func NacosAddr() string {
	addr := os.Getenv(NacosEnvServiceAddr)
	if len(addr) == 0 {
		return NacosDefaultServiceAddr
	}
	return addr
}

// NacosNameSpaceId Get Nacos namespace id from environment variables
func NacosNameSpaceId() string {
	return os.Getenv(NacosEnvNameSpaceId)
}
