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

package option

type Options struct {
	cluster string
	group   string
}

// NewOptions returns a new Options instance.
func NewOptions(cluster, group string) Options {
	return Options{
		cluster: cluster,
		group:   group,
	}
}

// GetCluster returns the cluster name.
func (o Options) GetCluster() string {
	return o.cluster
}

// GetGroup returns the group name.
func (o Options) GetGroup() string {
	return o.group
}

// Option is nacos option.
type Option func(o *Options)

// WithCluster with cluster option.
func WithCluster(cluster string) Option {
	return func(o *Options) { o.cluster = cluster }
}

// WithGroup with group option.
func WithGroup(group string) Option {
	return func(o *Options) { o.group = group }
}
