package models

import "testing"

func TestWebhookRequestDataCommentUpdated_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataCommentUpdated
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataCommentUpdated{
				CallId:    1234567890987654400,
				ChannelId: 123456789,
				Tags:      []string{"tag1", "tag2"},
				Comments:  "Valid comment",
			},
			wantErr: false,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataCommentUpdated{
				CallId:    0,
				ChannelId: 123456789,
				Tags:      []string{"tag1", "tag2"},
				Comments:  "Valid comment",
			},
			wantErr: true,
		},
		{
			name: "Empty Tags",
			data: WebhookRequestDataCommentUpdated{
				CallId:    1234567890987654400,
				ChannelId: 123456789,
				Tags:      []string{},
				Comments:  "Valid comment",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataTagUpdated_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataTagUpdated
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataTagUpdated{
				CallId: 1234567890987654400,
				Tags:   []string{"tag1"},
			},
			wantErr: false,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataTagUpdated{
				CallId: 0,
				Tags:   []string{"tag1"},
			},
			wantErr: true,
		},
		{
			name: "Empty Tags",
			data: WebhookRequestDataTagUpdated{
				CallId: 1234567890987654400,
				Tags:   []string{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataRecordAvailable_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataRecordAvailable
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataRecordAvailable{
				CallId:         1234567890987654400,
				RecordLink:     "https://example.com/record.mp3",
				RecordDuration: "30s",
			},
			wantErr: false,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataRecordAvailable{
				CallId:         0,
				RecordLink:     "https://example.com/record.mp3",
				RecordDuration: "30s",
			},
			wantErr: true,
		},
		{
			name: "Missing RecordLink",
			data: WebhookRequestDataRecordAvailable{
				CallId:         1234567890987654400,
				RecordLink:     "",
				RecordDuration: "30s",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataVoicemailAvailable_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataVoicemailAvailable
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataVoicemailAvailable{
				CallId:            1234567890987654400,
				VoicemailLink:     "https://example.com/voicemail.mp3",
				VoicemailDuration: "30s",
			},
			wantErr: false,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataVoicemailAvailable{
				CallId:            0,
				VoicemailLink:     "https://example.com/voicemail.mp3",
				VoicemailDuration: "30s",
			},
			wantErr: true,
		},
		{
			name: "Missing VoicemailLink",
			data: WebhookRequestDataVoicemailAvailable{
				CallId:            1234567890987654400,
				VoicemailLink:     "",
				VoicemailDuration: "30s",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataCall_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataCall
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataCall{
				Id:          "57622f20-a020-4f9e-814c-62a6123412fa",
				CallId:      "1234567890987654400",
				ChannelId:   "123456789",
				StartTime:   1554823493.762305,
				Direction:   "inbound",
				NumberFrom:  "33123456789",
				NumberTo:    "33123456789",
				UserId:      "12345678",
				IsInternal:  true,
				IsAnonymous: false,
				IsIVR:       true,
				IVRData:     IVRData{},
				User:        User{UserId: 123456789, FirstName: "Jean", LastName: "Dupont", Email: "jean.dupont@ringover.com", Photo: "https://cdn.ringover.com/img/users/default.jpg"},
				Status:      "ringing",
			},
			wantErr: false,
		},
		{
			name: "Missing Id",
			data: WebhookRequestDataCall{
				CallId:      "1234567890987654400",
				ChannelId:   "123456789",
				StartTime:   1554823493.762305,
				Direction:   "inbound",
				NumberFrom:  "33123456789",
				NumberTo:    "33123456789",
				UserId:      "12345678",
				IsInternal:  true,
				IsAnonymous: false,
				IsIVR:       true,
				IVRData:     IVRData{},
				User:        User{UserId: 123456789, FirstName: "Jean", LastName: "Dupont", Email: "jean.dupont@ringover.com", Photo: "https://cdn.ringover.com/img/users/default.jpg"},
				Status:      "ringing",
			},
			wantErr: true,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataCall{
				Id:          "57622f20-a020-4f9e-814c-62a6123412fa",
				ChannelId:   "123456789",
				StartTime:   1554823493.762305,
				Direction:   "inbound",
				NumberFrom:  "33123456789",
				NumberTo:    "33123456789",
				UserId:      "12345678",
				IsInternal:  true,
				IsAnonymous: false,
				IsIVR:       true,
				IVRData:     IVRData{},
				User:        User{UserId: 123456789, FirstName: "Jean", LastName: "Dupont", Email: "jean.dupont@ringover.com", Photo: "https://cdn.ringover.com/img/users/default.jpg"},
				Status:      "ringing",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataMessage_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataMessage
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataMessage{
				Id:             "1234567-234567",
				MessageId:      1234567,
				ConversationId: 234567,
				Direction:      "inbound",
				NumberFrom:     33123456789,
				NumberTo:       33123456789,
				Body:           "My SMS content...",
				UserId:         1234567890987654400,
			},
			wantErr: false,
		},
		{
			name: "Missing Id",
			data: WebhookRequestDataMessage{
				MessageId:      1234567,
				ConversationId: 234567,
				Direction:      "inbound",
				NumberFrom:     33123456789,
				NumberTo:       33123456789,
				Body:           "My SMS content...",
				UserId:         1234567890987654400,
			},
			wantErr: true,
		},
		// Add more invalid cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataSmartRouting_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataSmartRouting
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataSmartRouting{
				CallId:     1234567890987654400,
				Direction:  "inbound",
				NumberFrom: 33123456789,
				NumberTo:   33123456789,
			},
			wantErr: false,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataSmartRouting{
				Direction:  "inbound",
				NumberFrom: 33123456789,
				NumberTo:   33123456789,
			},
			wantErr: true,
		},
		// Add more invalid cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataContactCall_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataContactCall
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataContactCall{
				CallId:     "1234567890987654400",
				Direction:  "inbound",
				NumberFrom: "33123456789",
				NumberTo:   "33123456789",
			},
			wantErr: false,
		},
		{
			name: "Missing CallId",
			data: WebhookRequestDataContactCall{
				Direction:  "inbound",
				NumberFrom: "33123456789",
				NumberTo:   "33123456789",
			},
			wantErr: true,
		},
		{
			name: "Missing NumberFrom",
			data: WebhookRequestDataContactCall{
				CallId:    "1234567890987654400",
				Direction: "inbound",
				NumberTo:  "33123456789",
			},
			wantErr: true,
		},
		{
			name: "Missing NumberTo",
			data: WebhookRequestDataContactCall{
				CallId:     "1234567890987654400",
				Direction:  "inbound",
				NumberFrom: "33123456789",
			},
			wantErr: true,
		},
		{
			name: "Missing Direction",
			data: WebhookRequestDataContactCall{
				CallId:     "1234567890987654400",
				NumberFrom: "33123456789",
				NumberTo:   "33123456789",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataContactSearch_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataContactSearch
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataContactSearch{
				QuerySearch: "search query",
				UserId:      123456,
			},
			wantErr: false,
		},
		{
			name: "Missing QuerySearch",
			data: WebhookRequestDataContactSearch{
				UserId: 123456,
			},
			wantErr: true,
		},
		// Add more invalid cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWebhookRequestDataIVRCode_validate(t *testing.T) {
	tests := []struct {
		name    string
		data    WebhookRequestDataIVRCode
		wantErr bool
	}{
		{
			name: "Valid Data",
			data: WebhookRequestDataIVRCode{
				Code:       1234,
				NumberFrom: "33123456789",
				NumberTo:   "33123456789",
				Direction:  "inbound",
				CallId:     1234567890987654400,
			},
			wantErr: false,
		},
		{
			name: "Missing Code",
			data: WebhookRequestDataIVRCode{
				NumberFrom: "33123456789",
				NumberTo:   "33123456789",
				Direction:  "inbound",
				CallId:     1234567890987654400,
			},
			wantErr: true,
		},
		// Add more invalid cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.validate(); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
