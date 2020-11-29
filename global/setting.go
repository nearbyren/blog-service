package global

import (
	"github.com/nearbyren/blog-service/pkg/logger"
	"github.com/nearbyren/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
