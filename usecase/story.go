package usecase

import (
	"context"

	"github.com/DanielTitkov/weavle/domain"
	"github.com/DanielTitkov/weavle/repository"
)

type StoryUsecase struct {
	repo repository.Repository
}

func NewStoryUsecase(repo repository.Repository) *StoryUsecase {
	return &StoryUsecase{repo: repo}
}

func (u *StoryUsecase) CreateStory(ctx context.Context) (*domain.Story, error) {

	story := &domain.Story{
		Status: "open",
	}

	story, err := u.repo.CreateStory(ctx, story)
	if err != nil {
		return nil, err
	}

	return story, nil
}
