package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNewWebhookRequest_WebhookRequestDataMessage(t *testing.T) {

	type args struct {
		event     string
		resource  string
		timestamp int64
		data      json.RawMessage
	}
	tests := []struct {
		name    string
		args    args
		want    *WebhookRequest[*WebhookRequestDataMessage]
		wantErr bool
	}{
		// {
		// 	name: "Nominal",
		// 	args: args{
		// 		resource:  "sms",
		// 		event:     "sent",
		// 		timestamp: tms,
		// 		data: json.RawMessage(
		// 			`{
		// 				"id": "1234567-234567",
		// 				"message_id": 1234567,
		// 				"conversation_id": 234567,
		// 				"time": 0,
		// 				"direction": "inbound",
		// 				"from_number": 33123456789,
		// 				"to_number": 33123456789,
		// 				"body": "My SMS content...",
		// 				"is_internal": false,
		// 				"is_collaborative": false,
		// 				"user_id": 1234567890987654400
		// 			}`,
		// 		),
		// 	},
		// 	want: &WebhookRequest[*WebhookRequestDataMessage]{
		// 		Id:        0,
		// 		Resource:  "sms",
		// 		Event:     "sent",
		// 		Timestamp: tms,
		// 		Data: &WebhookRequestDataMessage{
		// 			Id:              "1234567-234567",
		// 			MessageId:       1234567,
		// 			ConversationId:  234567,
		// 			Time:            0,
		// 			Direction:       "inbound",
		// 			NumberFrom:      33123456789,
		// 			NumberTo:        33123456789,
		// 			Body:            "My SMS content...",
		// 			IsInternal:      false,
		// 			IsCollaborative: false,
		// 			UserId:          1234567890987654400,
		// 		},
		// 	},
		// 	wantErr: false,
		// },
		/*
			{
				name: "TPL",
				args: args{
					resource:  "sms",
					event:     "sent",
					timestamp: tms,
					data:      json.RawMessage{},
				},
				want: &WebhookRequest[*WebhookRequestDataMessage]{
					Id:        0,
					Resource:  "sms",
					Event:     "sent",
					Timestamp: tms,
					Data:      &WebhookRequestDataMessage{},
				},
				wantErr: true,
			},
		*/
		{
			name: "Empty",
			args: args{
				resource:  "sms",
				event:     "sent",
				timestamp: 1746544648,
				data:      json.RawMessage{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWebhookRequest[*WebhookRequestDataMessage](tt.args.event, tt.args.resource, tt.args.timestamp, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWebhookRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.want == nil && got != nil) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWebhookRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
