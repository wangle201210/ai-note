// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package note

import (
	"context"

	"github.com/wangle201210/ai-note/app/ai-note/api/note/v1"
)

type INoteV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	AddNote(ctx context.Context, req *v1.AddNoteReq) (res *v1.AddNoteRes, err error)
	GetNoteList(ctx context.Context, req *v1.GetNoteListReq) (res *v1.GetNoteListRes, err error)
	GetNoteInfo(ctx context.Context, req *v1.GetNoteInfoReq) (res *v1.GetNoteInfoRes, err error)
}
