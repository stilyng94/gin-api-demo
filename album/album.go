package album

import (
	"github.com/gin-api-demo/ent"
	"github.com/gin-gonic/gin"
)

// InitializeAlbum - initializes the album router
func InitializeAlbum(dbClient *ent.Client, router *gin.Engine) {
	albumRepo := NewAlbumRepository(dbClient)
	albumController := NewAlbumController(albumRepo)
	albumRouter := router.Group("/albums")
	{
		albumRouter.GET("/:id", albumController.GetAlbumByID)
		albumRouter.DELETE("/:id", albumController.DeleteAlbumByID)
		albumRouter.GET("", albumController.GetAlbums)
		albumRouter.POST("", albumController.PostAlbums)
	}

}
