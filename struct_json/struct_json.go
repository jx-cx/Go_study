package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name" ini:"name"` // 自定义在别的包引用时候的名称
	Age  int    `json:"age"`
}

func main() {
	a := person{
		"小红",
		18,
	}
	//序列化
	p, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%#v", string(p))
	fmt.Println()
	// 反序列化
	str := `{"name":"小明","age":18}`
	var p1 person
	json.Unmarshal([]byte(str), &p1) // 传指针是为了能在json.Unmarshal内部修改p1的值
	fmt.Printf("%#v", p1)
}
