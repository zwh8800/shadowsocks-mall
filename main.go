package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/zwh8800/shadowsocks-mall/conf"
	"github.com/zwh8800/shadowsocks-mall/crontab"
	"github.com/zwh8800/shadowsocks-mall/route"
)

func main() {
	defer glog.Flush()
	startServer()
	handleSignal()
	glog.Infoln("signal received")
	stopServer()
	glog.Infoln("gracefully shutdown")
}

func startServer() {
	glog.Infoln("starting...")

	if conf.Conf.Env.Prod {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	route.Route(r)

	go func() {
		err := r.Run(fmt.Sprintf("%v:%v", "", conf.Conf.Env.ServerPort))
		if err != nil {
			glog.Fatalf("error occored: %s", err)
			panic(err)
		}
	}()

	crontab.Go()
	glog.Infoln("server started")
}

func handleSignal() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	<-signalChan
}

func stopServer() {
}
