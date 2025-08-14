package httpcontrollers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ringover/ringover-webhooks/examples/project-webhooks-test-e2e/pkg/repositories/reposervice/repomysql"
	modelswebhooks "github.com/ringover/ringover-webhooks/models"
	servicewebhooks "github.com/ringover/ringover-webhooks/services"
)

/*
TODO:
- handlers for controllers
- HMAC signatue


*/

var svc *servicewebhooks.WebhookEventService

// responderImpl implements models.WebhookEventResponder for the example project
// You can add your own logic here to build real responses

type responderImpl struct{}

func (r *responderImpl) SmartRoutingResponse(req modelswebhooks.WebhookRequestDTO) (*modelswebhooks.EventResponseSmartRouting, error) {
	// Replace with real logic as needed based on the request, DB etc.
	return &modelswebhooks.EventResponseSmartRouting{
		Name:                 "redirections",
		Dispatch:             "ringall",
		MaxAttempts:          1,
		StartDelay:           0,
		IsStayNotConnected:   1,
		IsStayInCall:         0,
		IsStayPlannedSnoozed: 1,
		IsStaySnoozed:        0,
		RingOverlap:          0,
		Agents: []modelswebhooks.Agent{
			{
				AgentType:    "agent_external",
				RingDuration: 25,
				RingDelay:    0,
				Order:        1,
				Number:       33123456789,
				IsPreAnswer:  0,
				IsCallerID:   1,
				IsHeadLine:   0,
			},
		},
	}, nil
}

func (r *responderImpl) ContactCallResponse(req modelswebhooks.WebhookRequestDTO) (*modelswebhooks.EventResponseContactCall, error) {
	// Replace with real logic as needed
	return &modelswebhooks.EventResponseContactCall{
		UUID:      "b55f949b-c49b-4354-b10a-c8c4cdbd8690",
		Firstname: "Jean-Pierre",
		Lastname:  "De La Court",
		Company:   "Ringover",
		URL:       "https://mycrm.com/client/18192233",
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
			"keyN": "valueN",
		},
		IsShared: true,
	}, nil
}
func (r *responderImpl) ContactSearchResponse(req modelswebhooks.WebhookRequestDTO) (*modelswebhooks.EventResponseContactSearch, error) {

	return &modelswebhooks.EventResponseContactSearch{
		{
			Firstname: "Jean-Pierre",
			Lastname:  "De La Court",
			Company:   "Ringover",
			URL:       "https://mycrm.com/client/18192233",
			Numbers: []modelswebhooks.Number{
				{
					Number: 33123456789,
					Type:   "mobile",
				},
			},
		},
	}, nil
}

func init() {
	dbURL, dbURLfound := os.LookupEnv("WEBHOOKS_DB_URL")
	if !dbURLfound {
		slog.Error("DB URL missing from environment variable WEBHOOKS_DB_URL")
		os.Exit(1)
	}
	//db, err := sql.Open("mysql", "ringoverwebhooks:ringoverwebhooks@localhost:3306/ringoverwebhooks")
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		slog.Error("DB Connection error", slog.String("error", err.Error()))
		os.Exit(1)
	}
	svc = servicewebhooks.NewWebhookEventService(repomysql.NewRepositoryWebhookEventMysql(db), &responderImpl{})
}

// Export response types for huma registration

type EventResponseSmartRouting struct {
	Body modelswebhooks.EventResponseSmartRouting
}

type EventResponseContactCall struct {
	Body modelswebhooks.EventResponseContactCall
}

type EventResponseContactSearch struct {
	Body modelswebhooks.EventResponseContactSearch
}

// Add authentication middleware to all routes using JWT signature except for search

type WebhookRequest struct {
	// Authorization string `header:"Authorization" required:"true"`
	Body struct {
		Resource  string          `json:"resource" enum:"search,contact,sms,call,aftercall" required:"true" nullable:"false"`
		Event     string          `json:"event" required:"true" nullable:"false"`
		Timestamp float64         `json:"timestamp" required:"true" nullable:"false"`
		Data      json.RawMessage `json:"data" required:"true"`
		Attempt   int64           `json:"attempt" required:"false"`
	}
}

type WebhookResponse struct {
}

func PushWebhook(ctx context.Context, req *WebhookRequest) (rsp *WebhookResponse, err error) {
	_, err = svc.ProcessWebhookWithResponse(ctx, modelswebhooks.WebhookRequestDTO{
		Resource:  req.Body.Resource,
		Event:     req.Body.Event,
		Timestamp: int64(req.Body.Timestamp),
		Data:      req.Body.Data,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to handle webhook event: %w", err)
	}

	// For all other events, return no content
	return nil, nil
}

// Smart Routing controller
func PushWebhookSmartRouting(ctx context.Context, req *WebhookRequest) (*EventResponseSmartRouting, error) {
	resp, err := svc.ProcessWebhookWithResponse(ctx, modelswebhooks.WebhookRequestDTO{
		Resource:  req.Body.Resource,
		Event:     req.Body.Event,
		Timestamp: int64(req.Body.Timestamp), // Cast float64 to int64
		Data:      req.Body.Data,
	})
	// Print response

	if err != nil {
		slog.Error("failed to handle webhook event", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to handle webhook event: %w", err)
	}
	if v, ok := resp.(*modelswebhooks.EventResponseSmartRouting); ok {
		return &EventResponseSmartRouting{Body: *v}, nil
	}
	return nil, nil
}

// Contact Call controller
func PushWebhookContactCall(ctx context.Context, req *WebhookRequest) (*EventResponseContactCall, error) {
	resp, err := svc.ProcessWebhookWithResponse(ctx, modelswebhooks.WebhookRequestDTO{
		Resource:  req.Body.Resource,
		Event:     req.Body.Event,
		Timestamp: int64(req.Body.Timestamp), // Cast float64 to int64
		Data:      req.Body.Data,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to handle webhook event: %w", err)
	}
	if v, ok := resp.(*modelswebhooks.EventResponseContactCall); ok {
		return &EventResponseContactCall{Body: *v}, nil
	}

	return nil, nil
}

// Contact Search controller
func PushWebhookContactSearch(ctx context.Context, req *WebhookRequest) (*EventResponseContactSearch, error) {
	resp, err := svc.ProcessWebhookWithResponse(ctx, modelswebhooks.WebhookRequestDTO{
		Resource:  req.Body.Resource,
		Event:     req.Body.Event,
		Timestamp: int64(req.Body.Timestamp), // Cast float64 to int64
		Data:      req.Body.Data,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to handle webhook event: %w", err)
	}
	if v, ok := resp.(*modelswebhooks.EventResponseContactSearch); ok {
		return &EventResponseContactSearch{Body: *v}, nil
	}
	fmt.Print("not ok")
	return nil, nil
}
