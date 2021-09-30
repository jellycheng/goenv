# goenv
```
分析.env格式文件及提供对应的方法获取值
```

## 获取
```
go get -u github.com/jellycheng/goenv
或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/goenv

```

## Documentation
[https://pkg.go.dev/github.com/jellycheng/goenv](https://pkg.go.dev/github.com/jellycheng/goenv)

## .env文件格式示例
```
APP_ENV=dev
APP_NAME=go-xxx-service
APP_PORT=9000

#日志保存目录
LOG_DIR=/data/logs/app/
LOG_LEVEL=debug

#common redis 配置
COMMON_REDIS_HOST = 127.0.0.1
COMMON_REDIS_PORT = 6379
COMMON_REDIS_DATABASE = 0
COMMON_REDIS_PASSWORD =

```

## 示例
```
package main

import (
	"fmt"
	"github.com/jellycheng/goenv"
	"os"
)

func main()  {
    // 加载env文件
    curDir,_ := os.Getwd()
    err := goenv.LoadEnv(curDir + "/.env")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    
    // 获取env值 - 方式1，返回字符串
    appEnv := goenv.GetString("APP_ENV")
    fmt.Println(appEnv)
    
    noExitEnv := goenv.GetString("APP_XXX")
    fmt.Println(noExitEnv)
    
    // 获取env值 - 方式2，返回gosupport.StrTo类型
    host := goenv.Get("COMMON_REDIS_HOST")
    fmt.Println(host.String())
    
    port := goenv.Get("COMMON_REDIS_PORT")
    fmt.Println(port.MustInt())

}

```
