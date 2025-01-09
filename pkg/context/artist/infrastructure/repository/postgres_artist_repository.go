package repository

import (
	"context"
	"errors"
	"random-char/sql-comparison/pkg/context/artist/domain/entity"
	"random-char/sql-comparison/pkg/context/artist/domain/repository"
	"random-char/sql-comparison/pkg/db"

	"github.com/jackc/pgx/v5"
)

type postgresArtistRepository struct {
	psql *db.Postgres
}

func New(psql *db.Postgres) *postgresArtistRepository {
	return &postgresArtistRepository{
		psql: psql,
	}
}

func (ar *postgresArtistRepository) FindAll(ctx context.Context) ([]repository.GetArtistDTO, error) {
	query := "SELECT * FROM artist;"
	rows, err := ar.psql.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	artists := []repository.GetArtistDTO{}
	for rows.Next() {
		artist := repository.GetArtistDTO{}
		err := rows.Scan(&artist.Id, &artist.FirstName, &artist.LastName)
		if err != nil {
			return nil, err
		}

		artists = append(artists, artist)
	}

	return artists, nil
}

func (ar *postgresArtistRepository) Insert(ctx context.Context, artist *repository.InsertArtistDTO) (*entity.Artist, error) {
	query := `INSERT INTO artist 
        (first_name, last_name)
        VALUES (@firstName, @lastName)
        RETURNING id;`

	args := pgx.NamedArgs{
		"firstName": artist.FirstName,
		"lastName":  artist.LastName,
	}

	var id int
	err := ar.psql.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		return nil, err
	}

	return entity.NewArtist(id, artist.FirstName, artist.LastName), nil
}

func (ar *postgresArtistRepository) Update(ctx context.Context, artist *entity.Artist) error {
	return errors.New("not implemented")
	query := ``
	args := pgx.NamedArgs{}

	return ar.psql.Exec(ctx, query, args)
}

func (ar *postgresArtistRepository) Delete(ctx context.Context, artist *entity.Artist) error {
	return errors.New("not implemented")
}
