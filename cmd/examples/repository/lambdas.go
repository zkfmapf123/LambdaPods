package repository

import "gorm.io/gorm"

type LambdasRepository struct {
	db *gorm.DB
}

func NewLambdasRepository(db *gorm.DB) *LambdasRepository {
	return &LambdasRepository{
		db: db,
	}
}
