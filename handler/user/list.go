package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/apiserver/handler"
	"github.com/xiaozefeng/apiserver/pkg/errno"
	"github.com/xiaozefeng/apiserver/service"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	infoList, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infoList,
	})
}
