package user

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"

	"github.com/wangle201210/ai-note/app/ai-note/api/user/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.User.Ctx(ctx).Where(do.User{
		Name:   req.Name,
		OpenId: req.OpenId,
		Status: req.Status,
	}).Scan(&res.List)
	return
}
