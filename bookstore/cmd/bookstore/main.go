package main

import (
	_ "bookstore/internal/store"
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("mem") // 创建store实例
	if err != nil {
		panic(err)
	}

	// 创建 BookStoreServer 实例
	srv := server.NewBookStoreServer(":8080", s)

	// 运行 http server
	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Println("web server start failed:", err)
		return
	}
	log.Println("web server start ok")

	// 通过监视系统信号实现了 http 服务实例的优雅退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx) // 优雅关闭 http 服务
	}

	if err != nil {
		log.Println("bookstore program exit error: ", err)
		return
	}

	log.Println("bookstore program exit ok")
}
