package repository


import (
	"context"
	"errors"
	"random-char/sql-comparison/pkg/context/album/domain/entity"
	"random-char/sql-comparison/pkg/context/album/domain/repository"
	"random-char/sql-comparison/pkg/db"

	"github.com/jackc/pgx/v5"
)

type postgresAlbumRepository struct {
	psql *db.Postgres
}

func New(psql *db.Postgres) *postgresAlbumRepository {
	return &postgresAlbumRepository{
		psql: psql,
	}
}

func (ar *postgresAlbumRepository) FindAll(ctx context.Context) ([]repository.GetAlbumDTO, error) {
	query := "SELECT (id, name, date, artist_id) FROM album;"
	rows, err := ar.psql.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	artists := []repository.GetAlbumDTO{}
	for rows.Next() {
		album := repository.GetAlbumDTO{}
		err := rows.Scan(
            &album.Id,
            &album.Name,
            &album.Date,
            &album.ArtistId,
        )

		if err != nil {
			return nil, err
		}

		artists = append(artists, album)
	}

	return artists, nil
}

func (ar *postgresAlbumRepository) Insert(ctx context.Context, album *repository.InsertAlbumDTO) (*entity.Album, error) {
	query := `INSERT INTO album 
        (name, date, artist_id)
        VALUES (@name, @date, @artistId)
        RETURNING id;`

	args := pgx.NamedArgs{
		"name":  album.Name,
        "date": album.Date,
        "artistId": album.Arist.GetId(),
	}

	var id int
	err := ar.psql.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		return nil, err
	}

	return entity.NewAlbum(id, album.Name, album.Date, album.Arist), nil
}

func (ar *postgresAlbumRepository) Update(ctx context.Context, album *entity.Album) error {
	return errors.New("not implemented")
}

func (ar *postgresAlbumRepository) Delete(ctx context.Context, album *entity.Album) error {
	return errors.New("not implemented")
}
