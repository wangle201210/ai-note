package note

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	_, err = dao.Note.Ctx(ctx).Data(do.Note{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Amount:     req.Amount,
		OrgContent: req.OrgContent,
		Status:     req.Status,
	}).WherePri(req.Id).Update()
	return
}
