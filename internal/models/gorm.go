package models

import "gorm.io/gorm"

func Commit(db *gorm.DB, err error) {
	//if db.TranslateError == fa {
	//	return
	//}
	if err != nil {
		db.Rollback()
	} else {
		db.Commit()
	}
}
