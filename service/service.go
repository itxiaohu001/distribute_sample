package service

import (
	"context"
	"fmt"
	"net/http"
)

func Start(ctx context.Context, serviceName, host, port string, registerHandler func()) (context.Context, error) {
	registerHandler()
	return startService(ctx, serviceName, host, port), nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var svr http.Server
	svr.Addr = ":" + port

	// 启动http服务
	go func() {
		_ = svr.ListenAndServe()
		cancel()
	}()

	// 提示用户可以按任意键手动停止服务
	go func() {
		fmt.Printf("%v started, press any key to stop\n", serviceName)
		var s string
		_, _ = fmt.Scanln(&s)
		_ = svr.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
