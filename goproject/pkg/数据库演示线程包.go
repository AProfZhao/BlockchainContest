package pkg

import (
	"database/sql"
	"fmt"
	"time"
)

//单线程
func Insert1T(db sql.DB)  {
	//打印当前时间
	start_time:=time.Now()
	fmt.Println("当前时间：")

}
