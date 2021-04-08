## 定义生成服务

### 使用 truss

https://github.com/tuneinc/truss

* linux
```shell
go get -u -d github.com/metaverse/truss
cd $GOPATH/src/github.com/metaverse/truss
make dependencies
make
```

* win
```shell
go get -u -d github.com/metaverse/truss
cd %GOPATH%/src/github.com/metaverse/truss
wininstall.bat
```

```shell
truss tpuser.proto
```

###  使用 protoc

```shell
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
protoc tpuser.proto --go_out=plugins=grpc:.
```

* 生成接口

```go
type TpuserServer interface {
	Get(context.Context, *TpuserRequest) (*TpuserResponse, error)
	Post(context.Context, *TpuserRequest) (*TpuserResponse, error)
}

type TpuserClient interface {
	Get(ctx context.Context, in *TpuserRequest, opts ...grpc.CallOption) (*TpuserResponse, error)
	Post(ctx context.Context, in *TpuserRequest, opts ...grpc.CallOption) (*TpuserResponse, error)
}
```