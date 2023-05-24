package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"study_go/models"
	"study_go/routes"
)

func main() {
	// 打开数据库连接
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 数据迁移
	if err := models.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// 创建一个gin实例，并设置路由
	router := gin.Default()
	routes.Setup(router, db)

	// 启动HTTP服务器
	if err := http.ListenAndServe(":8888", router); err != nil {
		fmt.Println(err)
	}
}
