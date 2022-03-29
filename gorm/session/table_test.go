package session

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gorm/dialect"
	"testing"
)

type User struct {
	Name string `gorm:"PRIMARY KEY"`
	Age int
}

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestSession_CreateTable(t *testing.T) {
	TestDB, _ = sql.Open("sqlite3", "./gee.db")   //此处报错的原因是没有get
	                                                                     // _ "github.com/mattn/go-sqlite3"数据库，get之后就不报错了
	s := NewSession()
	s.Model(&User{})
	_ = s.DropTable()
	_ = s.CreateTable()
	if !s.HasTable() {
		t.Fatal("Failed to create User")
	}
}
