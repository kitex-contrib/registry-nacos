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
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// Option is the way to config a client.
type Option struct {
	F func(o *vo.NacosClientParam)
}

// NewDefaultNacosClient Create a default Nacos client
// It can create a client with default config by env variable.
// See: env.go
func NewDefaultNacosClient(opts ...Option) (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(NacosAddr(), uint64(NacosPort())),
	}
	cc := constant.ClientConfig{
		NamespaceId:         NacosNameSpaceId(),
		RegionId:            NACOS_DEFAULT_REGIONID,
		NotLoadCacheAtStart: true,
		CustomLogger:        NewCustomNacosLogger(),
	}
	param := vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	}
	for _, opt := range opts {
		opt.F(&param)
	}
	cli, err := clients.NewNamingClient(param)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
