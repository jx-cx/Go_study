package main

import (
	"fmt"
	"strconv"
)

func main() {
	//字符串转换成数字
	str := "100"
	ret1, _ := strconv.ParseInt(str, 10, 0) //10是进制  64 是int64
	fmt.Printf("%#v %T \n", ret1, ret1)

	retint, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T \n", retint, retint)

	//把数字转换成字符串
	str1 := int32(100)
	str2 := fmt.Sprintf("%d", str1)
	fmt.Printf("%#v %T \n", str2, str2)
	i := 100
	str3 := strconv.Itoa(i)
	fmt.Printf("%#v %T \n", str3, str3)

}
