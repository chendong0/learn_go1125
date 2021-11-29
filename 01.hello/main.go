package main

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello, golang!"
	var world = "world"
	fmt.Println(hello)

	fmt.Println(hello, world)

	var name string = "小强"
	var name1 = &name
	fmt.Println(name)
	fmt.Printf("name变量的类型是 %T\n,值是%v\n:", &name, &name)
	fmt.Printf("name1变量的类型是 %p\n,值是%p\n:", *name1, *name1)
	//不知道如何print出name的指针变量.

	var age int = 35
	fmt.Println(age)
	var tall float64 = 1.80
	fmt.Println(tall)

	//name = "小强测试变量" //变量不能重名?那么这个name是不是重名呢?
	//回答,name是重新赋值,指针地址是一致的.
	fmt.Println(name)
	//age = "测试一下类型" //类型不兼容报错
	//cannot use "测试一下类型" (type untyped string) as type int in assignment
	//tall = age //类型不兼容报错
	//cannot use "测试一下类型" (type untyped string) as type int in assignment

	var heng int = 3
	var shu int = 4
	fmt.Println(heng * shu)
	var mianji = heng * shu
	fmt.Println(mianji)

	var f1 float64 = 1.234
	var i1 int = int(f1)
	fmt.Println("f1:", f1, "i1:", i1)
	var a6 uint = math.MaxUint64
	var a7 int = int(a6)
	fmt.Println(a6, a7) //变量类型转换及其代价
	//补充知识,原码 反码 补码的知识

}
