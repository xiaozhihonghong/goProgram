package session

import (
	"errors"
	"gorm/clause"
	"reflect"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).GetRefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)   //有判断是否为nil，不会重复写入
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

//todo，反射还需要后续好好学习，东西很多
func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).GetRefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()   //这里是查询出符合条件的数据
	if err != nil {
		return err
	}

	//这里将查询出的数据写入user中
	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {  //将values写入
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

func (s *Session) Update(kv ...interface{}) (int64, error) {
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i:=0;i<len(kv);i+=2 {
			m[kv[i].(string)] = kv[i+1]
		}
	}
	s.clause.Set(clause.UPDATE, s.GetRefTable().Name, m)
	sql, vars := s.clause.Build(clause.UPDATE, clause.WHERE)  //完整的update语句
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()   //更新的条数
}

//返回删除成功的条数
func (s *Session) Delete() (int64, error) {
	s.clause.Set(clause.DELETE, s.GetRefTable().Name)
	sql, vars := s.clause.Build(clause.DELETE, clause.WHERE)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Count() (int64, error) {
	s.clause.Set(clause.COUNT, s.GetRefTable().Name)
	sql, vars := s.clause.Build(clause.COUNT, clause.WHERE)
	row := s.Raw(sql, vars...).QueryRow()
	var tmp int64
	if err := row.Scan(&tmp); err != nil {
		return 0, err
	}
	return tmp, nil
}

// 链式调用
func (s *Session) Limit(num int) *Session {
	s.clause.Set(clause.LIMIT, num)
	return s
}

// 链式调用
func (s *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	s.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return s
}

//查询一条记录
func (s *Session) First(value interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(value))  //找出对象
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()   //实例化对象
	if err := s.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("Not found")
	}
	dest.Set(destSlice.Index(0))
	return nil
}

func (s *Session) OrderBy(desc string) *Session {
	s.clause.Set(clause.ORDERBY, desc)
	return s
}
