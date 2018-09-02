package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/xiaozefeng/apiserver/config"
	"github.com/xiaozefeng/apiserver/model"
	v "github.com/xiaozefeng/apiserver/pkg/version"
	"github.com/xiaozefeng/apiserver/router"
	"github.com/xiaozefeng/apiserver/router/middleware"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *version {
		v := v.Get()
		marshaled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshaled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init db
	model.DB.Init()
	defer model.DB.Close()
	// set gin mode
	gin.SetMode(viper.GetString("runmode"))
	// Create the  Gin engine.
	g := gin.New()

	// Routers
	g = router.Load(
		// Cores
		g,
		// Middleware
		middleware.RequestId(),
		middleware.Logging(),
	)

	// Ping the sever to make sure the router is working
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no resoonse, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully")
	}()
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// pingServer pings the http server  to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/heath`
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
