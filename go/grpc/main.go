package main

import (
	"context"
	"net"

	"./proto"
	"github.com/go-kit/kit/endpoint"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

type TpuserServer struct {
	getHandler  grpc_transport.Handler
	postHandler grpc_transport.Handler
}

func (s *TpuserServer) Get(ctx context.Context, in *proto.TpuserRequest) (*proto.TpuserResponse, error) {
	_, rsp, err := s.getHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.TpuserResponse), err
}

func (s *TpuserServer) Post(ctx context.Context, in *proto.TpuserRequest) (*proto.TpuserResponse, error) {
	_, rsp, err := s.postHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*proto.TpuserResponse), err
}

func getEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		bl := new(proto.TpuserResponse)
		bl.Uid = 1
		bl.Unionid = "go-kit入门到精通"
		return bl, nil
	}
}

func postEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*proto.TpuserRequest)
		b := new(proto.TpuserResponse)
		b.Uid = req.Uid
		b.Unionid = "Go与微服务"
		return b, nil
	}
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func main() {

	userServer := new(TpuserServer)
	getHandler := grpc_transport.NewServer(
		getEndpoint(),
		decodeRequest,
		encodeResponse,
	)
	userServer.getHandler = getHandler

	//创建bookInfo的Handler
	postHandler := grpc_transport.NewServer(
		postEndpoint(),
		decodeRequest,
		encodeResponse,
	)

	userServer.postHandler = postHandler

	//启动grpc服务
	serviceAddress := ":50052"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer()
	proto.RegisterTpuserServer(gs, userServer)
	gs.Serve(ls)
}
