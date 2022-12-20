package user

import (
	"context"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
	"go_blog/utility/utils"
)

type sUser struct {
}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) Login(ctx context.Context, username, password string) (out *model.UserLoginOutput, err error) {
	out = &model.UserLoginOutput{}
	user := &entity.User{}
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username, "admin").Scan(user)
	utils.MyCopy(ctx, out, user)
	return
}
