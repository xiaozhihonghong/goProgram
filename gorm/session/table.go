package session

import (
	"fmt"
	"gorm/log"
	"gorm/schema"
	"reflect"
	"strings"
)

//初始化解析函数
func (s *Session) Model(value interface{}) *Session {
	//表不存在或两个表的类型不一致
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

// 调用refTable
func (s *Session) GetRefTable() *schema.Schema  {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}

//建表，其实就是执行解析后的sql语句
func (s *Session) CreateTable() error {
	table := s.GetRefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

//删表
func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", s.GetRefTable().Name)).Exec()
	return err
}

//判断表是否存在,从数据库中查找是否有这个表
func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistSQL(s.GetRefTable().Name)
	row := s.Raw(sql, values...).QueryRow() //从数据库中查找这个表
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.GetRefTable().Name
}