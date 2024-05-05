package models

import "rhim/middleware"

func InitSql() {
	db := middleware.GetDb()
	// 重命名字段
	// 手动迁移，重命名字段
	// 检查旧的字段名是否存在
	if db.Migrator().HasColumn(&UserBasic{}, "pass_word") {
		// 重命名字段
		db.Migrator().RenameColumn(&UserBasic{}, "pass_word", "password")
	}

	// 新表才需要
	db.AutoMigrate(
	//&UserBasic{},
	//&GroupBasic{},
	//&Message{},
	//&Contact{},
	)
}
