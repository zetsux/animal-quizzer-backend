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
			Hint:         "Hewan reptil air yang memiliki rahang sangat kuat",
		},
		{
			Name:         "Ular Kobra",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         2,
			Hint:         "Dikenal dengan leher yang dapat mengembang",
		},
		{
			Name:         "Komodo",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         3,
			Hint:         "Reptil besar dari Indonesia yang dapat berlari cepat",
		},
		{
			Name:         "Bunglon",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         4,
			Hint:         "Dapat mengubah warna kulitnya",
		},
		{
			Name:         "Biawak",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         5,
			Hint:         "Reptil yang mirip komodo tetapi lebih kecil",
		},
		{
			Name:         "Ular Sanca Kembang",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         6,
			Hint:         "Ular besar yang melilit mangsanya",
		},
		{
			Name:         "Penyu Hijau",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         7,
			Hint:         "Hewan laut yang kembali ke pantai kelahirannya untuk bertelur",
		},
		{
			Name:         "Ular Piton Reticulatus",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         8,
			Hint:         "Salah satu jenis ular terpanjang di dunia",
		},
		{
			Name:         "Penyu Belimbing",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         9,
			Hint:         "Memiliki cangkang berbentuk bintang",
		},
		{
			Name:         "Cicak Tokek",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         10,
			Hint:         "Dikenal karena suaranya yang khas",
		},
		{
			Name:         "Ular Weling",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         11,
			Hint:         "Ular berbisa dengan pola warna yang mencolok",
		},
		{
			Name:         "Ular Kapak Hijau",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         12,
			Hint:         "Memiliki keunikan pada bentuk kepala yang pipih dan lebar",
		},
		{
			Name:         "Kadal",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         13,
			Hint:         "Reptil kecil yang sering ditemukan di dinding atau pohon",
		},
		{
			Name:         "Penyu Lekang",
			AnimalTypeID: animalTypeMap["Reptil"],
			Step:         14,
			Hint:         "Penyu yang lebih sering ditemukan di perairan Asia Tenggara",
		},
		{
			Name:         "Elang Jawa",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         1,
			Hint:         "Burung pemangsa ini adalah lambang nasional Indonesia",
		},
		{
			Name:         "Cenderawasih",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         2,
			Hint:         "Burung dengan bulu yang sangat indah dan berwarna-warni",
		},
		{
			Name:         "Jalak Bali",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         3,
			Hint:         "Burung kecil yang hampir punah dari Bali",
		},
		{
			Name:         "Beo Nias",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         4,
			Hint:         "Burung yang dapat menirukan suara manusia",
		},
		{
			Name:         "Kakatua Raja",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         5,
			Hint:         "Memiliki jambul kuning mencolok",
		},
		{
			Name:         "Enggang Gading",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         6,
			Hint:         "Burung dengan paruh besar dan melengkung",
		},
		{
			Name:         "Burung Hantu Celepuk",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         7,
			Hint:         "Burung hantu kecil yang aktif di malam hari",
		},
		{
			Name:         "Kakatua Maluku",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         8,
			Hint:         "Burung ini berasal dari kepulauan Maluku",
		},
		{
			Name:         "Anis Merah",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         9,
			Hint:         "Dikenal karena kicauannya yang merdu",
		},
		{
			Name:         "Burung Merpati Hutan",
			AnimalTypeID: animalTypeMap["Aves"],
			Step:         10,
			Hint:         "Burung ini sering terlihat di hutan-hutan Indonesia",
		},
		{
			Name:         "Ikan Arwana",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         1,
			Hint:         "Ikan hias yang dipercaya membawa keberuntungan",
		},
		{
			Name:         "Ikan Napoleon",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         2,
			Hint:         "Ikan karang besar dengan tonjolan dahi khas",
		},
		{
			Name:         "Ikan Pari Manta",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         3,
			Hint:         "Ikan berbentuk piring dengan sirip dada yang lebar",
		},
		{
			Name:         "Ikan Betta",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         4,
			Hint:         "Ikan hias kecil yang agresif saat bertemu sesamanya",
		},
		{
			Name:         "Ikan Guppy",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         5,
			Hint:         "Ikan kecil yang mudah berbiak di akuarium",
		},
		{
			Name:         "Ikan Tuna",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         6,
			Hint:         "Ikan yang sering diolah menjadi sashimi",
		},
		{
			Name:         "Ikan Parrot",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         7,
			Hint:         "Ikan tropis dengan warna cerah seperti burung beo",
		},
		{
			Name:         "Ikan Kerapu",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         8,
			Hint:         "Populer di restoran seafood, terkenal akan teksturnya yang lembut",
		},
		{
			Name:         "Ikan Teri",
			AnimalTypeID: animalTypeMap["Pisces"],
			Step:         9,
			Hint:         "Ikan kecil yang sering dijadikan sebagai camilan gurih",
		},
		{
			Name:         "Orangutan",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         1,
			Hint:         "Primata besar dengan bulu merah",
		},
		{
			Name:         "Harimau Sumatera",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         2,
			Hint:         "Kucing besar dengan loreng hitam",
		},
		{
			Name:         "Gajah Sumatera",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         3,
			Hint:         "Mamalia besar dengan belalai panjang",
		},
		{
			Name:         "Badak Jawa",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         4,
			Hint:         "Hewan dengan satu cula di atas hidung",
		},
		{
			Name:         "Anoa",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         5,
			Hint:         "Bison kecil endemik dari Sulawesi",
		},
		{
			Name:         "Tapir",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         6,
			Hint:         "Memiliki moncong panjang seperti belalai",
		},
		{
			Name:         "Monyet Ekor Panjang",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         7,
			Hint:         "Monyet yang sering ditemui di seluruh wilayah Indonesia",
		},
		{
			Name:         "Tarsius Sulawesi",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         8,
			Hint:         "Primata kecil dengan mata besar",
		},
		{
			Name:         "Beruang Madu",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         9,
			Hint:         "Beruang kecil dari Asia Tenggara",
		},
		{
			Name:         "Rusa Timor",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         10,
			Hint:         "Rusa yang banyak ditemukan di Nusa Tenggara",
		},
		{
			Name:         "Banteng",
			AnimalTypeID: animalTypeMap["Mamalia"],
			Step:         11,
			Hint:         "Mempunyai dua tanduk dan sering ditemukan di Jawa",
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
