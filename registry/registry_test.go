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
	"net"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/stretchr/testify/assert"
)

func getNacosClient() (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}

	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
}

func TestNewNacosRegistry(t *testing.T) {
	client, err := getNacosClient()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	got := NewNacosRegistry(client, WithCluster("DEFAULT"), WithCluster("DEFAULT_GROUP"))
	assert.NotNil(t, got)
}

func Test_nacosRegistry_Register(t *testing.T) {
	client, err := getNacosClient()
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
			fields: fields{client},
			args: args{info: &registry.Info{
				ServiceName: "demo.kitex-contrib.local",
				Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8848},
				Weight:      999,
				StartTime:   time.Now(),
				Tags:        map[string]string{"env": "local"},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNacosRegistry(tt.fields.cli, WithCluster("DEFAULT"), WithCluster("DEFAULT_GROUP"))
			if err := n.Register(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nacosRegistry_Deregister(t *testing.T) {
	client, err := getNacosClient()
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
				Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8848},
				Weight:      999,
				StartTime:   time.Now(),
				Tags:        map[string]string{"env": "local"},
			}},
			fields:  fields{client},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewNacosRegistry(tt.fields.cli, WithCluster("DEFAULT"), WithCluster("DEFAULT_GROUP"))
			if err := n.Deregister(tt.args.info); (err != nil) != tt.wantErr {
				t.Errorf("Deregister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
