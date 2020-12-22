package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	//1、定义dsn字符串，指定准备使用数据库的
	var dsn string="root:123@tcp(127.0.0.1:3306)/student_mis"

	//2、定义DB对象，该对象负责操作数据库，该对象是个指针变量
	var db *sql.DB
	var err error
	//3、创建DB对象，并指定参数，第一个参数是数据库驱动程序名称，第二个参数为数据库服务器的dsn
	db,err=sql.Open("mysql",dsn)
	//判断为DB指定参数时，是否有错误产生
	if err != nil {
		panic(err)//panic方法，直接终止整个程序的执行，效果相当于return，同时在屏幕上打印 参数
	}
	//4、Ping（）方法指挥DB对象与数据库进行连接，如果成功，就可以通过DB对象对数据库进行操作了
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功!!!")//思考：该输出语句并未写在else中，为何连接错误时不被执行

	//读取数据
	rows,err:=db.Query("select * from sys_user ")
	if err!=nil{
		panic(err)
	}
	var user_id int
	var user_name,user_password string
	for rows.Next(){
		rows.Scan(&user_id,&user_name,&user_password)
		fmt.Println(user_id,user_name,user_password)
	}

	//打印菜单
	fmt.Println("")
	for true{

	}

}

