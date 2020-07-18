package main

import "log"

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "AAA", Age: 10},
		{Name: "BBB", Age: 20},
		{Name: "CCC", Age: 30},
	}

	log.Printf("%p", &(stus[0])) // 0xc0000be050
	log.Printf("%p", &(stus[1])) // 0xc0000be068
	log.Printf("%p", &(stus[2])) // 0xc0000be080

	for _, stu := range stus { // 都是 0xc0000ac040
		log.Printf("%p", &stu)
	}

	for _, stu := range stus { // 都是 0xc0000ac060
		log.Printf("%p", &stu)
	}

	/* bug
	output:
	2020/07/09 18:09:34 {CCC 30}
	2020/07/09 18:09:34 {CCC 30}
	2020/07/09 18:09:34 {CCC 30}
	*/
	for _, stu := range stus {
		/*
			因为 for range在遍历值类型时，其中的 stu变量是一个值的拷贝，
			当使用&获取指针时，实际上是获取到 stu这个临时变量的指针，
			而 stu变量在for range中只会创建一次，之后循环中会被一直重复使用，
			所以在 m[stu.Name]赋值的时候其实都是 stu变量的指针，而 &stu最终会指向 stus最后一个元素的值拷贝。
		*/
		m[stu.Name] = &stu
	}

	/* 解法
	output:
	2020/07/09 18:20:30 {AAA 10}
	2020/07/09 18:20:30 {BBB 20}
	2020/07/09 18:20:30 {CCC 30}
	*/
	// for _, stu := range stus {
	// 	tmp := stu
	// 	m[stu.Name] = &tmp
	// }

	/* 解法
	output:
	2020/07/09 18:20:30 {AAA 10}
	2020/07/09 18:20:30 {BBB 20}
	2020/07/09 18:20:30 {CCC 30}
	*/
	// for i, _ := range stus {
	// 	m[stus[i].Name] = &(stus[i])
	// }

	/* 解法
	output:
	2020/07/09 18:20:30 {AAA 10}
	2020/07/09 18:20:30 {BBB 20}
	2020/07/09 18:20:30 {CCC 30}
	*/
	// for i := 0; i < len(stus); i++ {
	// 	m[stus[i].Name] = &(stus[i])
	// }

	for _, v := range m {
		log.Println(*v)
	}
}
