/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 8:07 下午
 **/
package models

import (
	"fmt"
	. "gaea/gateway/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Cfg.DBUser, Cfg.DBPwd, Cfg.DBHost, Cfg.DBPort, Cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	// 初始化表
	return db.AutoMigrate(
		new(App),
		new(Task),
		new(Job),
	)
}
