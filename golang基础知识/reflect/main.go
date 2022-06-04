package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}
/*
转换成
type Config struct {
	Name    string `json:"server-name"` // CONFIG_SERVER_NAME
	IP      string `json:"server-ip"`   // CONFIG_SERVER_IP
	URL     string `json:"server-url"`  // CONFIG_SERVER_URL
	Timeout string `json:"timeout"`     // CONFIG_TIMEOUT
}
 */

func readConfig() *Config {
	//反射值有两个，一个type和一个value，分别使用Typeof和Valueof获取
	config := Config{}
	t := reflect.TypeOf(config)    //接口的类型描述信息，比如“int”和“string”
	//v := reflect.ValueOf(config)   //返回接口的值类型
	//fmt.Println(v.Kind())   //返回一个类型，这里是个指针
	//fmt.Println(v.Elem())  //可以通过 reflect.Elem() 方法获取这个指针指向的元素类型
	value := reflect.Indirect(reflect.ValueOf(&config)) //reflect.Indirect()函数用于获取v指向的值，即，如果v是nil指针，
	// 则Indirect返回零值。如果v不是指针，则Indirect返回v。

	//t.NumField()用于获取结构t中的字段数
	for i:=0;i<t.NumField();i++ {
		f := t.Field(i)   // 获取反射字段
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))   //根据字段获取值，然后set写入值
			}
		}
	}
	return &config
}

func main() {
	//os.Setenv("CONFIG_SERVER_NAME", "global_server")
	//os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	//os.Setenv("CONFIG_SERVER_URL", "geektutu.com")
	//c := readConfig()
	//fmt.Printf("%+v", c)
	a := 8
	v := reflect.ValueOf(a)
	fmt.Println("v=", v)
	k := v.Kind()
	fmt.Println(k)
}
