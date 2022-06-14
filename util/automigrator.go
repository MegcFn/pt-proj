package util

import "github.com/MegcFn/pt-proj/model"

func Migrate() {
	db := GetDB()
	err := db.AutoMigrate(&model.Application{})
	if err != nil {
		panic(err.Error())
	}
}
