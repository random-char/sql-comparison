package repository

import (
	"time"
)

type GetAlbumDTO struct {
    Id int
    Name string
    Date time.Time
    ArtistId int
}
