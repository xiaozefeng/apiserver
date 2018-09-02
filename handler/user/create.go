package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/xiaozefeng/apiserver/handler"
	"github.com/xiaozefeng/apiserver/model"
	"github.com/xiaozefeng/apiserver/pkg/errno"
	"github.com/xiaozefeng/apiserver/util"
)

func Create(c *gin.Context) {
	log.Info("User Create function called. ", lager.Data{"X-Request-Id": util.GetReqId(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// Insert the user to the database
	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, resp)
}
