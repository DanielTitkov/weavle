package usecase

import "github.com/DanielTitkov/weavle/repository"

type SentenceUsecase struct {
	repo repository.Repository
}

func NewSentenceUsecase(repo repository.Repository) *SentenceUsecase {
	return &SentenceUsecase{repo: repo}
}
