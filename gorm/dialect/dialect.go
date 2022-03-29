package dialect

import "reflect"


var dialectMap  = map[string]Dialect{}   //key为那种数据库，Dialect为该数据库的转换
type Dialect interface {
	DataTypeof(t reflect.Value) string   //将go语言类型转换成数据类型
	TableExistSQL(tableName string) (string, []interface{})   //返回某个表是否存在的sql语句，参数为表名，返回值是语句和占位符
}

//注册dialect实例
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

//获取dialect实例
func GetDialect(name string) (Dialect, bool)  {
	dialect, ok := dialectMap[name]
	return dialect, ok
}
