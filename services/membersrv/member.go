package membersrv

import (
	"context"
	"../../api/dao"
	"../../proto/member"
)

type memberService struct {
}

func NewMemberService() *memberService {
	return &memberService{}
}

func (*memberService) Info(ctx context.Context, req *member.Request) (*member.Response, error) {
	d := dao.NewMemberDao();
	m, _ := d.Fetch(req.Id);

	return &member.Response{Id: m["id"].(int64), Nickname: m["nickname"].(string), Avatar: m["avatar"].(string)}, nil
}
