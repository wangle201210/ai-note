package note

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	insertId, err := dao.Note.Ctx(ctx).Data(do.Note{
		UserId:     req.UserId,
		CategoryId: req.CategoryId,
		Amount:     req.Amount,
		OrgContent: req.OrgContent,
		Status:     v1.StatusOK,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{
		Id: insertId,
	}
	return
}
