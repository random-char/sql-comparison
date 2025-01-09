package main

import (
	"context"
	"fmt"
	getallartists "random-char/sql-comparison/pkg/context/artist/application/query/getAllArtists"
	//artistRepo "random-char/sql-comparison/pkg/context/artist/domain/repository"
	postgresArtistRepo "random-char/sql-comparison/pkg/context/artist/infrastructure/repository"
	"random-char/sql-comparison/pkg/db"
)

func main() {
	psql, err := db.New(context.Background())
	if err != nil {
		panic(err)
	}
	defer psql.Close()

	repo := postgresArtistRepo.New(psql)

	//artist1 := artistRepo.InsertArtistDTO{
	//	FirstName: "D",
	//	LastName:  "B",
	//}
	//a1, err := repo.Insert(context.Background(), &artist1)

	artists, err := getallartists.
		NewHandler(repo).
		Handle(context.Background(), getallartists.Query{})

	fmt.Printf("all: %+v\n", artists)
	fmt.Println(err)
}
