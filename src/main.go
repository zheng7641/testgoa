package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type test struct {
	Id int
	A  string
	B  string
}

func main() {
	router := gin.Default()
	router.GET("/hello", func(context *gin.Context) {
		db, err := gorm.Open("mysql", "root:@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")

		if err != nil {
			panic(err)
		}

		db.LogMode(true)

		defer db.Close()

		var test1 = new(test)
		find := db.Table("test").First(test1)
		fmt.Println(find)
		//fmt.Println(len(list))
		//fmt.Println(list)
		//for i := range list {
		//	t := list[i]
		//	fmt.Println(t.a)
		//	println(t.b)
		//}

		//db.AutoMigrate(&UserInfo{})
		//u1 := UserInfo{1, "七米", "男", "篮球"}
		//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
		//// 创建记录
		//db.Create(&u1)
		//db.Create(&u2)
		//// 查询
		//var u = new(UserInfo)
		//db.First(u)
		//fmt.Printf("%#v\n", u)

		context.JSON(200, gin.H{
			"code":    200,
			"success": true,
		})
	})
	_ = router.Run("localhost:9090")
}
