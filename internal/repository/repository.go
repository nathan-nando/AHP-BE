package repository

import (
	"ahp-be/internal/alternative"
	"ahp-be/internal/collection"
)

type Repository interface {
	CollectionRepository() collection.Repository
	AlternativeRepository() alternative.Repository
}
