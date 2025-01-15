package note

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"
)

func (c *ControllerV1) GetNoteList(ctx context.Context, req *v1.GetNoteListReq) (res *v1.GetNoteListRes, err error) {
	res = &v1.GetNoteListRes{}
	user := &entity.User{}
	var list []*model.NoteInfo
	dao.User.Ctx(ctx).Fields(dao.User.Columns().Id).Where(do.User{Name: req.UserName}).Scan(&user)
	dao.Note.Ctx(ctx).Where(dao.Note.Columns().UserId, user.Id).WithAll().Scan(&list)
	for _, info := range list {
		res.List = append(res.List, &model.NoteSimpleInfo{
			Id:           info.Id,
			Amount:       info.Amount,
			CreatedAt:    info.CreatedAt,
			UserName:     info.User.Name,
			CategoryName: info.Category.Name,
		})
	}
	return
}
