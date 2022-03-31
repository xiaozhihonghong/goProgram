package clause

import "strings"

type Type int

type Clause struct {
	sql map[Type]string  //golang函数到sql语句映射
	SqlVars map[Type][]interface{}   //golang函数到数据映射
}

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)    //todo 暂时不知道这里为什么这么写


//func NewClause() *Clause {
//	return &Clause{
//		sql: make(map[Type]string),
//		SqlVars: make(map[Type][]interface{}),
//	}
//}

// 构建sql语句个变量
func (c *Clause) Set(name Type, vars ...interface{})  {
	if c.sql == nil {
		c.sql = make(map[Type]string)
		c.SqlVars = make(map[Type][]interface{})
	}
	sql, vars := generators[name](vars...)
	c.sql[name] = sql
	c.SqlVars[name] = vars
}

//将所有构建的语句连接起来
func (c *Clause) Build(orders ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, order := range orders {
		if sql, ok := c.sql[order];ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.SqlVars[order]...)
		}
	}
	return strings.Join(sqls, " "), vars
}

