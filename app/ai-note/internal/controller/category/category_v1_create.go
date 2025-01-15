package category

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"

	"github.com/wangle201210/ai-note/app/ai-note/api/category/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	_, err = dao.Category.Ctx(ctx).Data(do.Category{
		Name:   req.Name,
		Status: v1.StatusOK,
	}).InsertAndGetId()
	return
}
