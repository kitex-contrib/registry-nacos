# registry-nacos (*This is a community driven project*)

[中文](https://github.com/kitex-contrib/registry-nacos/blob/main/README_CN.md)

Nacos as service discovery.

## How to use?

### Basic

#### Server

```go
import (
    // ...
    "github.com/kitex-contrib/registry-nacos/registry"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
    "github.com/cloudwego/kitex/pkg/rpcinfo"
    // ...
)

func main() {
    // ... 
    r, err := registry.NewDefaultNacosRegistry()
    if err != nil {
        panic(err)
    }
    svr := echo.NewServer(
        new(EchoImpl), 
        server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "echo"}),
        server.WithRegistry(r), 
	)
    if err := svr.Run(); err != nil {
        log.Println("server stopped with error:", err)
    } else {
        log.Println("server stopped")
    }
    // ...
}
```

#### Client

```go
import (
    // ...
    "github.com/kitex-contrib/registry-nacos/resolver"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
    // ...
)

func main() {
    // ... 
    r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
	    panic(err)	
    }
    client, err := echo.NewClient("echo", client.WithResolver(r))
    if err != nil {
        log.Fatal(err)
    }
    // ...
}
```

### Custom Nacos Client Configuration

#### Server

```go
import (
    // ...
    "github.com/kitex-contrib/registry-nacos/registry"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
    // ...
)
func main() {
    // ...
    sc := []constant.ServerConfig{
        *constant.NewServerConfig("127.0.0.1", 8848),
    }
    
    cc := constant.ClientConfig{
        NamespaceId:         "public",
        TimeoutMs:           5000,
        NotLoadCacheAtStart: true,
        LogDir:              "/tmp/nacos/log",
        CacheDir:            "/tmp/nacos/cache",
        LogLevel:            "info",
        Username:            "your-name",
        Password:            "your-password",
    }
    
    cli, err := clients.NewNamingClient(
        vo.NacosClientParam{
            ClientConfig:  &cc,
            ServerConfigs: sc,
        },
    )
    if err != nil {
        panic(err)
    }
    
    svr := echo.NewServer(new(EchoImpl), 
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "echo"}),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
    )
    if err := svr.Run(); err != nil {
        log.Println("server stopped with error:", err)
    } else {
        log.Println("server stopped")
    }
    // ...
}
```

#### Client

```go
import (
    // ...
    "github.com/kitex-contrib/registry-nacos/resolver"
    "github.com/nacos-group/nacos-sdk-go/clients"
    "github.com/nacos-group/nacos-sdk-go/clients/naming_client"
    "github.com/nacos-group/nacos-sdk-go/common/constant"
    "github.com/nacos-group/nacos-sdk-go/vo"
    // ...
)
func main() {
    // ... 
    sc := []constant.ServerConfig{
	    *constant.NewServerConfig("127.0.0.1", 8848),
	}
    cc := constant.ClientConfig{
        NamespaceId:         "public",
        TimeoutMs:           5000,
        NotLoadCacheAtStart: true,
        LogDir:              "/tmp/nacos/log",
        CacheDir:            "/tmp/nacos/cache",
        LogLevel:            "info",
        Username:            "your-name",
        Password:            "your-password",
    }
    
    cli, err := clients.NewNamingClient(
        vo.NacosClientParam{
            ClientConfig:  &cc,
            ServerConfigs: sc,
        },
    )
    if err != nil {
	    panic(err)	
    }
    client, err := echo.NewClient("echo", client.WithResolver(resolver.NewNacosResolver(cli))
    if err != nil {
        log.Fatal(err)
    }
    // ...
}
```

### Environment Variable

| Environment Variable Name | Environment Variable Default Value | Environment Variable Introduction |
| ------------------------- | ---------------------------------- | --------------------------------- |
| serverAddr               | 127.0.0.1                          | nacos server address              |
| serverPort               | 8848                               | nacos server port                 |
| namespace                 |                                    | the namespaceId of nacos          |
| KITEX_NACOS_ENV_TAGS           | ""                                  | support inject default tags|

### More Info

Refer to [example](example) for more usage.

## Compatibility
This Package use Nacos1.x client. The Nacos2.0 and Nacos1.0 Server are fully compatible with it. [see](https://nacos.io/en-us/docs/v2/upgrading/2.0.0-compatibility.html)

maintained by: [baiyutang](https://github.com/baiyutang)
