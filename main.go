package main

import (
	"fmt"
	"log/slog"

	"github.com/Brandon-lz/gmqtt/config"
	httpservice "github.com/Brandon-lz/gmqtt/http_service"
	"github.com/Brandon-lz/gmqtt/log"
	"github.com/Brandon-lz/gmqtt/utils"
)


func main() {
	defer utils.RecoverAndLog()
	if err := config.LoadConfig("config.toml"); err != nil {
		slog.Error(fmt.Sprintf("Failed to load config: %v", err))
		return
	}

	log.Init(config.Config.Logging.Level)
	httpservice.Start()
}

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./data_collector .
