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
	"github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoskitex/v2/nacos"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
)

// Option is the way to config a client.
type Option = nacos.Option

// NewDefaultNacosClient Create a default Nacos client
// It can create a client with default config by env variable.
// See: env.go
func NewDefaultNacosClient(opts ...Option) (naming_client.INamingClient, error) {
	return nacos.NewDefaultNacosClient(opts...)
}
