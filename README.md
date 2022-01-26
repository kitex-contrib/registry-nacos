# registry-nacos (*This is a community driven project*)

Nacos as service discovery.

## How to use?

### Server

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
        RotateTime:          "1h",
        MaxAge:              3,
        LogLevel:            "info",
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
    
    svr := echo.NewServer(new(EchoImpl), server.WithRegistry(registry.NewNacosRegistry(cli)))
    if err := svr.Run(); err != nil {
        log.Println("server stopped with error:", err)
    } else {
        log.Println("server stopped")
    }
    // ...
}

```

### Client

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
            RotateTime:          "1h",
            MaxAge:              3,
            LogLevel:            "info",
	}

	cli,err := clients.NewNamingClient(
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
## Compatibility
The server of Nacos2.0 is fully compatible with 1.X nacos-sdk-go. [see](https://nacos.io/en-us/docs/2.0.0-compatibility.html)


maintained by: [baiyutang](https://github.com/baiyutang)
