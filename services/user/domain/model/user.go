package model

import "time"

// User GORM
type User struct {
	// 主键
	ID uint64 `gorm:"primaryKey;notNull;autoIncrement"`
	// 用户名称
	UserName string  `gorm:"uniqueIndex;notNull"`
	// 添加需要的字段
	FirstName string
	// 密码
	HashPassword string
	//
	CreatedAt time.Time
	UpdatedAt time.Time
}