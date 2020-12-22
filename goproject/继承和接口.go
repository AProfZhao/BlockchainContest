package main

import "fmt"

/*
1、go中没有“继承”，只是用 结构体的  组合  来模拟继承
 */
//定义结构体  动物
type animal struct {
	name string
	weight int
}
//定义结构体 猫，它是动物的“子类
type cat struct {
	animal
	name string
}

func main()  {
	//结构体cat没有定义name字段，但已然可以使用name字段
	myCar:=new(cat)
	myCar.animal.name="波斯猫"
	myCar.name="老虎"

	fmt.Println(&myCar.animal.name)
	fmt.Println(&myCar.name)

}