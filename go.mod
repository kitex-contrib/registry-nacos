module github.com/kitex-contrib/registry-nacos

go 1.16

replace (
	github.com/kitex-contrib/registry-nacos/registry => ./registry
	github.com/kitex-contrib/registry-nacos/resolver => ./resolver
)

require (
	github.com/cloudwego/kitex v0.0.8
	github.com/nacos-group/nacos-sdk-go v1.0.9
	github.com/stretchr/testify v1.7.0
)
