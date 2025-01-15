package category

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/api/category/v1"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
)

func (c *ControllerV1) GetCategoryNameList(ctx context.Context, req *v1.GetCategoryNameListReq) (res *v1.GetCategoryNameListRes, err error) {
	res = &v1.GetCategoryNameListRes{}
	allList := &v1.GetListRes{}
	err = dao.Category.Ctx(ctx).Fields(dao.Category.Columns().Name).Scan(&allList.List)
	for _, category := range allList.List {
		res.Name = append(res.Name, category.Name)
	}
	return
}
