package seeder

import (
	"errors"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

func AnimalSeeder(db *gorm.DB, animalTypeMap map[string]string) (map[string]string, error) {
	var dummyAnimals = []entity.Animal{
		{
			Name:         "Buaya",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         1,
		},
		{
			Name:         "Ular Kobra",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         2,
		},
		{
			Name:         "Komodo",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         3,
		},
		{
			Name:         "Bunglon",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         4,
		},
		{
			Name:         "Biawak",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         5,
		},
		{
			Name:         "Ular Sanca Kembang",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         6,
		},
		{
			Name:         "Penyu Hijau",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         7,
		},
		{
			Name:         "Ular Piton Reticulatus",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         8,
		},
		{
			Name:         "Penyu Belimbing",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         9,
		},
		{
			Name:         "Cicak Tokek",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         10,
		},
		{
			Name:         "Ular Weling",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         11,
		},
		{
			Name:         "Ular Kapak Hijau",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         12,
		},
		{
			Name:         "Kadal",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         13,
		},
		{
			Name:         "Penyu Lekang",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         14,
		},
		{
			Name:         "Elang Jawa",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         1,
		},
		{
			Name:         "Cenderawasih",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         2,
		},
		{
			Name:         "Jalak Bali",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         3,
		},
		{
			Name:         "Beo Nias",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         4,
		},
		{
			Name:         "Kakatua Raja",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         5,
		},
		{
			Name:         "Enggang Gading",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         6,
		},
		{
			Name:         "Burung Hantu Celepuk",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         7,
		},
		{
			Name:         "Kakatua Maluku",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         8,
		},
		{
			Name:         "Anis Merah",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         9,
		},
		{
			Name:         "Burung Merpati Hutan",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         10,
		},
		{
			Name:         "Ikan Arwana",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         1,
		},
		{
			Name:         "Ikan Napoleon",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         2,
		},
		{
			Name:         "Ikan Pari Manta",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         3,
		},
		{
			Name:         "Ikan Betta",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         4,
		},
		{
			Name:         "Ikan Guppy",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         5,
		},
		{
			Name:         "Ikan Tuna",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         6,
		},
		{
			Name:         "Ikan Parrot",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         7,
		},
		{
			Name:         "Ikan Kerapu",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         8,
		},
		{
			Name:         "Ikan Teri",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         9,
		},
		{
			Name:         "Orangutan",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         1,
		},
		{
			Name:         "Harimau Sumatera",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         2,
		},
		{
			Name:         "Gajah Sumatera",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         3,
		},
		{
			Name:         "Badak Jawa",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         4,
		},
		{
			Name:         "Anoa",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         5,
		},
		{
			Name:         "Tapir",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         6,
		},
		{
			Name:         "Monyet Ekor Panjang",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         7,
		},
		{
			Name:         "Tarsius Sulawesi",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         8,
		},
		{
			Name:         "Beruang Madu",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         9,
		},
		{
			Name:         "Rusa Timor",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         10,
		},
		{
			Name:         "Banteng",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         11,
		},
	}

	hasTable := db.Migrator().HasTable(&entity.Animal{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Animal{}); err != nil {
			return nil, err
		}
	}

	for _, data := range dummyAnimals {
		var animal entity.Animal
		err := db.Where(&entity.Animal{Name: data.Name}).First(&animal).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		isData := db.Find(&animal, "name = ?", data.Name).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return nil, err
			}
		}
	}

	var curAnimals []entity.Animal
	if err := db.Find(&curAnimals).Error; err != nil {
		return nil, err
	}

	animalMap := map[string]string{}
	for _, data := range curAnimals {
		animalMap[data.Name] = data.ID.String()
	}
	return animalMap, nil
}
