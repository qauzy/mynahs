package job

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zzpu/mynahs/internal/controller/v1/ws"
	"time"
)

//# 文件格式說明
//# ┌──分鐘（0 - 59）
//# │  ┌──小時（0 - 23）
//# │  │  ┌──日（1 - 31）
//# │  │  │  ┌─月（1 - 12）
//# │  │  │  │  ┌─星期（0 - 6，表示从周日到周六）
//# │  │  │  │  │
//# *  *  *  *  * 被執行的命令

func init() {
	c := cron.New()
	EntryID, err := c.AddFunc("* */1 * * *", func() {
		ws.WebsocketManager.SendAll([]byte("clock"))
	})
	fmt.Println(time.Now(), EntryID, err)

	c.Start()
}
