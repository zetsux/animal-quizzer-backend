package database

import (
	"fmt"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"github.com/zetsux/gin-gorm-clean-starter/database/seeder"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.User{},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := DBSeed(db); err != nil {
		panic(err)
	}
}

func DBSeed(db *gorm.DB) error {
	if err := seeder.UserSeeder(db); err != nil {
		return err
	}

	animalTypeMap, err := seeder.AnimalTypeSeeder(db)
	if err != nil {
		return err
	}

	animalMap, err := seeder.AnimalSeeder(db, animalTypeMap)
	if err != nil {
		return err
	}

	if err := seeder.QuizSeeder(db, animalMap); err != nil {
		return err
	}

	return nil
}
