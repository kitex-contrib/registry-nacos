package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// NewNacosClient  Create a Nacos client
func NewNacosClient() (naming_client.INamingClient, error) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(NacosAddr(), uint64(NacosPort())),
	}
	cc := constant.ClientConfig{
		NamespaceId: NacosNameSpaceId(),
		RegionId:    DefaultRegionId,
	}
	// todo add kitex logger
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
