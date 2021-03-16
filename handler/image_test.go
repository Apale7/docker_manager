package handler

import (
	"context"
	"docker_manager/proto/docker_manager"
	"fmt"
	"reflect"
	"testing"
)

var ctx = context.Background()

func TestCreateContainer(t *testing.T) {
	type args struct {
		ctx context.Context
		req *docker_manager.CreateContainerRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *docker_manager.CreateContainerResponse
		wantErr  bool
	}{
		{
			name: "create container python",
			args: args{ctx: ctx, req: &docker_manager.CreateContainerRequest{
				ContainerName: "apale5",
				UserId:        1,
				Username:      "apale",
				ImageId:       "sha256:80d28bedfe5dec59da9ebf8e6260224ac9008ab5c11dbbe16ee3ba3e4439ac2c",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateContainer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDeleteContainer(t *testing.T) {
	type args struct {
		ctx context.Context
		req *docker_manager.DeleteContainerRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *docker_manager.DeleteContainerResponse
		wantErr  bool
	}{
		{
			name:    "delete apale3",
			args:    args{ctx: ctx, req: &docker_manager.DeleteContainerRequest{UserId: 1, ContainerId: "879ec25f61a6bc63cba72a4fc72676fd1b18a98c81d3c20142ab00b8e84bf06e"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DeleteContainer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("DeleteContainer() gotResp = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}

func TestDeleteImage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *docker_manager.DeleteImageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *docker_manager.DeleteImageResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := DeleteImage(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DeleteImage() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetContainer(t *testing.T) {
	type args struct {
		ctx context.Context
		req *docker_manager.GetContainerRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *docker_manager.GetContainerResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetContainer(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContainer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetContainer() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestGetImage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *docker_manager.GetImageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *docker_manager.GetImageResponse
		wantErr  bool
	}{
		{
			name: "get all",
			args: args{ctx: ctx, req: &docker_manager.GetImageRequest{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := GetImage(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp == nil || len(gotResp.Images) <= 0 {
				t.Errorf("GetImage() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			for _, tt := range gotResp.Images {
				fmt.Println(tt.Id)
			}
		})
	}
}

func TestCreateImage(t *testing.T) {
	type args struct {
		ctx context.Context
		req *docker_manager.CreateImageRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *docker_manager.CreateImageResponse
		wantErr  bool
	}{
		{
			name: "create image apale_img1",
			args: args{ctx: ctx, req: &docker_manager.CreateImageRequest{
				UserId:     1,
				Dockerfile: []byte("FROM ubuntu:18.04\nCMD [\"/bin/bash\"]"),
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateImage(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("CreateImage() = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}
