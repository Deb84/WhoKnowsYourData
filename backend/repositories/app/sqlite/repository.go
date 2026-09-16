package sqlite

import (
	"whoknowsyourdata/domain"

	"gorm.io/gorm"
)

type AppRepository struct {
	log domain.Logger
	db  *gorm.DB
}

func NewAppRepository(log domain.Logger, db *gorm.DB) *AppRepository {
	return &AppRepository{
		log: log,
		db:  db,
	}
}
