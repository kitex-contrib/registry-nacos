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

package resolver

import (
	nacosresolver "github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/nacoskitex/v2/resolver"
	nacosOption "github.com/cloudwego-contrib/cwgo-pkg/registry/nacos/options"
	"github.com/cloudwego/kitex/pkg/discovery"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
)

// Option is nacos option.
type Option = nacosOption.ResolverOption

// WithCluster with cluster option.
func WithCluster(cluster string) Option {
	return nacosOption.WithResolverCluster(cluster)
}

// WithGroup with group option.
func WithGroup(group string) Option {
	return nacosOption.WithResolverGroup(group)
}

// NewDefaultNacosResolver create a default service resolver using nacos.
func NewDefaultNacosResolver(opts ...Option) (discovery.Resolver, error) {
	return nacosresolver.NewDefaultNacosResolver(opts...)
}

// NewNacosResolver create a service resolver using nacos.
func NewNacosResolver(cli naming_client.INamingClient, opts ...Option) discovery.Resolver {
	return nacosresolver.NewNacosResolver(cli, opts...)
}
