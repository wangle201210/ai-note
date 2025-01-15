package note

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = dao.Note.Ctx(ctx).WherePri(req.Id).Delete()
	return
}
