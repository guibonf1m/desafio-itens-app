package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConectarGORM() (*gorm.DB, error) {
	dsn := "admin:admin@tcp(desafio-mysql:3306)/desafio_itens?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com GORM: %w", err)
	}

	err = db.AutoMigrate(&ItemModel{}, &UserModel{})
	if err != nil {
		return nil, fmt.Errorf("erro na migration: %w", err)
	}

	return db, nil
}
