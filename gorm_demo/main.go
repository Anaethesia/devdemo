package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	GLOBAL_DB *gorm.DB
)
func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gva_",		//表名前缀 'User'-> 't_Users'
			SingularTable: true, 	//适用单数表名'User' -> 't_User'
		},
		DisableForeignKeyConstraintWhenMigrating: true,  // 逻辑外键 （代码里面自动体现外键关系）
	})

	GLOBAL_DB = db
	// TestUserCreate()
	// UserCreate()
	UserFind()
	// UserUpdate()
	// UserDelete()

	// One2one()
	// HasOne()
	fmt.Println(db, err)
}
