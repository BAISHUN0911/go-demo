package handler

import (
	"gin-demo/model/system/request"
	"gin-demo/service"

	"github.com/gin-gonic/gin"
)

func SysAuthorities(c *gin.Context) {
	var req request.SysAuthoritiesReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "参数格式错误", "error": err.Error()})
		return
	}
	tree, error := service.GetRoleTree(req.TenantID)
	if error != nil {
		c.JSON(500, gin.H{"message": error.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": tree})
}
