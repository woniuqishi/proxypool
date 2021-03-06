package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/Alex950808/proxypool/api"
	"github.com/Alex950808/proxypool/internal/app"
	"github.com/Alex950808/proxypool/internal/cron"
	"github.com/Alex950808/proxypool/internal/database"
	"github.com/Alex950808/proxypool/pkg/proxy"
)

var configFilePath = ""

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	flag.StringVar(&configFilePath, "c", "", "path to config file: config.yaml")
	flag.Parse()

	if configFilePath == "" {
		configFilePath = os.Getenv("CONFIG_FILE")
	}
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	err := app.InitConfigAndGetters(configFilePath)
	if err != nil {
		panic(err)
	}

	database.InitTables()
	proxy.InitGeoIpDB()
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	go cron.Cron()
	api.Run()
}
