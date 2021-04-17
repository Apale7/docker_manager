package rpc

import (
	containerManager "docker_manager/proto/container_server"
	user_center "docker_manager/proto/user-center"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	containerManagerClient containerManager.ManagerClient
	userCenterClient       user_center.UserCenterClient
)

var (
	containerManagerAddr string = "[::]:8666"
	userCenterAddr       string = ":9999"
)

func init() {
	containerManagerClient = containerManager.NewManagerClient(getConn(containerManagerAddr))
	// userCenterClient = user_center.NewUserCenterClient(getConn("111.230.172.240:9999"))
	userCenterClient = user_center.NewUserCenterClient(getConn(userCenterAddr))
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

func initAddrs() {
	userCenterAddr = os.Getenv("user_center_addr")
	if userCenterAddr == "" {
		userCenterAddr = ":9999"
	}

	fmt.Println(userCenterAddr)
}
