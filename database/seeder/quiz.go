package seeder

import (
	"errors"

	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

func QuizSeeder(db *gorm.DB, animalMap map[string]string) error {
	var dummyQuizs = []entity.Quiz{
		{
			Question:      "Di mana Orangutan dapat ditemukan di alam liar?",
			CorrectAnswer: "Asia Tenggara",
			WrongAnswer:   []string{"Afrika", "Amerika Selatan", "Australia"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Apa arti nama 'Orangutan'?",
			CorrectAnswer: "Manusia Rimba",
			WrongAnswer:   []string{"Hewan Pintar", "Pengembara Hutan", "Penjaga Rimba"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Apa ciri khas fisik utama dari Orangutan?",
			CorrectAnswer: "Mereka memiliki lengan yang sangat panjang",
			WrongAnswer:   []string{"Mereka memiliki ekor panjang", "Mereka berbulu biru", "Mereka berjalan dengan dua kaki"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Apa makanan utama Orangutan?",
			CorrectAnswer: "Buah-buahan",
			WrongAnswer:   []string{"Daging", "Ikan", "Serangga"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Berapa lama periode kehamilan Orangutan betina?",
			CorrectAnswer: "10-11 bulan",
			WrongAnswer:   []string{"3-4 bulan", "6-7 bulan", "8-9 bulan"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Apa yang paling mengancam kelangsungan hidup Orangutan di habitatnya?",
			CorrectAnswer: "Penggundulan hutan",
			WrongAnswer:   []string{"Pencemaran laut", "Pemburuan ilegal", "Perubahan iklim"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Bagaimana Orangutan biasanya berkomunikasi satu sama lain?",
			CorrectAnswer: "Dengan suara keras",
			WrongAnswer:   []string{"Melalui bau", "Dengan berjabat tangan", "Melalui gerakan tarian"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Berapa umur rata-rata Orangutan di alam liar?",
			CorrectAnswer: "30-40 tahun",
			WrongAnswer:   []string{"Kurang dari 20 tahun", "20-30 tahun", "40-50 tahun"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Spesies Orangutan mana yang paling langka?",
			CorrectAnswer: "Orangutan Tapanuli",
			WrongAnswer:   []string{"Orangutan Sumatera", "Orangutan Borneo", "Orangutan Jawa"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Bagaimana Orangutan mempersiapkan tempat tidurnya?",
			CorrectAnswer: "Mereka membuat sarang di pohon",
			WrongAnswer:   []string{"Mereka tidur di tanah", "Mereka tidur di gua", "Mereka menggali lubang"},
			AnimalID:      animalMap["Orangutan"],
		},
		{
			Question:      "Apa karakteristik utama yang membedakan buaya dari aligator?",
			CorrectAnswer: "Bentuk moncong yang lebih panjang dan sempit",
			WrongAnswer:   []string{"Ukuran tubuh yang lebih kecil", "Bisa hidup di air tawar saja", "Tidak memiliki gigi taring"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Di mana habitat asli buaya dapat ditemukan?",
			CorrectAnswer: "Di sungai, danau, dan rawa",
			WrongAnswer:   []string{"Di pegunungan", "Di padang pasir", "Di hutan belantara"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Buaya dikenal memiliki indra yang sangat tajam untuk?",
			CorrectAnswer: "Mengesan getaran di air",
			WrongAnswer:   []string{"Mendengar suara jauh", "Mengidentifikasi warna", "Mencium bau di udara"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Apa yang menjadi makanan utama buaya?",
			CorrectAnswer: "Ikan dan mamalia kecil",
			WrongAnswer:   []string{"Tumbuhan dan buah-buahan", "Serangga dan amfibi", "Burung dan telur"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Bagaimana cara buaya mengatur suhu tubuhnya?",
			CorrectAnswer: "Berjemur di bawah sinar matahari",
			WrongAnswer:   []string{"Menggigil untuk menghasilkan panas", "Berganti kulit secara berkala", "Berenang di air dingin"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Buaya termasuk dalam kelompok hewan apa?",
			CorrectAnswer: "Reptil",
			WrongAnswer:   []string{"Amfibi", "Pisces", "Mamalia"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Apa metode reproduksi buaya?",
			CorrectAnswer: "Bertelur",
			WrongAnswer:   []string{"Melahirkan", "Fragmentasi", "Berkembang biak dengan tunas"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Buaya dikenal sebagai hewan yang?",
			CorrectAnswer: "Bersifat soliter",
			WrongAnswer:   []string{"Bersifat sosial", "Hidup dalam kelompok besar", "Migrasi jarak jauh"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Buaya biasanya aktif pada waktu?",
			CorrectAnswer: "Malam hari (nokturnal)",
			WrongAnswer:   []string{"Siang hari (diurnal)", "Senja (krepuskular)", "Tidak memiliki pola aktivitas tetap"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Apa perilaku khas buaya saat berburu?",
			CorrectAnswer: "Mengintai mangsanya dari bawah air",
			WrongAnswer:   []string{"Berlari cepat di darat", "Memanjat pohon", "Membuat jebakan"},
			AnimalID:      animalMap["Buaya"],
		},
		{
			Question:      "Apa ciri khas Ular Kobra ketika merasa terancam?",
			CorrectAnswer: "Mengembangkan lehernya membentuk seperti tudung",
			WrongAnswer:   []string{"Melilit mangsanya", "Menggali tanah", "Menyemburkan racun dari jarak jauh"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Bagaimana Ular Kobra mengeluarkan racunnya?",
			CorrectAnswer: "Melalui gigitannya",
			WrongAnswer:   []string{"Melalui sisiknya", "Melalui suaranya", "Melalui matanya"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Apa fungsi utama racun dari Ular Kobra?",
			CorrectAnswer: "Untuk melumpuhkan mangsa",
			WrongAnswer:   []string{"Untuk mencerna makanan", "Untuk menarik pasangan", "Untuk membersihkan diri"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Ular Kobra biasanya ditemukan di habitat apa?",
			CorrectAnswer: "Hutan dan padang rumput",
			WrongAnswer:   []string{"Gurun pasir", "Kutub", "Gunung berapi aktif"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Apa warna umum Ular Kobra?",
			CorrectAnswer: "Hitam atau coklat",
			WrongAnswer:   []string{"Merah muda", "Biru neon", "Hijau terang"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Apa periode aktivitas utama Ular Kobra?",
			CorrectAnswer: "Siang hari (diurnal)",
			WrongAnswer:   []string{"Malam hari (nokturnal)", "Senja (krepuskular)", "Variabel"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Berapa lama Ular Kobra bisa hidup di alam liar?",
			CorrectAnswer: "20 tahun",
			WrongAnswer:   []string{"1 tahun", "5 tahun", "50 tahun"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Bagaimana Ular Kobra melindungi dirinya dari predator?",
			CorrectAnswer: "Menyemburkan racun",
			WrongAnswer:   []string{"Berpura-pura mati", "Berlari sangat cepat", "Bersembunyi di dalam air"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Apa jenis makanan utama Ular Kobra?",
			CorrectAnswer: "Amfibi, seperti katak",
			WrongAnswer:   []string{"Buah-buahan", "Dedaunan", "Serangga"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Apa yang dilakukan Ular Kobra saat bertemu dengan Ular Kobra lainnya selama musim kawin?",
			CorrectAnswer: "Melakukan tarian kawin yang kompleks",
			WrongAnswer:   []string{"Menghindar dan bersembunyi", "Bertarung sampai mati", "Berkomunikasi dengan suara"},
			AnimalID:      animalMap["Ular Kobra"],
		},
		{
			Question:      "Di pulau mana Komodo dapat ditemukan di alam liar?",
			CorrectAnswer: "Pulau Komodo di Indonesia",
			WrongAnswer:   []string{"Pulau Kalimantan", "Pulau Sulawesi", "Pulau Sumatra"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Apa makanan utama Komodo?",
			CorrectAnswer: "Rusa, babi, dan kadang kala bangkai",
			WrongAnswer:   []string{"Buah-buahan", "Dedaunan", "Ikan"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Apa metode pemburuan yang sering digunakan oleh Komodo?",
			CorrectAnswer: "Menunggu dalam penyergapan dan mengandalkan gigitan beracun",
			WrongAnswer:   []string{"Berburu dalam kelompok", "Berpindah dari pohon ke pohon", "Menyelam di dalam air"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Berapa berat rata-rata Komodo dewasa?",
			CorrectAnswer: "Sekitar 70 kg",
			WrongAnswer:   []string{"Sekitar 15 kg", "Sekitar 40 kg", "Lebih dari 100 kg"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Apa yang unik dari gigitan Komodo?",
			CorrectAnswer: "Mengandung bakteri yang mematikan dan bisa mengakibatkan sepsis",
			WrongAnswer:   []string{"Menyebabkan pendarahan yang tidak berhenti", "Langsung mematikan", "Tidak berbahaya"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Bagaimana cara Komodo mengatur suhu tubuhnya?",
			CorrectAnswer: "Berjemur di bawah matahari",
			WrongAnswer:   []string{"Berganti kulit seperti ular", "Berenang", "Menggali tanah"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Apa yang dilakukan Komodo untuk menentukan lokasi mangsanya?",
			CorrectAnswer: "Menggunakan lidahnya yang bercabang untuk mengesan bau",
			WrongAnswer:   []string{"Menggunakan ekornya untuk merasakan getaran", "Menggunakan mata yang sangat tajam", "Mendengarkan suara mangsa"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Apa perilaku reproduksi Komodo?",
			CorrectAnswer: "Bertelur",
			WrongAnswer:   []string{"Melahirkan", "Berkembang biak dengan spora", "Berkembang biak dengan tunas"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Berapa lama masa inkubasi telur Komodo?",
			CorrectAnswer: "Sekitar 8 bulan",
			WrongAnswer:   []string{"Sekitar 2 bulan", "Sekitar 4 bulan", "Sekitar 12 bulan"},
			AnimalID:      animalMap["Komodo"],
		},
		{
			Question:      "Apa status konservasi Komodo?",
			CorrectAnswer: "Terancam Punah",
			WrongAnswer:   []string{"Tidak Terancam", "Rentan", "Spesies Mengganggu"},
			AnimalID:      animalMap["Komodo"],
		},
	}

	hasTable := db.Migrator().HasTable(&entity.Quiz{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Quiz{}); err != nil {
			return err
		}
	}

	for _, data := range dummyQuizs {
		var quiz entity.Quiz
		err := db.Where(&entity.Quiz{Question: data.Question}).First(&quiz).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&quiz, "question = ?", data.Question).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	var curQuizs []entity.Quiz
	if err := db.Find(&curQuizs).Error; err != nil {
		return err
	}
	return nil
}
