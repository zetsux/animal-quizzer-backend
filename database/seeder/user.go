package seeder

import (
	"errors"
	"time"

	"github.com/zetsux/gin-gorm-clean-starter/common/constant"
	"github.com/zetsux/gin-gorm-clean-starter/core/entity"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) error {
	var dummyUsers = []entity.User{
		{
			Username:    "admin1",
			Password:    "admin1",
			Role:        constant.EnumRoleAdmin,
			LastAttempt: time.Now(),
		},
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, data := range dummyUsers {
		var user entity.User
		err := db.Where(&entity.User{Username: data.Username}).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&user, "username = ?", data.Username).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
