package main

import (
	"context"
	"flag"

	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1r0npipe/url-requestor/pkg/config"
	"github.com/1r0npipe/url-requestor/pkg/server"
)

func main() {
	configFile := flag.String("config", "./configs/config.yaml", "define a config YAML file")
	debugLevel := flag.String("debug", "info", "defing a level of log (info(default), debug, error)")
	config, err := config.ReadConfigFile(*configFile)
	if err != nil {
		log.Fatalf("error with reading config: %v", err)
	}
	if config.Server.LogLevel == nil {
		config.Server.LogLevel = debugLevel
	}

	// section for logger description (true: debug messages activating)
	logger := server.NewLogger(*config.Server.LogLevel)

	// section for context description
	_, cancelFunc := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancelFunc()
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig,
		syscall.SIGINT,
		syscall.SIGTERM)

	// section main service initialisation
	mainS := server.Init(config, logger)

	logger.Debug().Msgf("...starting up...")

	go func() {
		if err := mainS.Run(); err != nil {
			logger.Fatal().Msgf("server drops down: %v", err)
		}
	}()
	<-osSig
	//mainS.
	logger.Debug().Msgf("...shutting down...")
}
