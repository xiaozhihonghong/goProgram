package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 注，原始的"github.com/go-sql-driver/mysql"和"database/sql"不提供创建表的操作，需要自己在mysql中建表
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:chenzhi@tcp(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func Insert() {
	result,err:=db.Exec("INSERT INTO test(idtest, name, money, birthday)" +
		"VALUES (?, ?, ?, ?)",4, "hong", "10", nil)
	if err!=nil {
		fmt.Println(err)
		return
	}
	//获取修改的最后一个id，但是每次都是0，没报错，暂时不清楚原因
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
}


func Update() {
	result, err := db.Exec("UPDATE test set money=money+10 WHERE idtest>=?", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取修改的行数
	rows,err := result.RowsAffected()
	if err != nil {
		fmt.Println(err); return
	}
	fmt.Println(rows)
}


func Select() {
	var id int64
	var Username string
	var Money float64
	var Birthday sql.NullString
	//参数绑定，可以避免sql注入
	rows := db.QueryRow("SELECT idtest,name,money,birthday FROM test WHERE idtest>?", 1) //查找一行
	rows2, _ := db.Query("SELECT idtest,name,money,birthday FROM test WHERE idtest>?", 1) //查找所有行
	err := rows.Scan(&id, &Username, &Money, &Birthday)
	if err != nil {
		fmt.Println("scan err:", err)
		return
	}
	fmt.Println(id, Username, Money, Birthday.String)

	for rows2.Next() {
		err := rows2.Scan(&id, &Username, &Money, &Birthday)
		if err != nil {
			fmt.Println("scan err:", err)
			return
		}
		fmt.Println(id, Username, Money, Birthday.String)
	}
}


func Delete() {
	result, err := db.Exec("DELETE FROM test WHERE idtest=?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取删除的行数
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rows)
}

func main() {
	//Insert()
	//Update()
	//Select()
	Delete()
}

