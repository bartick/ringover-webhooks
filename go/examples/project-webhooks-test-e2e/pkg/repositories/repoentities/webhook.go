package repoentities

import (
	"context"
	"encoding/json"

	"github.com/ringover/ringover-webhooks/models"
)

type CreateWebhookEventParams struct {
	Resource  string
	Event     string
	Timestamp int64
	Data      json.RawMessage
}

type RepositoryWebhookEvent interface {
	Create(ctx context.Context, arg *CreateWebhookEventParams) error
	Delete(ctx context.Context, id int64) error
	GetOneById(ctx context.Context, id int64) (models.WebhookRequestType, error)
	GetAll(ctx context.Context) ([]models.WebhookRequestType, error)
}
