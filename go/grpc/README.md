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
