package main

import "fmt"

type MInteger int

func (A MInteger) isBiggerThan(i MInteger) (bool) {
	if A > i {
		return true
	} else {
		return false
	}
}

func main() {
	// 后置类型声明方式
	var a int
	var b int = 1
	var c = 1
	fmt.Println(a, b, c)

	// 多个变量同时定义
	var (
		d, e int
		f float64
	)
	fmt.Println(d, e, f)

	// 简短声明方式
	g := 1
	h := int64(1)
	fmt.Println(g, h)

	// 常量
	const i = 32
	const (
		j = 4
		k = 0.1
	)
	fmt.Println(i, j, k)

	// 字符串
	var l = "string"
	var m = "\""

	fmt.Println(l)
	fmt.Println(l[0])
	fmt.Println(len(m))

	for _, char := range l {
		fmt.Println("%T", char)
	}

	// 数组
	var arr [5]int
	for i:=0;i<5;i++ {
		arr[i] = i
	}
	printHelper("arr:", arr)
	
	n := [...]string{"Hello", "World", "Tomorrow", "Is", "Better"}
	o := n[:]
	fmt.Println(o)

	// 自定义类型
	var mint MInteger = 10
	fmt.Println(mint.isBiggerThan(9))
	fmt.Println(mint.isBiggerThan(10))
	fmt.Println(mint.isBiggerThan(11))
}

func printHelper(name string, arr [5]int) {
    for i := 0; i < 5; i++ {
        fmt.Printf("%v[%v]: %v\n", name, i, arr[i])
    }

    // len 获取长度
    fmt.Printf("len of %v: %v\n", name, len(arr))

    // cap 也可以用来获取数组长度
    fmt.Printf("cap of %v: %v\n", name, cap(arr))

    fmt.Println()
}