// Package cmd
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"hotgo/internal/crons"
	"hotgo/internal/service"
	"os"
)

var (
	Cron = &gcmd.Command{
		Name:        "cron",
		Brief:       "定时任务，用来部署一些可独立运行的定时任务，通过tcp方式和后台保持长连接通讯，动态调整任务属性。",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 启动定时任务
			service.SysCron().StartCron(ctx)

			// tcp客户端
			service.CronClient().Start(ctx)

			// 退出信号监听
			signalListen(ctx, func(sig os.Signal) {
				service.CronClient().Stop(ctx)
				crons.StopALL()
				serverCloseSignal <- struct{}{}
			})

			select {
			case <-serverCloseSignal:
				// ...
			}
			return
		},
	}
)