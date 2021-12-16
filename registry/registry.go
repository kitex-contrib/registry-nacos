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
	"fmt"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type nacosRegistry struct {
	cli naming_client.INamingClient
}

func NewNacosRegistry(cli naming_client.INamingClient) registry.Registry {
	return &nacosRegistry{cli: cli}
}

var _ registry.Registry = (*nacosRegistry)(nil)

func (n *nacosRegistry) Register(info *registry.Info) error {
	if info.ServiceName == "" {
		return fmt.Errorf("registry.Info cannot is empty")
	}
	host, port, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		return fmt.Errorf("parse registry info addr error:%v", err)
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	_, e := n.cli.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(p),
		ServiceName: info.ServiceName,
		Weight:      float64(info.Weight),
		Enable:      true,
		Healthy:     true,
		Metadata:    info.Tags,
	})
	if e != nil {
		return fmt.Errorf("RegisterInstance err:%v", e)
	}
	return nil
}

func (n *nacosRegistry) Deregister(info *registry.Info) error {
	host, port, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		return err
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return err
	}
	if _, err = n.cli.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          host,
		Port:        uint64(p),
		ServiceName: info.ServiceName,
	}); err != nil {
		return err
	}
	return nil
}
