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

package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/v2/example/hello/kitex_gen/api"
	"github.com/kitex-contrib/registry-nacos/v2/example/hello/kitex_gen/api/hello"
	"github.com/kitex-contrib/registry-nacos/v2/resolver"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// the nacos server config
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848,
			func(sc *constant.ServerConfig) {
				// it is not recommended to modify GRPC-Port unless necessary,default by server port +1000
				sc.GrpcPort = 9848
			}),
	}

	// the nacos client config
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		// more ...
	}

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	newClient := hello.MustNewClient(
		"Hello",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	for {
		resp, err := newClient.Echo(context.Background(), &api.Request{Message: "Hello"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
