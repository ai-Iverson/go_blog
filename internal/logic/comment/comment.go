package comment

import (
	"context"
	"github.com/gogf/gf/v2/encoding/ghash"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"go_blog/internal/dao"
	"go_blog/internal/model"
	"go_blog/internal/model/entity"
	"go_blog/internal/service"
)

type sComment struct {
}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) Comment(ctx context.Context, in model.CommentInput) (err error) {
	// 通过Nickname映射头像
	var avatar string
	nickname := in.Nickname
	nicknameBytes := gconv.Bytes(nickname)
	nicknameHash := ghash.DJB(nicknameBytes)
	switch nicknameHash % 5 {
	case 0:
		avatar =
			"https://cdn.naccl.top/blog/comment/avatar/fc85d2c5-e3d0-40a5-913f-95a3b400db4d.jpeg"
	case 1:
		avatar =
			"https://cdn.naccl.top/blog/comment/avatar/72861ed5-94c5-49f8-b1e7-82edc810f94f.jpeg"
	case 2:
		avatar = "https://cdn.naccl.top/blog/comment/avatar/4330aa33-de59-4a65-9ef4-ed06424e1b30.jpeg"
	case 3:
		avatar = "https://cdn.naccl.top/blog/comment/avatar/60d06f07-aefe-4a99-9900-b0b97497168f.jpeg"
	case 4:
		avatar = "https://cdn.naccl.top/blog/comment/avatar/a9474efb-61ca-4e5a-8c2e-ad1e39dd6d20.jpeg"
	case 5:
		avatar = "https://cdn.naccl.top/blog/comment/avatar/d25ec1e1-d100-452a-a45d-c6ff1f352172.jpeg"

	}

	// 获取用户ip
	requestIp := g.RequestFromCtx(ctx).GetClientIp()

	var isAdminComment = 0
	if nickname == "admin" {
		isAdminComment = 1
	}
	_, err = dao.Comment.Ctx(ctx).Data(entity.Comment{
		Nickname:        nickname,
		Email:           in.Email,
		Content:         in.Content,
		Avatar:          avatar,
		CreateTime:      gtime.Now(),
		Ip:              requestIp,
		IsPublished:     1,
		IsAdminComment:  isAdminComment,
		Page:            in.Page,
		IsNotice:        gconv.Int(in.Notice),
		BlogId:          gconv.Int64(in.BlogID),
		ParentCommentId: gconv.Int64(in.ParentCommentID),
		Website:         in.Website,
	}).Insert()

	return
}

func (s *sComment) CommentList(ctx context.Context) (out *model.CommentListOutput, err error) {
	requestPayload := g.RequestFromCtx(ctx).GetMap()
	res, _ := dao.Comment.Ctx(ctx).Where(dao.Comment.Columns().BlogId, requestPayload["blogId"]).All()
	glog.Info(ctx, res)
	return
}
