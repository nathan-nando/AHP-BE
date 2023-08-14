package mysql

import (
	"ahp-be/internal/model"
	"ahp-be/pkg/constants"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"os"
)

type SubCriteriaRepoIMPL struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *SubCriteriaRepoIMPL {
	return &SubCriteriaRepoIMPL{Db: db}
}

func (s *SubCriteriaRepoIMPL) FindSubCriteria(ctx context.Context) (*model.SubCriteria, error) {

	jsonFile, err := os.ReadFile(constants.FileSubCriteria)

	if err != nil {
		return nil, err
	}

	var data model.SubCriteria
	err = json.Unmarshal(jsonFile, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
