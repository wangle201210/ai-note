package category

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"

	"github.com/wangle201210/ai-note/app/ai-note/api/category/v1"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	res = &v1.GetOneRes{}
	err = dao.Category.Ctx(ctx).WherePri(req.Id).Scan(&res.Category)
	return
}
