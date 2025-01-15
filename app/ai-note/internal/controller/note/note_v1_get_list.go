package note

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	res = &v1.GetListRes{}
	err = dao.Note.Ctx(ctx).Where(do.Note{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Status:     req.Status,
	}).Scan(&res.List)
	return
}
