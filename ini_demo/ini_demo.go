package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini 文件解析器

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"passwrod"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"passwrod"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
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

	//传进来的data参数必须是结构体类型指针（因为配置文件中各类键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}

	//读文件得到的字节类型数据
	buy, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(buy) 将字节类型的文件内容转换成字符串
	linSlice := strings.Split(string(buy), "\r\n")
	//fmt.Printf("%#v",linSlice)
	var structName string
	for idx, line := range linSlice {
		line = strings.TrimSpace(line) //去掉空格
		if len(line) == 0 {
			continue
		}
		//如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//如果是[开头就表示是节（section）
		if strings.HasPrefix(line, "[") {
			// 如果是[开头或者]结尾
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax err", idx+1) // idx+表示该行
				return
			}
			//当判断首位是[]之后，去掉，取出中间的内容并去掉空格
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line :%d syntax error", idx+1)
				return

			}
			// 根据字符串sectionName 去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					//找到了对接的嵌套结构体，把字段名记录下来
					structName = field.Name
					fmt.Printf("找到了%s对应的结构嵌套,结构体是%s \n", sectionName, structName)
				}
			}
		} else {
			//以等号分隔这一行
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//根据strucrName 去 data 里面把对应的结构体嵌套取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) //拿到嵌套结构体的值信息
			sType := sValue.Type()                     //拿到嵌套结构体的类型信息
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s应该是一个结构体", structName)
				return
			}
			var fileName string
			var fileType reflect.StructField
			//遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				fileType = filed
				if filed.Tag.Get("ini") == key { //找到对应的字段
					fileName = filed.Name
					break
				}

			}
			if len(fileName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fileName)
			//如果key =tag 给这个字段赋值
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("lin :%d value type error", idx+1) // 解析错误
					return
				}
				fileObj.SetInt(valueInt)
			}

		}
	}

	return
}

func main() {

	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed err:%s\n", err)
		return
	}
	fmt.Printf("%#v\r\n", cfg)
}
