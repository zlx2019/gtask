package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	. "zero.com/gtask-common/session"
)

// Run 	服务优雅启动与停止
// engine 	  要启动的gin服务引擎
// serverName 服务名
// serverPort 服务端口号
func Run(engine *gin.Engine, serverName, serverPort string) {
	// 通过原生httpServer + gin引擎开启服务
	server := &http.Server{
		Addr:    serverPort,
		Handler: engine,
	}
	// 这里通过一个协程来开启 http服务,保证服务能够主动性的优雅停止
	go func() {
		Logger.Infof("HTTP Server %s Running Port In %s \n", serverName, server.Addr)
		// 启动
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Logger.Panicf("HTTP Server %s Running Error: %s", err.Error())
		}
	}()

	// 由于服务由一个协程开启的,所以主协程这里我们需要手动阻塞一下,直到接受到停止信号
	// 创建一个接收操作系统信号的通道
	exit := make(chan os.Signal)
	// 这里表示如果接收到了SIGINT或者SIGTERM系统信号,则会把信号向exit通道发送.
	// syscall.SIGINT: 		用户发送INTR字符,例如在终端执行(Ctrl+C) 触发 kill -2 pid然后进程结束
	// syscall.SIGTERM: 	结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞,直到接收到两种信号其中一种...
	<-exit

	// 信号接收到后,需要一定的时间释放相应的资源。 这里延迟3秒,模拟释放资源
	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	// 停止服务
	if err := server.Shutdown(closeCtx); err != nil {
		log.Println(err)
	}

	// 等待5秒时间过后,结束程序.
	select {
	case <-closeCtx.Done():
		Logger.Info("Wait Close Resource Timeout...")
	}
	Logger.Info("HTTP Server Shutdown Success")
}
