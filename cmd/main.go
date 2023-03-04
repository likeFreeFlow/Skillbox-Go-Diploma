package main

import (
	"flag"
	"log"

	"sps/include/server"
	"sps/pkg/simulator"

	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Println("Configuration file not loaded")
	}

	mode := flag.String("mode", "server", "a string")
	addr := flag.String("addr", viper.GetString("server.addr"), "a string")
	port := flag.String("port", viper.GetString("server.port"), "a string")
	flag.Parse()

	simUrl := viper.GetString("simulator.addr") + ":" + viper.GetString("simulator.port")
	url := *addr + ":" + *port

	switch {
	case *mode == "server":
		server.StartServer(url)
	case *mode == "simulator":
		simulator.Start(simUrl)
	}
}

// initConfig() - initializes configuration file in yaml format
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
