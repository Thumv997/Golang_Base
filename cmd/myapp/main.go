package main

import (
	config "lore_project/configs"
	server "lore_project/internal/config"
	logger "lore_project/pkg/logging"

	"github.com/joho/godotenv"
)

func main() {

    // Initial Logrus
	logger.Init()

	//Load .ENV
	err := godotenv.Load(".env")
	if err != nil {
		logger.GetLogger().Error("Không thể đọc tệp .env")
	}

	// Start the HTTP server
	server.Run(config.LoadConfig().Server.Port)
}
