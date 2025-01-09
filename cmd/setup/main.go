package main

import (
	"context"
	"random-char/sql-comparison/pkg/db"
)

func main() {
	psql, err := db.New(context.Background())
	if err != nil {
		panic(err)
	}

	psql.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS artist (
            id INT GENERATED ALWAYS AS IDENTITY,
            first_name VARCHAR(255) NOT NULL,
            last_name VARCHAR(255) NOT NULL,
            PRIMARY KEY(id)
        )`,
	)
    psql.Exec(
        context.Background(),
        `CREATE TABLE IF NOT EXISTS album (
            id INT GENERATED ALWAYS AS IDENTITY,
            artist_id INT NOT NULL,
            name VARCHAR(255) NOT NULL,
            date TIMESTAMP WITHOUT TIMEZONE NOT NULL,

            PRIMARY KEY(id),
            CONSTRAINT fk_artist_id FOREIGN KEY(artist_id) REFERENCES artist(id)
        )`,
    )
	//create table
}
