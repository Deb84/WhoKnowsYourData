package sqlite

import (
	"whoknowsyourdata/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(env models.SqliteEnv) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(env.PATH), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
