package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...interface{}) (string, []interface{}) //输入golang语句，输出sql语句和占位符的值

var generators map[Type]generator  //输入golang对应插入和查询名字,对应sql语句


func init()  {
	generators = make(map[Type]generator)
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
	generators[LIMIT] = _limit
}

func genBindVars(num int) string {
	var vars []string
	for i:=0;i<num;i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ",")
}

//例如输入user, (name, age)， 输出INSERT INTO user (name, age)
func _insert(values ...interface{}) (string, []interface{})  {
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

//VALUES (?), (?), (v1), (v2)
func _values(values ...interface{}) (string, []interface{})  {
	//fields := strings.Join(values[0].([]string), ",")
	//return fmt.Sprintf("VALUES (%v)", fields), []interface{}{}
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, value := range values {
		v := value.([]interface{}) // 我认为这里的目的是能计算len,并且匹配返回值
		if bindStr == "" {
			bindStr = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindStr))
		if i+1 != len(values) {
			sql.WriteString(", ")
		}
		vars = append(vars, v...)
	}
	return sql.String(), vars
}

//SELECT $fields FROM $rableName
func _select(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	fileds := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fileds, tableName), []interface{}{}
}

//LIMIT $num  ,limit 可以是num，也可以是num,num,todo,这里只给出了limit num的形式
func _limit(values ...interface{}) (string, []interface{}) {
	return "LIMIT ?", values
}

// WHERE $desc,这里一个where支持一个条件，比如where a=1, desc是a=?, vars是1
func _where(values ...interface{}) (string, []interface{}) {
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

// id desc或者id asc
func _orderBy(values ...interface{}) (string, []interface{})  {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

