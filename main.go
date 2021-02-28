package main

import (
	"context"
	"net"

	h "docker_manager/proto/docker_manager"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8888"
)

var (
	conn   *grpc.ClientConn
	ctx    context.Context
	cancel context.CancelFunc
	err    error
)

func main() {
	logrus.SetReportCaller(true)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %+v", err)
	}
	server := grpc.NewServer()
	h.RegisterDockerManagerServer(server, &DockerManagerServer{})
	logrus.Infof("server listen at %s", lis.Addr().String())
	server.Serve(lis)
}
