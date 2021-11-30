package main

import "gorm.io/gorm"

// belongs to
//type Dog struct {
//	gorm.Model
//	Name string
//	GirlGodID uint
//	GirlGod GirlGod
//}
//
//type GirlGod struct {
//	gorm.Model
//	Name string
//}
//
//func One2one() {
//	g := GirlGod{
//		Model: gorm.Model{
//			ID: 2,
//		},
//		Name: "女神2",
//	}
//
//	d := Dog{
//		Model: gorm.Model{
//			ID :1,
//		},
//		Name: "xiaohu",
//		GirlGod: g,
//	}
//
//	GLOBAL_DB.Create(&d)
//	_ = GLOBAL_DB.AutoMigrate(&Dog{})
//	GLOBAL_DB.Model(&d).Association("GirlGod").Append(&g) // .Delete(&g) .Replace(&g, &g_new)
//}

// has one
//type Dog struct {
//	gorm.Model
//	Name string
//	GirlGodID uint 		// 狗链子
//
//}
//
//type GirlGod struct {
//	gorm.Model
//	Name string
//	Dog Dog				// has one
//}
//
//func HasOne() {
//	//d := Dog{
//	//	Model: gorm.Model{
//	//		ID :1,
//	//	},
//	//	Name: "xiaohu",
//	//}
//	//g := GirlGod{
//	//	Model: gorm.Model{
//	//		ID: 2,
//	//	},
//	//	Name: "女神2",
//	//	Dog: d,
//	//}
//	temp := GirlGod{}
//
//	GLOBAL_DB.Preload("Dog").Find(&temp, 2)     // 查询preload GirlGod结构体内的Dog（name）
//	fmt.Println("hasone", temp)
//	//_ = GLOBAL_DB.AutoMigrate(&GirlGod{}, &Dog{})
//}

// one2many
//func One2many() {
//	// 一对多关系和preload预加载用法
//	GLOBAL_DB.Preload("Dogs", func(db *gorm.DB) *gorm.DB{
//		return db.Joins("Info").Where("inquire condition")
//	}).Find(&d)
//}

// many2many
type Dog struct {
	gorm.Model
	Name string
	GirlGods []GirlGod		`gorm:"many2many:dog_girl_god"`    // 统一的外键记录下来

}

type GirlGod struct {
	gorm.Model
	Name string
	Dogs []Dog				`gorm:"many2many:dog_girl_god"`
}



