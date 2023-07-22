package repository

import (
	"context"

	"github.com/DanielTitkov/weavle/domain"
	"github.com/DanielTitkov/weavle/ent"
	"github.com/DanielTitkov/weavle/ent/user"
	"github.com/google/uuid"
)

func (r *Repository) CreateUser(ctx context.Context, ip string) (*domain.User, error) {
	user, err := r.client.User.
		Create().
		SetIPAddress(ip).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func (r *Repository) GetUser(ctx context.Context, cookieUUID uuid.UUID, ip string) (*domain.User, error) {
	user, err := r.client.User.
		Query().
		Where(
			user.ID(cookieUUID),
			user.IPAddress(ip),
		).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return entToDomainUser(user), nil
}

func entToDomainUser(user *ent.User) *domain.User {
	return &domain.User{
		ID:        user.ID,
		IPAddress: user.IPAddress,
	}
}
