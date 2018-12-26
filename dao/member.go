package dao

func NewMemberDao() *Tbl {
	m := &Tbl{
		Name: "member",   //表名
		Key: "id",        //主键字段名
		//Cols: new(map[string]interface{}),
	}

	return m;
}