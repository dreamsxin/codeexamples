## protobuffer

```shell
git clone https://github.com/golang/protobuf.git $GOPATH/src/github.com/golang/protobuf
# 安装protoc 
go install github.com/golang/protobuf/proto
#安装插件
go install github.com/golang/protobuf/protoc-gen-go
protoc --go_out=. *.proto //如果不加grpc就没有rpc代码实现
protoc --go_out=plugins=grpc:. *.proto //如果不加grpc就没有rpc代码实现

go get -u github.com/golang/protobuf/protoc-gen-go
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```

## gogoprotobuf

* 安装插件
gogoprotobuf有两个插件可以使用

protoc-gen-gogo：和protoc-gen-go生成的文件差不多，性能也几乎一样(稍微快一点点)

protoc-gen-gofast：生成的文件更复杂，性能也更高(快5-7倍) 
```shell
//gogo
go get github.com/gogo/protobuf/protoc-gen-gogo
//gofast
go get github.com/gogo/protobuf/protoc-gen-gofast
```
安装gogoprotobuf库文件
```shell
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/gogoproto //这个不装也没关系
```
生成go文件
```shell
//gogo
protoc --gogo_out=. *.proto
//gofast
protoc --gofast_out=. *.proto 
```
## 代理

https://github.com/golang/go/wiki/Modules#are-there-always-on-module-repositories-and-enterprise-proxies

Goproxy 中国

如果您使用的是 MAC/Linux, 终端中设置如下环境变量:
`export GOPROXY=https://proxy.golang.com.cn,direct`

如果您使用的是 Windows 系统, 终端中执行如下 Go 命令:
`go env -w GOPROXY=https://proxy.golang.com.cn,direct`

```shell
git clone https://github.com/grpc/grpc-go ./google.golang.org/grpc
git clone https://github.com/golang/net.git ./golang.org/x/net
git clone https://github.com/google/go-genproto.git ./google.golang.org/genproto
git clone https://github.com/golang/text.git ./golang.org/x/text
go install google.golang.org/grpc
```

## 常用库

https://github.com/go-kit/kit

https://github.com/stretchr/testify

https://github.com/prometheus/client_golang

https://github.com/opentracing/opentracing-go
