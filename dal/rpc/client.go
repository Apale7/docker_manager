package rpc

import (
	containerManager "docker_manager/proto/container_server"
	"docker_manager/proto/user-center"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	containerManagerClient containerManager.ManagerClient
	userCenterClient       user_center.UserCenterClient
)

func init() {
	containerManagerClient = containerManager.NewManagerClient(getConn("193.112.177.167:8666"))
	userCenterClient = user_center.NewUserCenterClient(getConn("111.230.172.240:9999"))
}

func getConn(addr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
	return conn
}
