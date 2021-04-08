package test

import (
	"context"
	"fmt"
	"testing"

	"../proto"
	"google.golang.org/grpc"
)

func TestGet(t *testing.T) {
	conn, _ := grpc.Dial(":50052", grpc.WithInsecure())
	defer conn.Close()

	client := proto.NewTpuserClient(conn)
	resp, _ := client.Get(context.Background(), &proto.TpuserRequest{
		Uid:     1,
		Unionid: "test",
	})
	fmt.Println(resp)
}
