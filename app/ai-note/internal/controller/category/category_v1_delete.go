package category

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"

	"github.com/wangle201210/ai-note/app/ai-note/api/category/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.Category.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
