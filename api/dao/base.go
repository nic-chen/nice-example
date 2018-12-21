package dao

import (
	"strings"
	"database/sql"
	"github.com/nic-chen/nice"
	"../../config"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)

type Columns struct {
}

type Tbl struct {
	Key string              //主键
	Name string             // 表名
	Cols map[string]string  // 表的所有列信息
}

func (d *Tbl) Fetch(value interface{}) (map[string]interface{}, error) {

	cache := nice.Instance(config.APP_NAME).Cache();
	if cache!=nil {
	}
	
	db  := nice.Instance(config.APP_NAME).Db();
	tbl := d.Name;
	col := d.Key;
	val := convertToString(value);
	sql := ` SELECT * FROM ` + tbl + ` WHERE ` + "`" + col + "`=?";
	tmp := make(map[string]interface{});
	res, err := db.Query(sql, val) // 执行语句并返回

	if err!=nil{
		return tmp, err;
	}

	if len(res)>0 {
		return res[0], err
	}

	return tmp, err;
}

func (d *Tbl) Insert(data map[string]interface{}, replace bool) (sql.Result, error) {
	db := nice.Instance(config.APP_NAME).Db(); //.(*Mysql)
	tbl := d.Name;

	cmd := ` INSERT INTO `;
	if(replace){
		cmd =  ` REPLACE INTO `
	}

	kv := d.implode(data);

	sql := cmd+tbl+" SET "+kv;

	res, err := db.Exec(sql) // 执行语句并返回

	return res, err
}

func (d *Tbl) implode(data map[string]interface{}) string{
	sql := ``
	for key, value := range data { // 遍历要传入的数据
		// 拼接set后的字符串  a=1,b='2',c=11
		if attr, ok := d.Cols[key]; ok { // 如果表中存在这个字段
			// 存在这个字段
			if attr!="" {
				sql = sql + "`" + key + "`" + `='` + convertToString(value) + `',`;
			}
		}
	}
	sql = strings.TrimRight(sql, `,`) // 去掉最后的逗号

	return sql;
}


// func (d *Tbl) Fetch_cache(value interface{}) (map[string]interface{}, error){

// }

// func (d *Tbl) Store_cache(value interface{}, data map[string]interface{}) error{

// }

// func (d *Tbl) Delete_cache(value interface{}) error{

// }


// 把数据转换为字符串
func convertToString(m interface{}) string {
	switch m.(type) {
	case int64:
		return strconv.FormatInt(m.(int64),10)
	case int32:
		return strconv.Itoa(int(m.(int32)))
	case int:
		return strconv.Itoa(m.(int))
		break
	default:
		return m.(string)
	}
	return ""
}
