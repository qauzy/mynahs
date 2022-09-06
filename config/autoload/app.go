package autoload

import (
	"github.com/zzpu/mynahs/pkg/convert"
	"github.com/zzpu/mynahs/pkg/utils"
)

type AppConfig struct {
	AppEnv         string `ini:"app_env" yaml:"app_env"`
	Debug          bool   `ini:"debug" yaml:"debug"`
	Language       string `ini:"language" yaml:"language"`
	StaticBasePath string `ini:"base_path" yaml:"base_path"`
}

var App = AppConfig{
	AppEnv:         "local",
	Debug:          true,
	Language:       "zh_CN",
	StaticBasePath: getDefaultPath(),
}

func getDefaultPath() (path string) {
	path, _ = utils.GetDefaultPath()
	path = convert.GetString(utils.If(path != "", path, "/tmp"))
	return
}
