package global

import (
    "github.com/spf13/viper"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "gin_test/conf"
)

type Application struct {
    ConfigViper *viper.Viper
    Config conf.Configuration
    Log *zap.Logger
    DB *gorm.DB
}

var App = new(Application)