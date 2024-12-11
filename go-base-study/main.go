package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 数组
func arrayDemo() {
	var numArray = [...]int{1, 2, 5: 3}
	strArray := [...]string{0: "han", 1: "peng"}

	for i := 0; i < len(numArray); i++ {
		fmt.Println(numArray[i])
	}

	for index, value := range strArray {
		fmt.Println(index, value)
	}

	a := [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i, val := range a {
		fmt.Println(i, val)
	}
}

// 切片
func sliceDemo() {
	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3] // s := a[low:high]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	//fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	var a = [...]int{3, 7, 8, 9, 1}
	s := a[:]
	sort.Ints(s)

	for i, v := range s {
		fmt.Println(i, v)
	}
}

// map
func mapDemo1() {
	m := make(map[string]int)
	m["han"] = 10
	m["zhi"] = 20
	m["peng"] = 30
	fmt.Println(m)

	// 判断是否存在
	val, ok := m["zh"]
	fmt.Println(val, ok)
}
func mapDemo2() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
func mapDemo3() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
func mapDemo4() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
func mapDemo5() {
	str := "how do you do"
	splitArray := strings.Split(str, " ")
	m := make(map[string]int, cap(splitArray))

	for _, v := range splitArray {
		val, ok := m[v]
		if ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}

	for key, value := range m {
		fmt.Printf("key:%v value:%v\n", key, value)
	}
}

func main() {
	//arrayDemo()
	//sliceDemo()
	//mapDemo4()
	mapDemo5()
}
