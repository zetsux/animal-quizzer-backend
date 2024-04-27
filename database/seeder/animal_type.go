package seeder

import (
	"errors"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

func AnimalTypeSeeder(db *gorm.DB) (map[string]string, error) {
	var animalTypes = []entity.AnimalType{
		{
			Name: "Reptil",
		},
		{
			Name: "Aves",
		},
		{
			Name: "Pisces",
		},
		{
			Name: "Mamalia",
		},
	}

	hasTable := db.Migrator().HasTable(&entity.AnimalType{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.AnimalType{}); err != nil {
			return nil, err
		}
	}

	for _, data := range animalTypes {
		var animal_type entity.AnimalType
		err := db.Where(&entity.AnimalType{Name: data.Name}).First(&animal_type).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		isData := db.Find(&animal_type, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return nil, err
			}
		}
	}

	var curAnimalTypes []entity.AnimalType
	if err := db.Find(&curAnimalTypes).Error; err != nil {
		return nil, err
	}

	animalTypeMap := map[string]string{}
	for _, data := range curAnimalTypes {
		animalTypeMap[data.Name] = data.ID.String()
	}
	return animalTypeMap, nil
}
