package main

import (
	"fmt"
	"gorm.io/gorm"
)

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:lc"`
	Age  int    `gorm:"comment:年龄"`
}

func TestUserCreate() {
	_ = GLOBAL_DB.AutoMigrate(&TestUser{})
}

func UserCreate() {
	dbRes := GLOBAL_DB.Create(&[]TestUser{
		{Name: "凉城", Age: 23},
		{Name: "大毛", Age: 24},
		{Name: "汪汪", Age: 22},
	})
	fmt.Println("User Create ", dbRes.Error, dbRes.RowsAffected)
}

func UserFind() {
	// temp := TestUser{} // 	存放查询结果的
	//result := GLOBAL_DB.Model(&TestUser{}).First(&temp)   // model放结构体说明是那个表，查询结果放到res里面 ｜｜ 主键升序
	//fmt.Println(result.Error, result.RowsAffected)

	//var temp []*TestUser  	// 多类数据查询的
	//GLOBAL_DB.Where("Name = ?", "汪汪").Or("Name = ?", "大毛").Order("id").Find(&temp)

	var temp *TestUser
	GLOBAL_DB.Where("Name = ?", "汪汪").Or("Name = ?", "大毛").Order("id").First(temp) // where条件检索  --> sql 语句
	//GLOBAL_DB.Where(&TestUser{Name: "汪汪", Age: 22}).First(&temp)   // struct or map
	//GLOBAL_DB.Where(map[string]interface{}{"name": "汪汪", "age": 22}).First(&temp)

	//GLOBAL_DB.Select("age").Where("Name = ?", "汪汪") // 挑选字段
	fmt.Println("User Find", temp)
	//// 高级查询
	//// 智能选择字段
	//type APIUser struct {
	//	Name string
	//}
	//u := APIUser{}
	//GLOBAL_DB.Model(&TestUser{}).Where("Name <> ?", "汪汪").Find(&u)
}

func UserUpdate() {
	var temp []TestUser
	// update 		只更新指定的字段
	// updates 		更新所有的字段
	// save         无论如何更新所有的字段 包括0值
	//GLOBAL_DB.Model(&TestUser{}).Where("name LIKE ?", "%凉%").Update("name", "l")
	dbRes := GLOBAL_DB.Model(&TestUser{}).Where("name LIKE ?", "%凉%").Find(&temp) // 搜到的结构体
	for k := range temp {
		temp[k].Age = 1
	}
	dbRes.Save(&temp)
}

func UserDelete() {
	var temp []TestUser
	GLOBAL_DB.Model(&TestUser{}).Where("name LIKE ?", "%凉%").Delete(&temp)
}
