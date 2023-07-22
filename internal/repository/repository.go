package repository

import (
	"context"

	"github.com/DanielTitkov/weavle/internal/domain"
	"github.com/DanielTitkov/weavle/internal/ent"
	"github.com/DanielTitkov/weavle/internal/ent/story"
	"github.com/DanielTitkov/weavle/internal/logger"
	"github.com/google/uuid"
)

type Repository struct {
	client *ent.Client
	logger *logger.Logger
}

func NewRepository(client *ent.Client, logger *logger.Logger) *Repository {
	return &Repository{
		client: client,
		logger: logger,
	}
}

// CreateStory creates new empty story (without sentences)
func (r *Repository) CreateStory(ctx context.Context, newStory *domain.Story) (*domain.Story, error) {
	dbStory, err := r.client.Story.
		Create().
		SetStatus(story.Status(newStory.Status)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	newStory.ID = dbStory.ID
	return newStory, nil
}

// GetStory gets story with all it's sentences
func (r *Repository) GetStory(ctx context.Context, id uuid.UUID) (*domain.Story, error) {
	dbStory, err := r.client.Story.
		Query().
		Where(story.ID(id)).
		WithSentences().
		First(ctx)

	if err != nil {
		return nil, err
	}

	return entToDomainStory(dbStory), nil
}

// UpdateStore updates any fields of the story (status etc)
func (r *Repository) UpdateStory(ctx context.Context, newStory *domain.Story) (*domain.Story, error) {
	_, err := r.client.Story.
		UpdateOneID(newStory.ID).
		SetStatus(story.Status(newStory.Status)).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return newStory, nil
}

// CreateSentence creates sentence if sentence structs has a non-zero sentence.StoryID.
func (r *Repository) CreateSentence(ctx context.Context, sentence *domain.Sentence) (*domain.Sentence, error) {
	dbSentence, err := r.client.Sentence.
		Create().
		SetStoryID(sentence.StoryID).
		SetText(sentence.Text).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	sentence.ID = dbSentence.ID
	return sentence, nil
}

func entToDomainStory(dbStory *ent.Story) *domain.Story {
	return &domain.Story{
		ID:        dbStory.ID,
		Status:    dbStory.Status.String(),
		Sentences: entToDomainSentences(dbStory.Edges.Sentences),
	}
}

func entToDomainSentence(dbSentence *ent.Sentence) *domain.Sentence {
	return &domain.Sentence{
		ID:   dbSentence.ID,
		Text: dbSentence.Text,
	}
}

func entToDomainSentences(dbSentences []*ent.Sentence) []*domain.Sentence {
	sentences := make([]*domain.Sentence, len(dbSentences))
	for i, dbSentence := range dbSentences {
		sentences[i] = entToDomainSentence(dbSentence)
	}

	return sentences
}
