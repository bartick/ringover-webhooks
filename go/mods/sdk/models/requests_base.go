package models

import (
	"encoding/json"
	"fmt"
)

type WebhookRequestType interface {
	validate() error
	MarhsalEventData() (json.RawMessage, error)
	ResourceEvent() string

	//
	GetResource() string
	GetEvent() string
}

type WebhookRequestDataType interface {
	validate() error
	MatchResourceEvent(resource string, event string) bool
}

var _ WebhookRequestDataType = &WebhookRequestDataContactSearch{}

type WebhookRequest[WebhookData WebhookRequestDataType] struct {
	Id        int64
	Resource  string
	Event     string
	timestamp int64
	Data      WebhookData
	Attempt   int64
}

func NewWebhookRequest[WebhookData WebhookRequestDataType](
	event string,
	resource string,
	Timestamp int64,
	data json.RawMessage,
) (
	WebhookRequestType,
	error,
) {
	var err error
	var webhookData WebhookData

	// Validation Pre-Creation
	if webhookData.MatchResourceEvent(resource, event) {
		return nil, fmt.Errorf("failed to build EventRequest: failed to validate EventRequest: EventData does not match input for resource/event: eventData=T(%T) // input=V(resource=%q/event=%q)", webhookData, resource, event)
	}

	err = json.Unmarshal(data, &webhookData)
	if err != nil {
		return nil, fmt.Errorf("failed to build EventRequest: failed to unmarshal webhookData: %w", err)
	}

	ev := &WebhookRequest[WebhookData]{
		Id:        0,
		Event:     event,
		Resource:  resource,
		timestamp: Timestamp,
		Data:      webhookData,
	}

	// Validation Post-Creation
	err = ev.validate()
	if err != nil {
		return nil, fmt.Errorf("failed to build EventRequest: failed to validate EventRequest: %w", err)
	}

	return ev, nil
}

func (ev *WebhookRequest[EventData]) validate() error {
	if ev.Event == "" {
		return fmt.Errorf("invalid EventRequest: Event is required")
	}
	if ev.Resource == "" {
		return fmt.Errorf("invalid EventRequest: Resource is required")
	}
	if err := ev.Data.validate(); err != nil {
		return fmt.Errorf("invalid EventRequest: EventData is invalid: %w", err)
	}
	return nil
}

func (ev *WebhookRequest[EventData]) MarhsalEventData() (json.RawMessage, error) {
	evDataJson, err := json.Marshal(ev.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to build EventRequest: failed to unmarshal webhookData: %w", err)
	}

	return json.RawMessage(evDataJson), nil
}

func (ev *WebhookRequest[EventData]) ResourceEvent() string {
	return fmt.Sprintf("%s/%s", ev.Resource, ev.Event)
}

func (ev *WebhookRequest[EventData]) GetResource() string {
	return ev.Resource
}

func (ev *WebhookRequest[EventData]) GetEvent() string {
	return ev.Event
}
