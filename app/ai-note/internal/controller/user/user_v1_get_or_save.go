package user

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"

	"github.com/wangle201210/ai-note/app/ai-note/api/user/v1"
)

func (c *ControllerV1) GetOrSave(ctx context.Context, req *v1.GetOrSaveReq) (res *v1.GetOrSaveRes, err error) {
	user := &entity.User{}
	dao.User.Ctx(ctx).Fields(dao.User.Columns().Id).Where(do.User{Name: req.Name}).Scan(&user)
	if user.Id == 0 {
		dao.User.Ctx(ctx).Data(do.User{
			Name:   req.Name,
			OpenId: req.OpenId,
		}).Insert()
	}
	return
}
