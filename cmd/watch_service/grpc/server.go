package grpc

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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

func (s *Server) WatchPod(in *protocol.MatchCondition, stream protocol.WatchService_WatchPodServer) error {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic: ", err)
		}
	}()

	namespace := in.GetNamespace()
	label := in.GetLabelSelector()
	if namespace == "" || label == "" {
		stream.Send(&protocol.WatchResponse{
			Status: protocol.WatchResponse_ParamError,
		})
		return nil
	}

	watcher, err := s.clientset.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: label,
	})
	if err != nil {
		stream.Send(&protocol.WatchResponse{
			Status: protocol.WatchResponse_UnknownError,
		})
		log.Println(err)
		return err
	}

	for {
		event, ok := <-watcher.ResultChan()
		if !ok {
			stream.Send(&protocol.WatchResponse{
				Status: protocol.WatchResponse_UnknownError,
			})
			log.Println("watch_service stop")
			watcher.Stop()
			return nil
		}

		switch event.Type {
		case watch.Added:
			log.Println("receive add event")
			pod, ok := event.Object.(*v1.Pod)
			if !ok {
				continue
			}
			buff, _ := json.Marshal(pod)
			stream.Send(&protocol.WatchResponse{
				Status: protocol.WatchResponse_Ok,
				Action: protocol.WatchResponse_Add,
				Name:   pod.Name,
				Detail: string(buff),
			})

		case watch.Deleted:
			log.Println("receive delete event")
			pod, ok := event.Object.(*v1.Pod)
			if !ok {
				continue
			}
			buff, _ := json.Marshal(pod)
			stream.Send(&protocol.WatchResponse{
				Status: protocol.WatchResponse_Ok,
				Action: protocol.WatchResponse_Delete,
				Name:   pod.Name,
				Detail: string(buff),
			})
		case watch.Modified:
			log.Println("receive modify event")

		case watch.Error:
			log.Println("receive error event")

		case watch.Bookmark:
			log.Println("receive bookmark event")
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

func (s *Server) SetPodLabels() {
	namespace := os.Getenv("NAMESPACE")
	podName := os.Getenv("HOSTNAME")
	log.Println("pod namespace: ", namespace)
	log.Println("pod name: ", podName)

	_, err := s.clientset.CoreV1().Pods(namespace).Patch(context.Background(), podName, types.MergePatchType, []byte(`{ "metadata": { "labels": { "status": "ready" } }}`), metav1.PatchOptions{})
	if err != nil {
		log.Println(err)
	}
}
