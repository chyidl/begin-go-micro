package main

import (
	"fmt"
	"github.com/chyidl/begin-go-micro/services/user/domain/repository"
	service2 "github.com/chyidl/begin-go-micro/services/user/domain/service"
	"github.com/chyidl/begin-go-micro/services/user/handler"
	pb "github.com/chyidl/begin-go-micro/services/user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
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
	db, err := gorm.Open("mysql", "user:password@/db?charset=utf8&parseTime=True&local=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.SingularTable(true)

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
