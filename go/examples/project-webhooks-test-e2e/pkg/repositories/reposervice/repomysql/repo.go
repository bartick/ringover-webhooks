package repomysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ringover/ringover-webhooks/adapters/adaptersrepositories"
	"github.com/ringover/ringover-webhooks/mappers"
	"github.com/ringover/ringover-webhooks/models"

	"github.com/ringover/ringover-webhooks/examples/project-webhooks-test-e2e/pkg/repositories/persistence/persmysql/persmysql_sqlc_gen"
)

// Ensure the implementation satisfies the SDK's repository interface
var _ adaptersrepositories.AdapterRepositoryWebhookEvent = &RepositoryWebhookEventMysql{}

type RepositoryWebhookEventMysql struct {
	queries *persmysql_sqlc_gen.Queries
}

func NewRepositoryWebhookEventMysql(db *sql.DB) *RepositoryWebhookEventMysql {
	return &RepositoryWebhookEventMysql{
		queries: persmysql_sqlc_gen.New(db),
	}
}

// Update the Create method to use the SDK's CreateWebhookEventParams
func (r *RepositoryWebhookEventMysql) Create(ctx context.Context, webhook *adaptersrepositories.CreateWebhookEventParams) error {
	if webhook == nil {
		return errors.New("event cannot be nil")
	}
	err := r.queries.CreateWebhookEvent(ctx, persmysql_sqlc_gen.CreateWebhookEventParams{
		Resource:  webhook.Resource,
		Event:     webhook.Event,
		Timestamp: time.Unix(webhook.Timestamp, 0).UTC(),
		Data:      webhook.Data,
	})
	if err != nil {
		return fmt.Errorf("failed in repo.Create/CreateWebhookEvent: %w", err)
	}
	return nil
}

func (r *RepositoryWebhookEventMysql) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteWebhookEvent(ctx, id)
	if err != nil {
		return fmt.Errorf("failed in repo.Delete/DeleteWebhookEvent: %w", err)
	}
	return nil
}

func (r *RepositoryWebhookEventMysql) GetOneById(ctx context.Context, id int64) (models.WebhookRequestType, error) {
	var (
		err error
		evt models.WebhookRequestType
	)
	eventDB, err := r.queries.GetWebhookEventOneById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed in repo.GetOneById/GetWebhookEventById: %w", err)
	}

	evt, err = mappers.MapToResourceEvent(models.WebhookRequestDTO{
		Id:        eventDB.ID,
		Event:     eventDB.Event,
		Resource:  eventDB.Resource,
		Timestamp: eventDB.Timestamp.Unix(),
		Data:      eventDB.Data,
	})
	if err != nil {
		return nil, fmt.Errorf("failed in repo.GetOneById/MapToResourceEvent: failed to map webhook event as output (%s/%s): %w", eventDB.Resource, eventDB.Event, err)
	}

	return evt, nil
}

func (r *RepositoryWebhookEventMysql) GetAllByResourceEvent(ctx context.Context, resource string, event string) ([]models.WebhookRequestType, error) {
	eventsDB, err := r.queries.GetWebhookEventAllByResourceEvent(ctx, persmysql_sqlc_gen.GetWebhookEventAllByResourceEventParams{
		Resource: resource,
		Event:    event,
	})
	if err != nil {
		return nil, fmt.Errorf("failed in repo.GetAllByResourceEvent/GetWebhookEventAllByResourceEvent: %w", err)
	}

	events := make([]models.WebhookRequestType, len(eventsDB))
	for _, eventDB := range eventsDB {
		var (
			err error
			evt models.WebhookRequestType
		)
		evt, err = mappers.MapToResourceEvent(models.WebhookRequestDTO{
			Id:        eventDB.ID,
			Event:     eventDB.Event,
			Resource:  eventDB.Resource,
			Timestamp: eventDB.Timestamp.Unix(),
			Data:      eventDB.Data,
		})
		if err != nil {
			return nil, fmt.Errorf("failed in repo.GetAllByResourceEvent/MapToResourceEvent: failed to map webhook event as output (%s/%s): %w", eventDB.Resource, eventDB.Event, err)
		}
		events = append(events, evt)
	}

	return events, nil
}

func (r *RepositoryWebhookEventMysql) GetAll(ctx context.Context) ([]models.WebhookRequestType, error) {
	eventsDB, err := r.queries.GetWebhookEventAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed in repo.GetAll/GetWebhookEventAll: %w", err)
	}

	events := make([]models.WebhookRequestType, len(eventsDB))
	for _, eventDB := range eventsDB {
		var (
			err error
			evt models.WebhookRequestType
		)
		evt, err = mappers.MapToResourceEvent(models.WebhookRequestDTO{
			Id:        eventDB.ID,
			Event:     eventDB.Event,
			Resource:  eventDB.Resource,
			Timestamp: eventDB.Timestamp.Unix(),
			Data:      eventDB.Data,
		})
		if err != nil {
			return nil, fmt.Errorf("failed in repo.GetAll/MapToResourceEvent: failed to map webhook event as output (%s/%s): %w", eventDB.Resource, eventDB.Event, err)
		}
		events = append(events, evt)
	}

	return events, nil
}
