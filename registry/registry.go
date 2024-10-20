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

package registry

import (
	nacosRegistry "github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoskitex/registry"
	nacosOption "github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/options"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
)

// Option is nacos option.
type Option = nacosOption.Option

// WithCluster with cluster option.
func WithCluster(cluster string) Option {
	return nacosOption.WithCluster(cluster)
}

// WithGroup with group option.
func WithGroup(group string) Option {
	return nacosOption.WithGroup(group)
}

// NewDefaultNacosRegistry create a default service registry using nacos.
func NewDefaultNacosRegistry(opts ...Option) (registry.Registry, error) {
	return nacosRegistry.NewDefaultNacosRegistry(opts...)
}

// NewNacosRegistry create a new registry using nacos.
func NewNacosRegistry(cli naming_client.INamingClient, opts ...Option) registry.Registry {
	return nacosRegistry.NewNacosRegistry(cli, opts...)
}

// should not modify the source data.
func mergeTags(ts ...map[string]string) map[string]string {
	if len(ts) == 0 {
		return nil
	}
	if len(ts) == 1 {
		return ts[0]
	}
	tags := map[string]string{}
	for _, t := range ts {
		for k, v := range t {
			tags[k] = v
		}
	}
	return tags
}
