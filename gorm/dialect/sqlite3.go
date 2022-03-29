package dialect

import (
	"fmt"
	"reflect"
	"time"
)

//通过dialect实现go到sqlite3的映射


//var _ Dialect = (*sqlite3)(nil)  //这个看起来没什么用

type sqlite3 struct {

}

//服务启动之后自动注册，且为sync.Once
func init()  {
	RegisterDialect("sqlite3", &sqlite3{})
}

func (s *sqlite3) DataTypeof(t reflect.Value) string {
	switch t.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := t.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", t.Type().Name(), t.Kind()))
}

func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}
