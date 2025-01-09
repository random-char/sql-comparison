package repository

import (
	"context"
	"random-char/sql-comparison/pkg/context/artist/domain/entity"
)

type ArtistRepository interface {
	FindAll(context.Context) ([]GetArtistDTO, error)
	Insert(context.Context, *InsertArtistDTO) (*entity.Artist, error)
    Update(context.Context, *entity.Artist) (error)
    Delete(context.Context, *entity.Artist) (error)
}
