package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/apiserver/handler"
	"github.com/xiaozefeng/apiserver/model"
	"github.com/xiaozefeng/apiserver/pkg/errno"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	u, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
	}
	handler.SendResponse(c, nil, u)
}
