package global

import (
	"go_blog/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config       *config.Config
	DB           *gorm.DB
	Log          *logrus.Logger
	MysqlLog     logger.Interface
	InterceptApi *config.InterceptApiYaml
)
