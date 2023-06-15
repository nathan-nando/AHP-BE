package app

import (
	"ahp-be/config"
	"ahp-be/docs"
	collection "ahp-be/internal/collection/http"
	"ahp-be/internal/middleware"
	"ahp-be/internal/model"
	"ahp-be/internal/repository"
	"ahp-be/internal/repository/mysql"
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
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	docs.SwaggerInfo.Version = fmt.Sprintf("%s", cfg.Server.Version)

	log := logrus.New()

	var repo *mysql.RepositoryMysqlImpl

	if dbMode == "mysql" {
		repo = mysql.New(log, cfg)
		log.Info("DB MODE = MYSQL")
	} else if dbMode == "postgresql" {
		log.Info("DB MODE = POSTGRESQL")
	} else {
		log.Info("DB MODE = MONGODB")
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
}

func RegisterRepository(repo *mysql.RepositoryMysqlImpl, log *logrus.Logger) {
	err := repo.Db.AutoMigrate(&model.CollectionModel{})
	if err != nil {
		log.Error(err)
	}
}
