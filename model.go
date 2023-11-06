package go_package_db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func Init(host, port, username, password, dbname string) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库链接失败：%s", err.Error()))
	}

	DB = db
}

func AutoMigrate() {
	DB.Logger = logger.Default.LogMode(logger.Silent) // 关闭慢查询日志
	_ = DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(Collect{})
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}
