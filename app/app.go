package app

import (
	"ahp-be/config"
	"ahp-be/docs"
	alternative "ahp-be/internal/alternative/http"
	collection "ahp-be/internal/collection/http"
	criteria "ahp-be/internal/criteria/http"
	"ahp-be/internal/middleware"
	"ahp-be/internal/model"
	"ahp-be/internal/repository"
	"ahp-be/internal/repository/mysql"
	subCriteria "ahp-be/internal/subCriteria/http"
	"ahp-be/pkg/constants"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

type Api struct {
	repository repository.Repository
	cfg        *config.Config
	log        *logrus.Logger
	routes     *echo.Echo
}

func New(cfg *config.Config, dbMode string) *Api {
	docs.SwaggerInfo.Title = fmt.Sprintf("%s", cfg.Server.Name)
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	docs.SwaggerInfo.Version = fmt.Sprintf("%s", cfg.Server.Version)

	log := logrus.New()

	var repo *mysql.RepositoryMysqlImpl

	if dbMode == constants.Mysql {
		repo = mysql.New(log, cfg)
		log.Info(fmt.Sprintf("DB MODE = %s", constants.Mysql))
	} else if dbMode == constants.Postgresql {
		log.Info(fmt.Sprintf("DB MODE = %s", constants.Postgresql))
	} else {
		log.Info(fmt.Sprintf("DB MODE = %s", constants.Mongodb))
	}

	routes := echo.New()
	middleware.Init(routes)

	RegisterHandler(routes, repo)
	RegisterRepository(repo, log)

	return &Api{
		repository: repo,
		cfg:        cfg,
		log:        log,
		routes:     routes,
	}
}

func RegisterHandler(e *echo.Echo, repo *mysql.RepositoryMysqlImpl) {
	e.GET("/", func(c echo.Context) error {
		message := "API TPS LOCATION USING AHP"
		return c.String(http.StatusOK, message)
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	collection.NewHandler(repo.CollectionRepository(), repo.Db).Route(e.Group("/collection"))
	alternative.NewHandler(repo.AlternativeRepository(), repo.Db).Route(e.Group("/alternative"))
	criteria.NewHandler(repo.CriteriaRepository(), repo.AlternativeRepository(), repo.Db).Route(e.Group("/ahp"))
	subCriteria.NewHandler(repo.SubCriteriaRepository(), repo.Db).Route(e.Group("/ahp"))

}

func RegisterRepository(repo *mysql.RepositoryMysqlImpl, log *logrus.Logger) {
	err := repo.Db.AutoMigrate(&model.CollectionModel{}, &model.AlternativeModel{}, &model.ScoreModel{}, &model.FinalScoreModel{})
	if err != nil {
		log.Error(err)
	}
}
