package repository

import (
	"random-char/sql-comparison/pkg/context/artist/domain/entity"
	"time"
)

type InsertAlbumDTO struct {
    Name string
    Date time.Time

    Arist *entity.Artist
}
