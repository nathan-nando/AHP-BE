package repository

import (
	"ahp-be/internal/alternative"
	"ahp-be/internal/collection"
	"ahp-be/internal/criteria"
	"ahp-be/internal/subCriteria"
)

type Repository interface {
	CollectionRepository() collection.Repository
	AlternativeRepository() alternative.Repository
	CriteriaRepository() criteria.Repository
	SubCriteriaRepository() subCriteria.Repository
}
