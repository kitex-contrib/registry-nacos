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
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
)

type customNacosLogger struct{}

func NewCustomNacosLogger() logger.Logger {
	return customNacosLogger{}
}

func (m customNacosLogger) Info(args ...interface{}) {
	klog.Info(args...)
}

func (m customNacosLogger) Warn(args ...interface{}) {
	klog.Warn(args...)
}

func (m customNacosLogger) Error(args ...interface{}) {
	klog.Error(args...)
}

func (m customNacosLogger) Debug(args ...interface{}) {
	klog.Debug(args)
}

func (m customNacosLogger) Infof(fmt string, args ...interface{}) {
	klog.Infof(fmt, args...)
}

func (m customNacosLogger) Warnf(fmt string, args ...interface{}) {
	klog.Warnf(fmt, args...)
}

func (m customNacosLogger) Errorf(fmt string, args ...interface{}) {
	klog.Errorf(fmt, args...)
}

func (m customNacosLogger) Debugf(fmt string, args ...interface{}) {
	klog.Debugf(fmt, args...)
}
