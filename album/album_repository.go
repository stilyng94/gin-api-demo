package album

import (
	"context"

	"github.com/gin-api-demo/ent"
)

type AlbumRepository interface {
	AlbumByID(ctx context.Context, id int) (*ent.Album, error)
	AddAlbum(ctx context.Context, alb ent.Album) (*ent.Album, error)
	GetAlbums(ctx context.Context) ([]*ent.Album, error)
}

type albumRepository struct {
	dbClient *ent.Client
}

// Initialize album repository
func NewAlbumRepository(dbClient *ent.Client) AlbumRepository {
	return &albumRepository{dbClient: dbClient}
}

func (repo *albumRepository) AlbumByID(ctx context.Context, id int) (*ent.Album, error) {
	return repo.dbClient.Album.Get(ctx, id)
}

func (repo *albumRepository) AddAlbum(ctx context.Context, alb ent.Album) (*ent.Album, error) {
	return repo.dbClient.Album.Create().SetArtist(alb.Artist).SetTitle(alb.Title).SetPrice(alb.Price).Save(ctx)
}

func (repo *albumRepository) GetAlbums(ctx context.Context) ([]*ent.Album, error) {
	return repo.dbClient.Album.Query().All(ctx)
}
