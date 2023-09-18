package main

import (
	"fmt"

	"github.com/gin-api-demo/album"
	"github.com/gin-api-demo/config"
	"github.com/gin-gonic/gin"
)

func main() {
	settings := config.NewSetting()
	dbClient := config.OpenDB(settings.DatabaseUrl)
	defer dbClient.Close()

	router := gin.Default()
	album.InitializeAlbum(dbClient, router)

	router.Run(fmt.Sprintf("0.0.0.0:%v", settings.Port))
}
