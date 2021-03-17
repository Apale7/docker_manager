package rpc

import (
	"context"
	user_center "docker_manager/proto/user-center"
)

func GetGroup(ctx context.Context, groupInfo *user_center.Group, memberID uint) (groups []*user_center.Group, err error) {
	req := &user_center.GetGroupRequest{
		GroupInfo: groupInfo,
		MemberId: uint32(memberID),
	}
	resp, err := userCenterClient.GetGroup(ctx, req)
	if err != nil {
		return
	}

	return resp.Groups, nil
}
