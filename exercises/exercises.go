package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**

 */

var aa string = "how do you do"
var (
	bb    string = "hello小伙子"
	cc    int
	dd    = "上海自来水来自海上"
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

/*
   你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
   分配规则如下：
   a. 名字中每包含1个'e'或'E'分1枚金币
   b. 名字中每包含1个'i'或'I'分2枚金币
   c. 名字中每包含1个'o'或'O'分3枚金币
   d: 名字中每包含1个'u'或'U'分4枚金币
   写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
   程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
func dispatchCoin() (left int) {
	for _, name := range users {
		for _, a := range name {
			switch a {
			case 'e', 'E':
				distribution[name]++
				coins -= 1
			case 'i', 'I':
				distribution[name]++
				coins -= 2
			case 'o', 'O':
				distribution[name]++
				coins -= 3
			case 'u', 'U':
				distribution[name]++
				coins -= 4
			}
		}
	}
	return coins
}

func main() {
	//练习题 1
	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
	a1 := strings.Split(aa, " ") //把aa按照空格切割成切片a1
	//创建一个空数组，把s1遍历加到数组中
	a2 := make(map[string]int, 10)
	for _, a := range a1 {
		if _, ok := a2[a]; !ok { // 如果a1[a]不存在于a2[a]中，则a2[a]=1 ,存在则累积次数
			a2[a] = 1
		} else {
			a2[a]++
		}
	}
	//遍历数组a2
	for count, value := range a2 {
		fmt.Println(count, value)
	}

	// 练习题2
	//观察以下代码，写出打印结果
	//	type Map map[string][]int
	//	m := make(Map)
	//	s := []int{1, 2}
	//	s = append(s, 3)
	//	fmt.Printf("%+v\n", s)  s=1 2 3
	//	m["q1mi"] = s
	//	s = append(s[:1], s[2:]...)  s=1 3   切掉了2
	//	fmt.Printf("%+v\n", s)   s=1 3
	//	fmt.Printf("%+v\n", m["q1mi"])  底层数据改变，所以 m["q1mi"] =1 3 3

	// 练习题3 判断字符串中汉子的数量
	for _, h := range bb {
		if unicode.Is(unicode.Han, h) {
			cc++
		}
	}
	fmt.Println(cc)
	// 练习4  回文判断
	d := make([]rune, 0, len(dd))
	for _, a := range dd {
		d = append(d, a)
	}

	for i := 0; i < len(d)/2; i++ {
		if d[i] != d[len(d)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for i, v := range distribution {
		fmt.Printf("%s %d\n", i, v)
	}
}
