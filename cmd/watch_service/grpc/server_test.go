package grpc

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"github.com/liuqianhong6007/k8s/cmd/watch_service/grpc/protocol"
)

func newWatchServiceClient(addr string) protocol.WatchServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return protocol.NewWatchServiceClient(conn)
}

func TestServer_WatchPod(t *testing.T) {
	client := newWatchServiceClient("localhost:18083")
	stream, err := client.WatchPod(context.Background(), &protocol.MatchCondition{
		Namespace:     "default",
		LabelSelector: "app=k8s-service",
	})
	if err != nil {
		t.Fatal(err)
	}
	for {
		rsp, err := stream.Recv()
		if err != nil {
			t.Fatal(err)
		}
		if rsp.Status == protocol.WatchResponse_Ok {
			t.Logf("name: %s,addr: %s,action: %d", rsp.Name, rsp.Addr, rsp.Action)
		} else {
			t.Logf("status: %d", rsp.Status)
		}
	}
}
