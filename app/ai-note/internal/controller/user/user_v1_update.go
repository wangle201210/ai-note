package user

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"

	"github.com/wangle201210/ai-note/app/ai-note/api/user/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Name:   req.Name,
		Status: req.Status,
		OpenId: req.OpenId,
	}).WherePri(req.Id).Update()
	return
}
