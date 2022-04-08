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
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/kitex-contrib/registry-nacos/nacos"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type options struct {
	cluster string
	group   string
}

// Option is nacos option.
type Option func(o *options)

// WithCluster with cluster option.
func WithCluster(cluster string) Option {
	return func(o *options) { o.cluster = cluster }
}

// WithGroup with group option.
func WithGroup(group string) Option {
	return func(o *options) { o.group = group }
}

type nacosRegistry struct {
	cli  naming_client.INamingClient
	opts options
}

// NewDefaultNacosRegistry create a default service registry using nacos.
func NewDefaultNacosRegistry() (registry.Registry, error) {
	cli, err := nacos.NewDefaultNacosClient()
	if err != nil {
		return nil, err
	}
	return NewNacosRegistry(cli), nil
}

// NewNacosRegistry create a new registry using nacos.
func NewNacosRegistry(cli naming_client.INamingClient, opts ...Option) registry.Registry {
	op := options{
		cluster: "DEFAULT",
		group:   "DEFAULT_GROUP",
	}
	for _, option := range opts {
		option(&op)
	}
	return &nacosRegistry{cli: cli, opts: op}
}

var _ registry.Registry = (*nacosRegistry)(nil)

// Register service info to nacos.
func (n *nacosRegistry) Register(info *registry.Info) error {
	if info == nil {
		return errors.New("registry.Info can not be empty")
	}
	if info.ServiceName == "" {
		return errors.New("registry.Info ServiceName can not be empty")
	}
	if info.Addr == nil {
		return errors.New("registry.Info Addr can not be empty")
	}
	host, port, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		return fmt.Errorf("parse registry info addr error: %w", err)
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("parse registry info port error: %w", err)
	}
	if host == "" {
		host, err = n.getLocalIpv4Host()
		if err != nil {
			return fmt.Errorf("parse registry info addr error: %w", err)
		}
	}
	_, e := n.cli.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(p),
		ServiceName: info.ServiceName,
		Weight:      float64(info.Weight),
		Enable:      true,
		Healthy:     true,
		Metadata:    info.Tags,
		GroupName:   n.opts.group,
		ClusterName: n.opts.cluster,
		Ephemeral:   true,
	})
	if e != nil {
		return fmt.Errorf("register instance error: %w", e)
	}
	return nil
}

func (n *nacosRegistry) getLocalIpv4Host() (string, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addr {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", errors.New("not found ipv4 address")
}

// Deregister a service info from nacos.
func (n *nacosRegistry) Deregister(info *registry.Info) error {
	host, port, err := net.SplitHostPort(info.Addr.String())
	if err != nil {
		return err
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return fmt.Errorf("parse registry info port error: %w", err)
	}
	if host == "" {
		host, err = n.getLocalIpv4Host()
		if err != nil {
			return fmt.Errorf("parse registry info addr error: %w", err)
		}
	}
	if _, err = n.cli.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          host,
		Port:        uint64(p),
		ServiceName: info.ServiceName,
		Ephemeral:   true,
		GroupName:   n.opts.group,
		Cluster:     n.opts.cluster,
	}); err != nil {
		return err
	}
	return nil
}
