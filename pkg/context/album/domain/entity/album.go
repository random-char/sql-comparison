package entity

import (
	"random-char/sql-comparison/pkg/context/artist/domain/entity"
	"time"
)

type Album struct {
	id   int
	name string
	date time.Time

	artist *entity.Artist
}

func NewAlbum(
	id int,
	name string,
	date time.Time,
	artist *entity.Artist,
) *Album {
	return &Album{
		id:     id,
		name:   name,
		date:   date,
		artist: artist,
	}
}

func (a *Album) WithId(id int) *Album {
	return &Album{
		id:     id,
		name:   a.name,
		date:   a.date,
		artist: a.artist,
	}
}

func (a *Album) GetId() int {
	return a.id
}

func (a *Album) GetName() string {
	return a.name
}

func (a *Album) GetDate() time.Time {
	return a.date
}

func (a *Album) GetArtist() *entity.Artist {
	return a.artist
}
