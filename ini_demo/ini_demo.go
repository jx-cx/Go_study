package main

import (
	"errors"
	"fmt"
	"reflect"
)

// ini 文件解析器

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int64  `ini:"port"`
	Username string `int:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `redis:"host"`
	Port     int64  `redis:"port"`
	Password string `redis:"passwrod"`
	Database int8   `redis:"database"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 参数校验

	//传进来的data参数必须是指针类型（因为要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	return
}

func main() {
	var mc MysqlConfig
	err := loadIni("./conf.ini", &mc)
	if err != nil {
		fmt.Printf("load ini failed err:%s\n", err)
		return
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
