package main

import (
	"context"
	"net"
	"os"

	h "docker_manager/proto/docker_manager"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	addr   = ":8888"
	conn   *grpc.ClientConn
	ctx    context.Context
	cancel context.CancelFunc
	err    error
)

func main() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %+v", err)
	}
	server := grpc.NewServer()
	h.RegisterDockerManagerServer(server, &DockerManagerServer{})
	logrus.Infof("server listen at %s", lis.Addr().String())
	server.Serve(lis)
}

func init() {
	tmpAddr := os.Getenv("docker_manager_addr")
	if tmpAddr != "" {
		addr = tmpAddr
	}
}
