package note

import (
	"context"
	"fmt"
	"github.com/wangle201210/ai-note/app/ai-note/internal/dao"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/do"
	"github.com/wangle201210/ai-note/app/ai-note/internal/model/entity"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

func (c *ControllerV1) AddNote(ctx context.Context, req *v1.AddNoteReq) (res *v1.AddNoteRes, err error) {
	user := &entity.User{}
	category := &entity.Category{}
	fmt.Printf("req.UserName: %s, req.CategoryName: %s\n", req.UserName, req.CategoryName)
	dao.User.Ctx(ctx).Fields(dao.User.Columns().Id).Where(do.User{Name: req.UserName}).Scan(&user)
	dao.Category.Ctx(ctx).Fields(dao.Category.Columns().Id).Where(do.Category{Name: req.CategoryName}).Scan(&category)
	insertId, err := dao.Note.Ctx(ctx).Data(do.Note{
		UserId:     user.Id,
		CategoryId: category.Id,
		Amount:     req.Amount,
		OrgContent: req.OrgContent,
		Status:     v1.StatusOK,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	res = &v1.AddNoteRes{
		Id: insertId,
	}
	return
}
