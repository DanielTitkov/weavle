package prepare

import (
	"context"

	"github.com/DanielTitkov/weavle/ent"
	"github.com/DanielTitkov/weavle/ent/migrate"
)

func Migrate(ctx context.Context, client *ent.Client) error {
	err := client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		return err
	}
	return nil
}
