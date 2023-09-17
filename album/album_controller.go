package album

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-api-demo/ent"
	"github.com/gin-gonic/gin"
)

type albumController struct {
	albumRepo AlbumRepository
}

func NewAlbumController(albumRepo AlbumRepository) *albumController {
	return &albumController{albumRepo: albumRepo}
}

// [GetAlbums] responds with the list of all albums as JSON.
func (handler *albumController) GetAlbums(c *gin.Context) {
	albums, err := handler.albumRepo.GetAlbums(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": albums})
}

// [PostAlbums] adds an album from JSON received in the request body.
func (handler *albumController) PostAlbums(c *gin.Context) {
	var albumPayload ent.Album

	if err := c.BindJSON(&albumPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newAlbum, err := handler.albumRepo.AddAlbum(c, albumPayload)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newAlbum})
}

// [GetAlbumByID] locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (handler *albumController) GetAlbumByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("album with id %v not found", id)})
	}

	album, err := handler.albumRepo.AlbumByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("album with id %d not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": album})

}
