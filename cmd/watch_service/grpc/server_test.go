package grpc

import (
	"context"
	"testing"
	"time"

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

func TestServer_Watch(t *testing.T) {
	client := newWatchServiceClient("localhost:8082")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stream, err := client.Watch(ctx, &protocol.MatchPodCondition{
		Namespace:     "default",
		LabelSelector: "app=k8s-test",
	})
	if err != nil {
		t.Fatal(err)
	}
	for {
		rsp, err := stream.Recv()
		if err != nil {
			t.Fatal(err)
		}
		if rsp.Status == protocol.MatchPodResponse_Ok {
			t.Logf("[INFO] podname: %s,action: %d,detail: %s\n", rsp.Name, rsp.Action, rsp.Detail)
		} else {
			t.Logf("[ERROR] status: %d", rsp.Status)
		}
	}
}
