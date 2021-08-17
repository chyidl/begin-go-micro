package main

import (
	"github.com/chyidl/begin-go-micro/services/user/domain/repository"
	service2 "github.com/chyidl/begin-go-micro/services/user/domain/service"
	"github.com/chyidl/begin-go-micro/services/user/handler"
	pb "github.com/chyidl/begin-go-micro/services/user/proto/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"log"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("begin-go-micro.user"),
		service.Version("latest"),
	)

	// Init service
	srv.Init()

	// 创建数据库连接
	dsn := "chyi@29Wtdwv3hMmsj2gDEd7W*RDB@tcp(remote.mysql.baidu:3306)/develop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 维护数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)
	defer sqlDB.Close()

	// 只执行一次，数据库表初始化
	rp := repository.NewUserRepository(db)
	// 初始化表
	rp.InitTable()

	// 创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	// Register handler
	pb.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
