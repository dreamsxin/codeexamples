package main

import (
	"context"
	"log"
	"fmt"
	"flag"
	"net"

	"./proto"
	//"github.com/go-kit/kit/endpoint"
	//grpc_transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	//"github.com/golang/protobuf/proto"
)


type TpuserServer struct {
}


func (s *TpuserServer) Get(ctx context.Context, in *proto.TpuserRequest) (*proto.TpuserResponse, error) {
		bl := new(proto.TpuserResponse)
		bl.Uid = 1
		bl.Unionid = "go-kit入门到精通"
		return bl, nil
}

func (s *TpuserServer) Post(ctx context.Context, in *proto.TpuserRequest) (*proto.TpuserResponse, error) {
		b := new(proto.TpuserResponse)
		b.Uid = in.Uid
		b.Unionid = "Go与微服务"
		return b, nil
}

/*
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
*/

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	port       = flag.Int("port", 50052, "The server port")
)

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		/*
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
		*/
	}

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterTpuserServer(grpcServer, new(TpuserServer))
	grpcServer.Serve(listen)

	/*
	server := new(TpuserServer)

	getHandler := grpc_transport.NewServer(
		getEndpoint(),
		decodeRequest,
		encodeResponse,
	)
	server.getHandler = getHandler

	//创建bookInfo的Handler
	postHandler := grpc_transport.NewServer(
		postEndpoint(),
		decodeRequest,
		encodeResponse,
	)

	server.postHandler = postHandler

	//启动grpc服务
	serviceAddress := ":50052"
	ls, _ := net.Listen("tcp", serviceAddress)
	gs := grpc.NewServer()
	proto.RegisterTpuserServer(gs, server)
	gs.Serve(ls)
	*/
}
