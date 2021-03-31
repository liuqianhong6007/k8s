package grpc

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"

	"github.com/liuqianhong6007/k8s/cmd/watch_service/grpc/protocol"
	"github.com/liuqianhong6007/k8s/util"
)

func NewServer(addr string) *Server {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	clientset := util.NewKubernetesClientset(true, "")
	grpcServer := grpc.NewServer()
	server := &Server{
		addr:       addr,
		lis:        lis,
		grpcServer: grpcServer,
		clientset:  clientset,
	}
	protocol.RegisterWatchServiceServer(grpcServer, server)
	return server
}

type Server struct {
	protocol.UnimplementedWatchServiceServer
	addr       string
	lis        net.Listener
	grpcServer *grpc.Server
	clientset  *kubernetes.Clientset
}

func (s *Server) Watch(in *protocol.MatchPodCondition, stream protocol.WatchService_WatchServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic: ", err)
		}
	}()

	namespace := in.GetNamespace()
	label := in.GetLabelSelector()
	if namespace == "" || label == "" {
		stream.Send(&protocol.MatchPodResponse{
			Status: protocol.MatchPodResponse_ParamError,
		})
		return nil
	}

	podList, err := s.clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		stream.Send(&protocol.MatchPodResponse{
			Status: protocol.MatchPodResponse_UnknownError,
		})
		log.Println(err)
		return err
	}

	podMap := make(map[string]interface{})
	for _, pod := range podList.Items {
		podMap[pod.Name] = pod
		buf, _ := json.Marshal(pod)
		stream.Send(&protocol.MatchPodResponse{
			Status: protocol.MatchPodResponse_Ok,
			Name:   pod.Name,
			Detail: string(buf),
			Action: protocol.MatchPodResponse_Add,
		})
	}

	watcher, err := s.clientset.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		stream.Send(&protocol.MatchPodResponse{
			Status: protocol.MatchPodResponse_UnknownError,
		})
		log.Println(err)
		return err
	}

	for {
		event, ok := <-watcher.ResultChan()
		if !ok {
			stream.Send(&protocol.MatchPodResponse{
				Status: protocol.MatchPodResponse_UnknownError,
			})
			log.Println("watch_service stop")
			watcher.Stop()
			return nil
		}

		switch event.Type {
		case watch.Added:
			stream.Send(&protocol.MatchPodResponse{
				Status: protocol.MatchPodResponse_Ok,
				Action: protocol.MatchPodResponse_Add,
			})
		case watch.Deleted:
			stream.Send(&protocol.MatchPodResponse{
				Status: protocol.MatchPodResponse_Ok,
				Action: protocol.MatchPodResponse_Delete,
			})
		}
	}
}

func (s *Server) Listen() {
	defer func() {
		log.Println("game service grpc server stop")
		s.grpcServer.Stop()
		s.lis.Close()
	}()
	log.Println("game service grpc server listen at: ", s.addr)
	if err := s.grpcServer.Serve(s.lis); err != nil {
		panic(err)
	}
}