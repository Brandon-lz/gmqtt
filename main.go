package main

import (
	"fmt"
	"log/slog"

	"github.com/Brandon-lz/gmqtt/config"
	httpservice "github.com/Brandon-lz/gmqtt/http_service"
	"github.com/Brandon-lz/gmqtt/log"
	"github.com/Brandon-lz/gmqtt/utils"
)

// @title Gin Swagger Example API
// @version 2.0
// @description go subpub service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http
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
