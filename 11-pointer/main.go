package main

import (
	"fmt"
	"log"
)

func main() {
	v := "hello"
	ptr2 := &v        // 取得變數 v的地址
	log.Println(ptr2) // ptr的型態是 *string

	//----

	var house = "Malibu Point 10880, 90265"
	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr) // ptr type: *string
	fmt.Printf("address: %p\n", ptr)  // address: 0xc0000961f0

	value := *ptr
	fmt.Printf("value type: %T\n", value) // value type: string
	fmt.Printf("value: %s\n", value)      // value: Malibu Point 10880, 90265

	// 	取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
	//
	// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
	//  - 对变量进行取地址操作使用&操作符，可以获得这个变量的指针变量。
	//  - 指针变量的值是指针地址。
	//  - 对指针变量进行取值操作使用*操作符，可以获得指针变量指向的原变量的值。
	//  - Go語言中除了map、slice、chan外，其他類型在函數參數中都是值傳遞
	//  - Go語言不是面向對象的語言，很多時候實現結構體方法時需要用指針類型實現引用結構體對象
	//  - 指針也是一個類型，在實現接口interface時，結構體類型和其指針類型對接口的實現是不同的

	//----

	f := func(a int, b *int) { // b是 int指針型別
		fmt.Println(a, b) // 1 0xc00001a148
		a++
		*b = *b + 1 // 要拿值就需要 *來取得該記憶體位置的值
	}
	a := 1
	b := 1
	f(a, &b)
	fmt.Println(a, b)

	//-----

	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)

	//-----

	var aa int
	var ptraa *int
	var pptr **int // **代表指向指針的指針 , 当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：

	aa = 3000

	/* 指针 ptr 地址 */
	ptraa = &aa

	/* 指向指针 ptr 地址 */
	pptr = &ptraa

	aa = 1

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", aa)                 // 1
	fmt.Printf("指针变量 *ptr = %d\n", *ptraa)        // 1
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr) // 1
}
