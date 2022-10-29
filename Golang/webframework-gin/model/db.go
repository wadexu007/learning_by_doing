package model

import (
	"demo-gin/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBConfig.User,
		config.DBConfig.Password,
		config.DBConfig.Host,
		config.DBConfig.Name)

	DB, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})

	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to connect database")
	}

	if config.DBConfig.ConnPool != nil {
		sqlDB.SetConnMaxLifetime(time.Duration(config.DBConfig.ConnPool.MaxConnLifeTime))
		sqlDB.SetMaxIdleConns(int(config.DBConfig.ConnPool.MaxIdleConns))
		sqlDB.SetMaxOpenConns(int(config.DBConfig.ConnPool.MaxOpenConns))
	}

	return DB
}
