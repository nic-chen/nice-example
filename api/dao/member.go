package dao

type Member struct {
	Id int              //主键
	Nickname string            
	Avatar string            
	Columns
}

func NewMemberDao() *Tbl {
	cols := make(map[string]string);
	m := &Tbl{
		Name: "member",
		Key: "id",
		Cols: cols,
	}

	return m;
}