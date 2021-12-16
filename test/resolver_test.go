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

package test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/registry"
	nacosregistry "github.com/kitex-contrib/registry-nacos/registry"
	nacosresolver "github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/stretchr/testify/assert"
)

func Test_nacosResolver_Resolve(t *testing.T) {
	client, err := getNacosClient()
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	info := &registry.Info{
		ServiceName: "demo.kitex-contrib.local",
		Addr:        &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8848},
		Weight:      999,
		StartTime:   time.Now(),
		Tags:        map[string]string{"env": "local"},
	}

	err = nacosregistry.NewNacosRegistry(client).Register(info)
	if err != nil {
		t.Errorf("Register Fail:%v", err)
		return
	}

	type fields struct {
		cli naming_client.INamingClient
	}
	type args struct {
		ctx  context.Context
		desc string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    discovery.Result
		wantErr bool
	}{
		{
			name: "common",
			args: args{
				ctx:  context.Background(),
				desc: "demo.kitex-contrib.local",
			},
			fields: fields{cli: client},
		},
		{
			name: "wrong desc",
			args: args{
				ctx:  context.Background(),
				desc: "xxxx.kitex-contrib.local",
			},
			fields:  fields{cli: client},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := nacosresolver.NewNacosResolver(tt.fields.cli)
			got, err := n.Resolve(tt.args.ctx, tt.args.desc)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotEmpty(t, got.Instances)
				if len(got.Instances) > 0 {
					assert.Equal(t, got.Instances[0].Address().String(), "127.0.0.1:8848")
					assert.Equal(t, got.Instances[0].Weight(), 999)
				}
			}
		})
	}

	err = nacosregistry.NewNacosRegistry(client).Deregister(info)
	if err != nil {
		t.Errorf("Deregister Fail")
		return
	}
}
