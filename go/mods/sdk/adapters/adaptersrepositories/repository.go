package adaptersrepositories

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

type AdapterRepositoryWebhookEvent interface {
	Create(ctx context.Context, arg *CreateWebhookEventParams) error
	Delete(ctx context.Context, id int64) error
	GetOneById(ctx context.Context, id int64) (models.WebhookRequestType, error) // Replace `interface{}` with a proper type if needed
	GetAll(ctx context.Context) ([]models.WebhookRequestType, error)
}
