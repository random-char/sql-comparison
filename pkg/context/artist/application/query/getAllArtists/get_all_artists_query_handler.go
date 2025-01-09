package getallartists

import (
	"context"
	"random-char/sql-comparison/pkg/context/artist/domain/repository"
)

type handler struct {
	repo repository.ArtistRepository
}

func NewHandler(repo repository.ArtistRepository) *handler {
	return &handler{
		repo: repo,
	}
}

func (h *handler) Handle(ctx context.Context, query Query) ([]repository.GetArtistDTO, error) {
	return h.repo.FindAll(ctx)
}
