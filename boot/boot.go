package boot

import (
	"flag"
	"fmt"
	"github.com/qauzy/mynahs/config"
	"github.com/qauzy/mynahs/data"
	"github.com/qauzy/mynahs/internal/pkg/logger"
	"github.com/qauzy/mynahs/internal/routers"
	"github.com/qauzy/mynahs/internal/validator"
	"os"
)

func init() {
	var (
		configPath   string
		printVersion bool = true
	)

	flag.StringVar(&configPath, "c", "", "请输入配置文件绝对路径")
	flag.BoolVar(&printVersion, "v", false, "查看版本")
	flag.Parse()

	if printVersion {
		// 打印版本号
		println(version)
		os.Exit(0)
	}

	// 1、初始化配置
	config.InitConfig(configPath)

	// 2、初始化zap日志
	logger.InitLogger()

	// 3、初始化数据库
	data.InitData()

	// 4、初始化验证器
	validator.InitValidatorTrans("zh")
}

func Run() {
	r := routers.SetRouters()
	err := r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
	if err != nil {
		panic(err)
	}
}
