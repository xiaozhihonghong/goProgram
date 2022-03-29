package engine

import (
	"database/sql"
	"gorm/dialect"
	"gorm/log"
	"gorm/session"
)

//处理连接数据库和断开数据库的工作
type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}


//数据库的连接
func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return nil, err
	}
	e := &Engine{db:db, dialect: dial}
	log.Info("连接成功")
	return e, nil
}

//数据库的关闭
func (e *Engine) Close()  {
	if err := e.db.Close(); err != nil{
		log.Error("关闭数据库失败")
	}
	log.Info("数据库关闭成功")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}