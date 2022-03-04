package model

import (
	"log"

	"gorm.io/gorm"
)

func SetUp(db *gorm.DB) {
	err := db.AutoMigrate(
		&SysApi{}, &SysDept{}, &SysGroup{},
		&SysMenu{}, &SysOrg{}, &SysOrgConf{},
		&SysPost{}, &SysRole{}, &User{},
		&SysConf{}, &SysDict{})
	if err != nil {
		log.Panic(err)
	}
}
