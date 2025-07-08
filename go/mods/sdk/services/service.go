package services

import (
	"context"
	"fmt"

	"github.com/ringover/ringover-webhooks/adapters/adaptersrepositories"
	"github.com/ringover/ringover-webhooks/mappers"
	"github.com/ringover/ringover-webhooks/models"
)

type WebhookEventService struct {
	repo adaptersrepositories.AdapterRepositoryWebhookEvent
	// Add a responder field to allow injection of custom response logic
	responder models.WebhookEventResponder
}

func NewWebhookEventService(repo adaptersrepositories.AdapterRepositoryWebhookEvent, responder models.WebhookEventResponder) *WebhookEventService {
	return &WebhookEventService{repo: repo, responder: responder}
}

func (svc *WebhookEventService) ProcessWebhookRequest(
	ctx context.Context,
	whReq models.WebhookRequestDTO,
) (
	models.WebhookRequestType,
	error,
) {
	var (
		err error
		evt models.WebhookRequestType
	)

	evt, err = mappers.MapToResourceEvent(whReq)
	if err != nil {
		return nil, fmt.Errorf("failed in ProcessWebhookRequest/MapToResourceEvent: failed to map webhook as input (%s/%s): %w", whReq.Resource, whReq.Event, err)
	}

	evtData, err := evt.MarhsalEventData()
	if err != nil {
		return nil, fmt.Errorf("failed to process webhook: MarhsalEventData(%s): %w", evt.ResourceEvent(), err)
	}

	err = svc.repo.Create(ctx, &adaptersrepositories.CreateWebhookEventParams{
		Resource:  evt.GetResource(),
		Event:     evt.GetEvent(),
		Timestamp: 1746544648,
		Data:      evtData,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to process webhook: repo.Create: %w", err)
	}
	return evt, nil
}

func (svc *WebhookEventService) ProcessWebhookWithResponse(ctx context.Context, req models.WebhookRequestDTO) (any, error) {
	_, err := svc.ProcessWebhookRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if req.Resource == "call" && req.Event == "routing" && svc.responder != nil {
		return svc.responder.SmartRoutingResponse(req)
	} else if req.Resource == "call" && req.Event == "contact" && svc.responder != nil {
		return svc.responder.ContactCallResponse(req)
	} else if req.Resource == "contact" && req.Event == "search" && svc.responder != nil {
		return svc.responder.ContactSearchResponse(req)
	}
	return nil, nil
}
