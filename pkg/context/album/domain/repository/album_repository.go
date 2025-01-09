package repository

import (
	"context"
	"random-char/sql-comparison/pkg/context/album/domain/entity"
)

type AlbumRepository interface {
	FindAll(context.Context) ([]GetAlbumDTO, error)
	Insert(context.Context, *InsertAlbumDTO) (*entity.Album)
    Update(context.Context, *entity.Album) (error)
    Delete(context.Context, *entity.Album) (error)
}
