package user

import (
	"context"
	v1 "github.com/wangle201210/ai-note/app/ai-note/api/user/v1"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	insertId, err := dao.User.Ctx(ctx).Data(do.User{
		Name:   req.Name,
		Status: v1.StatusOK,
		OpenId: req.OpenId,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		Id: insertId,
	}
	return
}
