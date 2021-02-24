package handler

import (
	"context"
	"docker_manager/dal/rpc"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
)

func GetAllContainers(c *gin.Context) {
	ctx := context.Background()
	containers, err := rpc.GetAllContainers(ctx)
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"code": 0, "containers": containers})
}
