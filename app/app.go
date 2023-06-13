package app

import (
	"ahp-be/config"
	"ahp-be/docs"
	"ahp-be/internal/repository"
	"ahp-be/internal/repository/mysql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
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
	//log.SetFormatter(&logrus.JSONFormatter{})

	var repo repository.Repository

	if dbMode == "mysql" {
		repo = mysql.New(log, cfg)
		log.Info("SELECTED MYSQL")
	} else if dbMode == "postgresql" {
		log.Info("postgresql")
	} else {
		log.Info("mongodb")
	}

	routes := echo.New()

	return &Api{
		repository: repo,
		cfg:        cfg,
		log:        log,
		routes:     routes,
	}
}
