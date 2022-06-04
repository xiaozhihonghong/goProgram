package schema

import (
	"go/ast"
	"gorm/dialect"
	"reflect"
)

//这个结构体是创建表中的字段，类型和额外约束条件（主键等）
type Field struct {
	Name string
	Type string
	Tag string
}

//这个结构体是整个表的映射
type Schema struct {
	Model interface{}  //传入表的struct实例
	Name string        //表名
	Fields []*Field   //所有字段
	FieldNames []string  //所有字段的名字
	FieldMap map[string]*Field  //名字到字段的映射
}

func (s *Schema) GetField(name string) *Field {
	return s.FieldMap[name]
}

//进行解析,dest为输入的结构体，d为解析的规则，返回解析后的sql语句
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	//判断是否为一个指针，如果是指针就通过 Elem() 获取指针指向的变量值
	modeType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model: dest,
		Name: modeType.Name(),
		FieldMap: make(map[string]*Field),
	}

	for i:=0;i<modeType.NumField();i++ {
		p := modeType.Field(i)
		//忽略匿名字段和私有字段
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeof(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("gorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.FieldMap[p.Name] = field
		}
	}
	return schema
}

func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var filedValues []interface{}
	for _, field := range s.Fields {
		filedValues = append(filedValues, destValue.FieldByName(field.Name).Interface())
	}
	return filedValues
}