package session

import (
	"database/sql"
	"gorm/clause"
	"gorm/dialect"
	"gorm/log"
	"gorm/schema"
	"strings"
)

// session用来和数据库交互

type Session struct {
	db    *sql.DB  //sql原生数据库
	tx    *sql.Tx   //使用事务操作
	dialect dialect.Dialect   //解析的约定
	refTable *schema.Schema  //保存解析后的表
	sql   strings.Builder  //sql语句，sql的关键字
	sqlVars []interface{}  //占位符，方sql注入
	clause  clause.Clause
}

func New(db *sql.DB, d dialect.Dialect) *Session {
	return &Session{
		db: db,
		dialect: d,
	}
}

type CommonDB interface {
	Query(query string, args ...interface{})(*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}
var _ CommonDB = (*sql.DB)(nil)  //就是实例化一个结构体
var _ CommonDB = (*sql.Tx)(nil)

// 就是将语句和占位符都清空
func (s *Session) Clear()  {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

//这个函数本质上没什么用，直接使用s.db也是一样的，可能是为了封装的更彻底一点
func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

//exec封装两个目的，一个是统一打印日志，二是清空变量，复用session，一次会话多次执行sql
func (s *Session) Exec() (sql.Result, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err := s.DB().Exec(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return rows, nil
}

//查询返回一行
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

//select的所有返回
func (s *Session) QueryRows() (*sql.Rows, error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err := s.DB().Query(s.sql.String(), s.sqlVars...)
	if err != nil {
		log.Error(err)
	}
	return rows, nil
}
