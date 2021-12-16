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

package resolver

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type nacosResolver struct {
	cli naming_client.INamingClient
}

func NewNacosResolver(cli naming_client.INamingClient) discovery.Resolver {
	return &nacosResolver{cli: cli}
}

func (n nacosResolver) Target(_ context.Context, target rpcinfo.EndpointInfo) (description string) {
	return target.ServiceName()
}

func (n nacosResolver) Resolve(_ context.Context, desc string) (discovery.Result, error) {
	res, err := n.cli.SelectInstances(vo.SelectInstancesParam{
		ServiceName: desc,
		HealthyOnly: true,
	})
	if err != nil {
		return discovery.Result{}, err
	}
	var instances []discovery.Instance
	for _, in := range res {
		instances = append(instances, discovery.NewInstance(
			"tcp",
			fmt.Sprintf("%s:%d", in.Ip, in.Port),
			int(in.Weight),
			in.Metadata),
		)
	}
	return discovery.Result{
		Cacheable: true,
		CacheKey:  desc,
		Instances: instances,
	}, nil
}

func (n nacosResolver) Diff(cacheKey string, prev, next discovery.Result) (discovery.Change, bool) {
	return discovery.DefaultDiff(cacheKey, prev, next)
}

func (n nacosResolver) Name() string {
	return "nacos"
}

var _ discovery.Resolver = (*nacosResolver)(nil)
