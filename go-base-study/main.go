package main

import (
	"errors"
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

// 函数
func fun01(x int, y int) int {
	return x + y
}
func fun02(x, y int) int {
	return x * y
}
func fun03() {
	type calculation func(int, int) int
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))
}
func fun04(x, y int) (string, int) {
	return fmt.Sprintf("%d", x), y
}
func fun05(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func fun06(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
func fun07() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //return x 的含义是将当前 x 的值（即 5）作为返回值保存。此时返回值已经确定为 5。
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //在返回前，还会进行赋值x，函数有一个 命名返回值 x，其初始值为 0，return 5 表示将 5 赋值给命名返回值 x，此时 x = 5。执行 x++ 后，x 的值从 5 增加到 6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //x已经赋值给y
}

func f4() (x int) {
	defer func(x int) {
		x++ //修改的是传递给匿名函数的副本 x，而非外层的命名返回值 x
	}(x) // 此时传递的 x 是命名返回值 x 的当前值，即初始值 0
	return 5
}

func main() {
	//arrayDemo()
	//sliceDemo()
	//mapDemo4()
	//mapDemo5()
	//fun03()
	//s, i := fun04(1, 2)
	//fmt.Printf("type:%T value:%v\n", s, s)
	//fmt.Println(s)
	//fmt.Println(i)
	//println(fun05(1, 2, add))
	//f, err := fun06("-")
	//if err != nil {
	//	println(err.Error())
	//}
	//println(f(1, 2))
	//fun07()
	fmt.Println(f1())
	fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())
}
