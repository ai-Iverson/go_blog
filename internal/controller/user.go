package controller

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	v1 "go_blog/api/v1"
	"go_blog/internal/service"
	"go_blog/utility/token"
	"go_blog/utility/utils"
	"strconv"
)

var User = cUser{}

type cUser struct {
}

func (c cUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	user, err := service.User().Login(ctx, req.Name, req.Password)
	res.User = &v1.UserGetInfoRes{}
	glog.Info(ctx, "df", user)
	utils.MyCopy(ctx, res.User, user)
	glog.Info(ctx, "df", res)
	userKey := strconv.Itoa(user.Id)
	MyToken, err := token.Instance().GenerateToken(ctx, userKey, user)
	res.Token = "Bearer " + MyToken.Token
	return
}
