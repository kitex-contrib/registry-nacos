package nacos

import (
	"os"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	EnvServiceAddr     = "serviceAddr"
	EnvPort            = "servicePort"
	EnvNameSpaceId     = "namespace"
	DefaultServiceAddr = "127.0.0.1"
	DefaultPort        = 8848
	DefaultRegionId    = "cn-hangzhou"
)

// NacosPort Get Nacos port from environment variables
func NacosPort() int64 {
	portText := os.Getenv(EnvPort)
	if len(portText) == 0 {
		return DefaultPort
	}
	port, err := strconv.ParseInt(portText, 10, 64)
	if err != nil {
		klog.Errorf("ParseInt failed,err=%+v", err)
		return DefaultPort
	}
	return port
}

// NacosAddr Get Nacos addr from environment variables
func NacosAddr() string {
	addr := os.Getenv(EnvServiceAddr)
	if len(addr) == 0 {
		return DefaultServiceAddr
	}
	return addr
}

// NacosNameSpaceId Get Nacos namespace id from environment variables
func NacosNameSpaceId() string {
	return os.Getenv(EnvNameSpaceId)
}
