package user

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaozefeng/apiserver/handler"
	"github.com/xiaozefeng/apiserver/model"
	"github.com/xiaozefeng/apiserver/pkg/auth"
	"github.com/xiaozefeng/apiserver/pkg/errno"
	"github.com/xiaozefeng/apiserver/pkg/token"
)

// Login generates the authentication token
// if the password was matched with the specified account
func Login(c *gin.Context) {
	// Binding the data with user struct
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
	}

	// Get user info by the username
	user, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password
	if err := auth.Compare(user.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token
	t, err := token.Sign(c, token.Context{ID: user.Id, Username: user.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}
	handler.SendResponse(c, nil, model.Token{Token: t})

}
