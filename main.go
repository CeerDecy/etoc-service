package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"etoc-service/internal/app/http/handler"
	"etoc-service/internal/app/http/router"
)

const logo = ` 
  ____________________/\\\__________________________________      
   _____/\\\\\\\\___/\\\\\\\\\\\_____/\\\\\________/\\\\\\\\_     
    ___/\\\/////\\\_\////\\\////____/\\\///\\\____/\\\//////__    
     __/\\\\\\\\\\\_____\/\\\_______/\\\__\//\\\__/\\\_________   
      _\//\\///////______\/\\\_/\\__\//\\\__/\\\__\//\\\________  
       __\//\\\\\\\\\\____\//\\\\\____\///\\\\\/____\///\\\\\\\\_ 
        ___\//////////______\/////_______\/////________\////////__

`

func main() {

	print(logo)
	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT) // 监听
	go func() {
		sig := <-sigCh
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logrus.Infof("receive signal %s", sig.String())
			closeChan <- struct{}{}
		}
	}()

	engine := gin.Default()

	go func() {
		engine.GET("/shutdown", func(ctx *gin.Context) {
			closeChan <- struct{}{}
		})

		r := router.NewRouter(engine,
			router.WithSqlite(),
		)
		r.RegisterHandler(handler.Handler())

		if err := engine.Run(":26676"); err != nil {
			logrus.Errorf("start engine err, %s", err.Error())
			closeChan <- struct{}{}
		}
	}()

	for {
		select {
		case <-closeChan:
			logrus.Infof("closing etoc-service")
			for i := 0; i < 3; i++ {
				logrus.Infof("shutdown %d", 3-i)
				time.Sleep(1 * time.Second)
			}
			return
		}
	}
}
