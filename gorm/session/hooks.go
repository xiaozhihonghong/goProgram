package session

import (
	"gorm/log"
	"reflect"
)

//钩子的本质是设置回调函数，通过反射获取，然后在具体场景下实现自己需要的函数功能

const (
	BeforeQuery = "BeforeQuery"
	AfterQuery = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert = "AfterInsert"
)

//反射实现钩子的回调函数
func (s *Session) CallMethod(method string, value interface{})  {
	//首先获取会话的对象
	//反射获取对象的方法
	fm := reflect.ValueOf(s.GetRefTable().Model).MethodByName(method)
	//也可以通过value获取,value就是输入的一个对象
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method)
	}
	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {   //使用fm.Call调用函数
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
}
