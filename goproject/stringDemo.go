package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.IndexRune("小韩说课", '韩'))


	// Split() 字符串分割
	slice2 := strings.Split("hello-world-haha","-")  // 返回字符串切片 []string
	fmt.Println(slice2)  // [ hello world haha ]  (两端有两个空元素)
	slice3:= []byte(slice2[0])
	fmt.Printf("%c\n",slice3)
	fmt.Println(slice2[1])
	fmt.Println(len(slice2))  // 5

	fmt.Println(strings.Count("Ha nZhong Kanang", "an"))
	// 返回 3


	str := "Go is an open source prograiting language that makes it easy to build simple, reliable, and efficient it software."
	// 存储单词词频
	count := make(map[string]int)

	// 分词
	substr := strings.Split(str, " ")
	//统计词频
	for _, value := range substr {
		count[value]++
	}

	for key, value := range count {
		fmt.Printf("%s: %d \n", key, value)
	}

}
