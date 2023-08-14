package mysql

import (
	"ahp-be/config"
	"ahp-be/internal/alternative"
	repoAlternativeMySQL "ahp-be/internal/alternative/repository/mysql"
	"ahp-be/internal/collection"
	repoCollectionMySQL "ahp-be/internal/collection/repository/mysql"
	"ahp-be/internal/criteria"
	repoCriteriaMySQL "ahp-be/internal/criteria/repository/mysql"
	"ahp-be/internal/subCriteria"
	repoSubCriteriaMySQL "ahp-be/internal/subCriteria/repository/mysql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepositoryMysqlImpl struct {
	Db  *gorm.DB
	log *logrus.Logger
	cfg *config.Config
}

func New(log *logrus.Logger, cfg *config.Config) *RepositoryMysqlImpl {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Db.Sql.User, cfg.Db.Sql.Password, cfg.Db.Sql.Host, cfg.Db.Sql.Port, cfg.Db.Sql.Name)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		log.Error(err)
	}

	return &RepositoryMysqlImpl{
		Db:  db,
		log: log,
		cfg: cfg,
	}
}

func (r *RepositoryMysqlImpl) CollectionRepository() collection.Repository {
	return repoCollectionMySQL.New(r.Db)
}

func (r *RepositoryMysqlImpl) AlternativeRepository() alternative.Repository {
	return repoAlternativeMySQL.New(r.Db)
}

func (r *RepositoryMysqlImpl) CriteriaRepository() criteria.Repository {
	return repoCriteriaMySQL.New(r.Db)
}

func (r *RepositoryMysqlImpl) SubCriteriaRepository() subCriteria.Repository {
	return repoSubCriteriaMySQL.New(r.Db)
}
