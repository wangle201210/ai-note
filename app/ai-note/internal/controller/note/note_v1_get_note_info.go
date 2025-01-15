package note

import (
	"context"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

func (c *ControllerV1) GetNoteInfo(ctx context.Context, req *v1.GetNoteInfoReq) (res *v1.GetNoteInfoRes, err error) {
	res = &v1.GetNoteInfoRes{}
	user := &entity.User{}
	var list []*model.NoteInfo

	dao.User.Ctx(ctx).Fields(dao.User.Columns().Id).Where(do.User{Name: req.UserName}).Scan(&user)

	sql := dao.Note.Ctx(ctx).Where(dao.Note.Columns().UserId, user.Id).
		WhereBetween(dao.Note.Columns().CreatedAt, req.Start, req.End)
	if req.CategoryName != model.AllCategory {
		cate := &entity.Category{}
		dao.Category.Ctx(ctx).Fields(dao.Category.Columns().Id).Where(do.Category{Name: req.CategoryName}).Scan(&cate)
		sql = sql.Where(dao.Note.Columns().CategoryId, cate.Id)
	}
	res.Amount, _ = sql.Sum(dao.Note.Columns().Amount)
	sql.WithAll().Scan(&list)
	m := map[string]float64{}
	for _, info := range list {
		m[info.Category.Name] += info.Amount
		// res.List = append(res.List, &model.NoteSimpleInfo{
		// 	// Id:     info.Id,
		// 	Amount: info.Amount,
		// 	// CreatedAt:    info.CreatedAt,
		// 	UserName:     info.User.Name,
		// 	CategoryName: info.Category.Name,
		// })
	}
	for k, v := range m {
		res.List = append(res.List, &model.NoteSimpleInfo{
			Amount:       v,
			CategoryName: k,
		})
	}
	return
}
