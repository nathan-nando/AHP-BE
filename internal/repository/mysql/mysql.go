package mysql

import (
	"ahp-be/config"
	"ahp-be/internal/collection"
	repoMySQL "ahp-be/internal/collection/repository/mysql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepositoryMysqlImpl struct {
	db  *gorm.DB
	log *logrus.Logger
	cfg *config.Config
}

func New(log *logrus.Logger, cfg *config.Config) *RepositoryMysqlImpl {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Db.Sql.User, cfg.Db.Sql.Password, cfg.Db.Sql.Host, cfg.Db.Sql.Port, cfg.Db.Sql.Name)

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
		db:  db,
		log: log,
		cfg: cfg,
	}
}

func (r *RepositoryMysqlImpl) CollectionRepository() collection.Repository {
	return repoMySQL.New(r.db)
}
