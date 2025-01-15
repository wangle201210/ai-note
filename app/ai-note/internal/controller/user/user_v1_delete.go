package user

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"

	"github.com/wangle201210/ai-note/app/ai-note/api/user/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.User.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
