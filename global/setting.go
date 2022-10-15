package global

import (
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"github.com/jinzhu/gorm"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	DBEngine        *gorm.DB
	Logger          *logger.Logger
)
