package main

//1、程序自动生成一个随机数
//2、提示用户输入猜测的数据
//3、程序根据情况，给出“你猜的太小”，“你猜的太大”，“恭喜，你猜对了”。
//4、如果用户没有猜中，提示用户再次输入
//5、如果用户猜中了，程序结束
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//使用当前时间的数字，为随机数设置种子，否则会每次生成相同的数字
	rand.Seed(time.Now().UnixNano())
	//生成0到100的随即数
	number := rand.Intn(100)
	fmt.Printf("我已经想好了整数！(如果不会生成随机数，就自己写一个)\n")
	//开始循环接收用户输入，判断是否争取
	for i := 1; i < 100; i++ {
		//提示用户输入
		fmt.Print("请输入您猜的数字:")
		//接收用户输入的数据
		var inputNumber int
		fmt.Scanln(&inputNumber)
		//判断是否猜中
		if inputNumber > number {
			fmt.Printf("第%d次,您猜的太大了！\n", i)
		} else if inputNumber < number {
			fmt.Printf("第%d次,您猜的太小了！\n", i)
		} else {
			fmt.Printf("恭喜，你猜对了!\n")
			switch i {
			case 1, 2, 3:
				fmt.Printf("太厉害了，您只用%d次就猜中了！\n", i)
			case 4, 5, 6:
				fmt.Printf("不错哦，您用了%d次猜中了答案！\n", i)
			default:
				fmt.Printf("您用了%d次才猜中答案！太慢了，继续努力哦\n", i)
			}
			break
		}

	}
	fmt.Printf("\n按回车键退出程序")
	fmt.Scanln()
}
