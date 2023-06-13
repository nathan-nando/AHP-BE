package repository

import (
	"ahp-be/internal/collection"
)

type Repository interface {
	CollectionRepository() collection.Repository
}
