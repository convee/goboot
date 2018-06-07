package goboot

import (
	"fmt"

	"github.com/convee/goboot/conf"
	"github.com/convee/goboot/logger"
)

func Run(configPath string) {
	conf.LoadTomlConfig(configPath)
	logConfig := conf.Get().Log
	logger.Init(logConfig.LogName, logConfig.LogPath)
	defer logger.Close()
	appName := conf.Get().AppName
	logger.Info(fmt.Sprintf("%s starting up ...", appName))
	defer func() {
		if r := recover(); r != nil {
			logger.Error(fmt.Sprintf("%v", r))
		}
	}()
}
