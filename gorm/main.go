package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

)

//注意，gorm只是一个orm框架，只是提供一些sdk或者函数，并不需要起服务，所以不需要main函数
func main() {
	engine, _ := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()   //获取执行的条数
	fmt.Printf("Exec success, %d affected\n", count)
}
