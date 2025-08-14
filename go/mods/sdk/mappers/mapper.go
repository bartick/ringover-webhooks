package mappers

import (
	"fmt"

	"github.com/ringover/ringover-webhooks/models"
)

func MapToResourceEvent(evtReq models.WebhookRequestDTO) (models.WebhookRequestType, error) {

	var (
		err error
		evt models.WebhookRequestType
	)

	// TODO: switch over all evtReq.Resource, evtReq.Event

	switch evtReq.Resource {
	case "search":
		switch evtReq.Event {
		case "contact":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataContactSearch](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		default:
			return nil, fmt.Errorf("failed to process webhook: invalid event: %q", evtReq.Event)
		}
	case "sms":
		switch evtReq.Event {
		case "sent":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataMessage](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		default:
			return nil, fmt.Errorf("failed to process webhook: invalid event: %q", evtReq.Event)
		}
	// TOFIX: !!!
	case "call":
		switch evtReq.Event {
		case "contact":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataContactCall](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "routing":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataSmartRouting](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "ringing":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataCall](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "answered":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataCall](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "hangup":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataCall](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "missed":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataCall](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "voicemail":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataCall](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "ivr_response_code":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataIVRCode](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		default:
			return nil, fmt.Errorf("failed to process webhook: invalid event: %q", evtReq.Event)
		}
	case "aftercall":
		switch evtReq.Event {
		case "tags_updated":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataTagUpdated](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "record_available":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataRecordAvailable](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "voicemail_available":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataVoicemailAvailable](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		case "comments_updated":
			evt, err = models.NewWebhookRequest[*models.WebhookRequestDataCommentUpdated](evtReq.Resource, evtReq.Event, evtReq.Timestamp, evtReq.Data)
		default:
			return nil, fmt.Errorf("failed to process webhook: invalid event: %q", evtReq.Event)
		}
	default:
		return nil, fmt.Errorf("failed to process webhook: invalid resource: %q", evtReq.Resource)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to process webhook: NewWebhookRequest: %w", err)
	}

	return evt, nil
}
