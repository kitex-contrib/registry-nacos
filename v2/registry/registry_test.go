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

package registry

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/v2/example/hello/kitex_gen/api"
	"github.com/kitex-contrib/registry-nacos/v2/example/hello/kitex_gen/api/hello"
	"github.com/kitex-contrib/registry-nacos/v2/resolver"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/stretchr/testify/assert"
)

func getNacosClient() (naming_client.INamingClient, error) {
	// create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos"), constant.WithScheme("http")),
	}

	// create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithTimeoutMs(50000),
		constant.WithUpdateCacheWhenEmpty(true),
		constant.WithNotLoadCacheAtStart(true),
	)

	// create naming client
	newClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return newClient, err
}

// TestNewNacosRegistry test registry a service
func TestNacosRegistryRegister(t *testing.T) {
	nacosClient, err := getNacosClient()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	type fields struct {
		cli naming_client.INamingClient
	}
	type args struct {
		info *registry.Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "common",
			fields: fields{nacosClient},
			args: args{info: &registry.Info{
				ServiceName: "demo.kitex-contrib.local",
				Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080},
				Weight:      999,
				StartTime:   time.Now(),
				Tags:        map[string]string{"env": "local"},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNacosRegistry(tt.fields.cli, WithCluster("DEFAULT"), WithGroup("DEFAULT_GROUP"))
			if err := n.Register(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestNacosRegistryDeregister test deregister a service
func TestNacosRegistryDeregister(t *testing.T) {
	nacosClient, err := getNacosClient()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	type fields struct {
		cli naming_client.INamingClient
	}
	type args struct {
		info *registry.Info
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "common",
			args: args{info: &registry.Info{
				ServiceName: "demo.kitex-contrib.local",
				Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080},
				Weight:      999,
				StartTime:   time.Now(),
				Tags:        map[string]string{"env": "local"},
			}},
			fields:  fields{nacosClient},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNacosRegistry(tt.fields.cli, WithCluster("DEFAULT"), WithGroup("DEFAULT_GROUP"))
			if err := n.Deregister(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Deregister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestNacosMultipleInstancesWithDefaultNacosRegistry use DefaultNacosRegistry to test registry multiple service,then deregister one
func TestNacosMultipleInstancesWithDefaultNacosRegistry(t *testing.T) {
	var (
		svcName     = "MultipleInstances"
		clusterName = "TheCluster"
		groupName   = "TheGroup"
	)
	got, err := NewDefaultNacosRegistry(WithCluster(clusterName), WithGroup(groupName))
	assert.Nil(t, err)
	if !assert.NotNil(t, got) {
		t.Errorf("err: new registry fail")
		return
	}
	time.Sleep(time.Second)
	err = got.Register(&registry.Info{
		ServiceName: svcName,
		Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081},
	})
	assert.Nil(t, err)

	time.Sleep(time.Second * 1)
	nacosClient, err := getNacosClient()
	assert.Nil(t, err)
	res, err := nacosClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: svcName,
		GroupName:   groupName,
		Clusters:    []string{clusterName},
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, len(res))

	time.Sleep(time.Second)
	err = got.Deregister(&registry.Info{
		ServiceName: svcName,
		Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8081},
	})
	assert.Nil(t, err)

	time.Sleep(time.Second * 3)
	res, err = nacosClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: svcName,
		GroupName:   groupName,
		Clusters:    []string{clusterName},
		HealthyOnly: true,
	})
	assert.Equal(t, "instance list is empty!", err.Error())
	assert.Equal(t, 0, len(res))
}

func TestMergeTags(t *testing.T) {
	t1 := map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	t2 := map[string]string{
		"k3": "v3",
		"k4": "v4",
	}
	merged := mergeTags(t1, t2)
	assert.Equal(t, merged, map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	})
	assert.Equal(t, t1, map[string]string{
		"k1": "v1",
		"k2": "v2",
	})
	assert.Equal(t, t2, map[string]string{
		"k3": "v3",
		"k4": "v4",
	})
}

type HelloImpl struct{}

func (h *HelloImpl) Echo(_ context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{
		Message: req.Message,
	}
	return
}

func TestResolverDifferentGroup(t *testing.T) {
	r, err := NewDefaultNacosRegistry()
	assert.Nil(t, err)
	r2, err := NewDefaultNacosRegistry(WithGroup("OTHER"))
	assert.Nil(t, err)

	svr := hello.NewServer(
		new(HelloImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "demo1",
			Addr:        utils.NewNetAddr("tcp", "127.0.0.1:8080"),
			Weight:      10,
			Tags:        nil,
		}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "demo1"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}),
	)

	svr2 := hello.NewServer(
		new(HelloImpl),
		server.WithRegistry(r2),
		server.WithRegistryInfo(&registry.Info{
			ServiceName: "demo2",
			Addr:        utils.NewNetAddr("tcp", "127.0.0.1:8082"),
			Weight:      10,
			Tags:        nil,
		}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "demo2"}),
		server.WithServiceAddr(&net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8082}),
	)

	go svr.Run()  //nolint:errcheck
	go svr2.Run() //nolint:errcheck
	time.Sleep(2 * time.Second)

	resolver1, err := resolver.NewDefaultNacosResolver()
	assert.Nil(t, err)
	resolver2, err := resolver.NewDefaultNacosResolver(resolver.WithGroup("OTHER"))
	assert.Nil(t, err)

	client1 := hello.MustNewClient(
		"demo1",
		client.WithResolver(resolver1),
		client.WithRPCTimeout(time.Second*3),
	)
	client2 := hello.MustNewClient(
		"demo2",
		client.WithResolver(resolver2),
		client.WithRPCTimeout(time.Second*3),
	)
	resp, err := client1.Echo(context.Background(), &api.Request{Message: "Hello"})
	assert.Nil(t, err)
	assert.Equal(t, resp.Message, "Hello")

	resp, err = client2.Echo(context.Background(), &api.Request{Message: "Hello1"})
	assert.Nil(t, err)
	assert.Equal(t, resp.Message, "Hello1")

	defer func() {
		svr.Stop()  //nolint:errcheck
		svr2.Stop() //nolint:errcheck
	}()
}
