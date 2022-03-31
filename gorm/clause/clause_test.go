package clause

import (
	"reflect"
	"testing"
)

func testInsert(t *testing.T)  {
	var clause Clause
	clause.Set(INSERT, "User", []string{"Name", "Age"})
	v := make([]interface{}, 0)
	v = append(v, []interface{}{"Chenzhi", 6})
	v = append(v, []interface{}{"Tom", 3})
	clause.Set(VALUES, v...)    //这里注意输入是v...
	sql, vars := clause.Build(INSERT, VALUES)
	t.Log(sql, vars)
	if sql != "INSERT INTO User (Name,Age) VALUES (?,?), (?,?)" {
		t.Fatal("failed to build SQL")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Chenzhi", 6, "Tom", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}

func testSelect(t *testing.T) {
	var clause Clause
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(ORDERBY, "Age ASC")
	clause.Set(LIMIT, 3)
	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)
	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Age ASC LIMIT ?" {
		t.Fatal("failed to build SQL")
	}
	if !reflect.DeepEqual(vars, []interface{}{"Tom", 3}) {
		t.Fatal("failed to build SQLVars")
	}
}

//并行测试
func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
	t.Run("insert", func(t *testing.T) {
		testInsert(t)
	})
}
