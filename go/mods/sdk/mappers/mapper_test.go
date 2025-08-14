package mappers

import (
	"encoding/json"
	"testing"

	"github.com/ringover/ringover-webhooks/models"
)

func TestMapToResourceEvent(t *testing.T) {
	tests := []struct {
		name    string
		payload models.WebhookRequestDTO
		wantErr bool
	}{
		{
			name: "Valid comments_updated Event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "comments_updated",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"tags": ["tag1", "tag2"],
					"comments": "Valid comment"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid tags_updated Event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "tags_updated",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"tags": ["tag1"]
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid record_available Event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "record_available",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"record_link": "https://example.com/record.mp3",
					"record_duration": "30s"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid voicemail_available Event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "voicemail_available",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"voicemail_link": "https://example.com/voicemail.mp3",
					"voicemail_duration": "30s"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Invalid comments_updated Event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "comments_updated",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 0,
					"channel_id": 123456789,
					"tags": ["tag1", "tag2"],
					"comments": "Valid comment"
				}`),
			},
			wantErr: true,
		},
		{
			name: "Valid ringing Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "ringing",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "ringing"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid answered Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "answered",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "answered"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid hangup Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "hangup",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"hangup_time": 1554823442.762305,
					"duration_in_seconds": 32,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "hangup"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid missed Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "missed",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "missed"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid voicemail Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "voicemail",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "voicemail"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Invalid missing id",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "ringing",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "ringing"
				}`),
			},
			wantErr: true,
		},
		{
			name: "Invalid missing call_id",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "answered",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "answered"
				}`),
			},
			wantErr: true,
		},
		{
			name: "Invalid missing start_time",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "hangup",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					},
					"status": "hangup"
				}`),
			},
			wantErr: true,
		},
		{
			name: "Invalid missing status",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "voicemail",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "57622f20-a020-4f9e-814c-62a6123412fa",
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"start_time": 1554823493.762305,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789",
					"user_id": 12345678,
					"is_internal": 1,
					"is_anonymous": 0,
					"is_ivr": 1,
					"ivr_data": {},
					"user": {
						"user_id": 123456789,
						"firstname": "Jean",
						"lastname": "Dupont",
						"email": "jean.dupont@ringover.com",
						"photo": "https://cdn.ringover.com/img/users/default.jpg"
					}
				}`),
			},
			wantErr: true,
		},
		{
			name: "Valid sent Event",
			payload: models.WebhookRequestDTO{
				Resource:  "sms",
				Event:     "sent",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "1234567-234567",
					"message_id": 1234567,
					"conversation_id": 234567,
					"time": 1554823493,
					"direction": "inbound",
					"from_number": 33123456789,
					"to_number": 33123456789,
					"body": "My SMS content...",
					"is_internal": true,
					"is_collaborative": false,
					"user_id": 1234567890987654400
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid ivr_response_code Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "ivr_response_code",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"code": 1234,
					"from_number": "33123456789",
					"to_number": "33123456789",
					"direction": "inbound",
					"call_id": 1234567890987654400
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid contact Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "contact",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"direction": "inbound",
					"from_number": "33123456789",
					"to_number": "33123456789"
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid contact search Event",
			payload: models.WebhookRequestDTO{
				Resource:  "search",
				Event:     "contact",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"query_search": "search query",
					"user_id": 123456
				}`),
			},
			wantErr: false,
		},
		{
			name: "Valid smart routing Event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "routing",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"direction": "inbound",
					"from_number": 33123456789,
					"to_number": 33123456789
				}`),
			},
			wantErr: false,
		},

		{
			name: "Invalid Event",
			payload: models.WebhookRequestDTO{
				Resource:  "invalid_resource",
				Event:     "invalid_event",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,	
					"direction": "inbound",
					"from_number": 33123456789,
					"to_number": 33123456789
				}`),
			},
			wantErr: true,
		},
		{
			name: "Search Resource with invalid event",
			payload: models.WebhookRequestDTO{
				Resource:  "search",
				Event:     "invalid_event",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"query_search": "search query",
					"user_id": 123456
				}`),
			},
			wantErr: true,
		},
		{
			name: "SMS resource with invalid event",
			payload: models.WebhookRequestDTO{
				Resource:  "sms",
				Event:     "invalid_event",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"id": "1234567-234567",
					"message_id": 1234567,
					"conversation_id": 234567,
					"time": 1554823493,
					"direction": "inbound",
					"from_number": 33123456789,
					"to_number": 33123456789,
					"body": "My SMS content...",
					"is_internal": true,
					"is_collaborative": false,
					"user_id": 1234567890987654400
				}`),
			},
			wantErr: true,
		},
		{
			name: "Call resource with invalid event",
			payload: models.WebhookRequestDTO{
				Resource:  "call",
				Event:     "invalid_event",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"direction": "inbound",
					"from_number": 33123456789,
					"to_number": 33123456789
				}`),
			},
			wantErr: true,
		},
		{
			name: "Aftercall resource with invalid event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "invalid_event",
				Timestamp: 1746544648,
				Data: json.RawMessage(`{
					"call_id": 1234567890987654400,
					"channel_id": 123456789,
					"tags": ["tag1", "tag2"],
					"comments": "Valid comment"
				}`),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MapToResourceEvent(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToResourceEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMapToResourceEvent_InvalidResource(t *testing.T) {
	tests := []struct {
		name    string
		payload models.WebhookRequestDTO
		wantErr bool
	}{
		{
			name: "Invalid Resource",
			payload: models.WebhookRequestDTO{
				Resource:  "invalid_resource",
				Event:     "comments_updated",
				Timestamp: 1746544648,
				Data:      json.RawMessage(`{}`),
			},
			wantErr: true,
		},
		{
			name: "Invalid Event",
			payload: models.WebhookRequestDTO{
				Resource:  "aftercall",
				Event:     "invalid_event",
				Timestamp: 1746544648,
				Data:      json.RawMessage(`{}`),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := MapToResourceEvent(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapToResourceEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
