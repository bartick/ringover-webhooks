package models

import (
	"errors"
	"net/url"
)

type WebhookRequestDataAll interface {
	WebhookRequestDataIVRCode |
		WebhookRequestDataMessage |
		WebhookRequestDataSmartRouting |
		WebhookRequestDataContactCall |
		WebhookRequestDataContactSearch |
		WebhookRequestDataCall |
		WebhookRequestDataCommentUpdated |
		WebhookRequestDataTagUpdated |
		WebhookRequestDataRecordAvailable |
		WebhookRequestDataVoicemailAvailable
}

//
// --- Event Requests Data

type WebhookRequestDataIVRCode struct {
	Code       int    `json:"code"`
	NumberFrom string `json:"from_number"`
	NumberTo   string `json:"to_number"`
	Direction  string `json:"direction"`
	CallId     int64  `json:"call_id"`
}

func (e *WebhookRequestDataIVRCode) validate() error {
	if e.Code == 0 {
		return errors.New("code is required")
	}
	if e.NumberFrom == "" {
		return errors.New("from_number is required")
	}
	if e.NumberTo == "" {
		return errors.New("to_number is required")
	}
	if e.Direction == "" {
		return errors.New("direction is required")
	}
	if e.CallId == 0 {
		return errors.New("call_id is required")
	}
	return nil
}

func (e *WebhookRequestDataIVRCode) MatchResourceEvent(resource, event string) bool {
	return resource == "call" && event == "ivr_response_code"
}

type WebhookRequestDataMessage struct {
	Id              string `json:"id"`
	MessageId       int    `json:"message_id"`
	ConversationId  int    `json:"conversation_id"`
	Time            int    `json:"time"`
	Direction       string `json:"direction"`
	NumberFrom      int64  `json:"from_number"`
	NumberTo        int64  `json:"to_number"`
	Body            string `json:"body"`
	IsInternal      bool   `json:"is_internal"`
	IsCollaborative bool   `json:"is_collaborative"`
	UserId          int64  `json:"user_id"`
}

func (e *WebhookRequestDataMessage) validate() error {
	if e.Id == "" {
		return errors.New("id is required")
	}
	if e.MessageId == 0 {
		return errors.New("message_id is required")
	}
	if e.ConversationId == 0 {
		return errors.New("conversation_id is required")
	}
	if e.Direction == "" {
		return errors.New("direction is required")
	}
	if e.NumberFrom == 0 {
		return errors.New("from_number is required")
	}
	if e.NumberTo == 0 {
		return errors.New("to_number is required")
	}
	if e.Body == "" {
		return errors.New("body is required")
	}
	if e.UserId == 0 {
		return errors.New("user_id is required")
	}
	return nil
}

func (e *WebhookRequestDataMessage) MatchResourceEvent(resource, event string) bool {
	return resource == "sms" && event == "sent"
}

type WebhookRequestDataSmartRouting struct {
	CallId     int64  `json:"call_id"`
	Direction  string `json:"direction"`
	NumberFrom int64  `json:"from_number"`
	NumberTo   int64  `json:"to_number"`
}

func (e *WebhookRequestDataSmartRouting) validate() error {
	if e.CallId == 0 {
		return errors.New("call_id is required")
	}
	if e.Direction == "" {
		return errors.New("direction is required")
	}
	if e.NumberFrom == 0 {
		return errors.New("from_number is required")
	}
	if e.NumberTo == 0 {
		return errors.New("to_number is required")
	}
	return nil
}

func (e *WebhookRequestDataSmartRouting) MatchResourceEvent(resource, event string) bool {
	return resource == "call" && event == "routing"
}

type WebhookRequestDataContactCall struct {
	CallId     string `json:"call_id"`
	Direction  string `json:"direction"`
	NumberFrom string `json:"from_number"`
	NumberTo   string `json:"to_number"`
}

func (e *WebhookRequestDataContactCall) validate() error {
	if e.CallId == "" {
		return errors.New("call_id is required")
	}
	if e.Direction == "" {
		return errors.New("direction is required")
	}
	if e.NumberFrom == "" {
		return errors.New("from_number is required")
	}
	if e.NumberTo == "" {
		return errors.New("to_number is required")
	}
	return nil
}

func (e *WebhookRequestDataContactCall) MatchResourceEvent(resource, event string) bool {
	return resource == "call" && event == "contact"
}

type WebhookRequestDataContactSearch struct {
	QuerySearch string `json:"query_search"`
	UserId      int    `json:"user_id"`
}

func (e *WebhookRequestDataContactSearch) validate() error {
	if e.QuerySearch == "" {
		return errors.New("query_search is required")
	}
	if e.UserId == 0 {
		return errors.New("user_id is required")
	}
	return nil
}

func (e *WebhookRequestDataContactSearch) MatchResourceEvent(resource, event string) bool {
	return resource == "search" && event == "contact"
}

type WebhookRequestDataCall struct {
	/*
		Call Events:
		- ringing
		- answered
		- hangup
		- missed
		- voicemail
	*/
	CallEvent string // To differentiate between call events

	Id           string  `json:"id"`
	CallId       string  `json:"call_id"`
	ChannelId    string  `json:"channel_id"`
	StartTime    float64 `json:"start_time"`
	HangupTime   float64 `json:"hangup_time,omitempty"`
	AnsweredTime float64 `json:"answered_time,omitempty"`
	Duration     int     `json:"duration_in_seconds,omitempty"`
	Record       string  `json:"record,omitempty"`
	Direction    string  `json:"direction"`
	NumberFrom   string  `json:"from_number"`
	NumberTo     string  `json:"to_number"`
	UserId       string  `json:"user_id"`
	IsInternal   bool    `json:"is_internal"`
	IsAnonymous  bool    `json:"is_anonymous"`
	IsIVR        bool    `json:"is_ivr"`
	IVRData      IVRData `json:"ivr_data"`
	User         User    `json:"user,omitempty"`
	Status       string  `json:"status,omitempty"`
	Reason       string  `json:"reason,omitempty"`
}

func (e *WebhookRequestDataCall) validate() error {
	if e.Id == "" {
		return errors.New("id is required")
	}
	if e.CallId == "" {
		return errors.New("call_id is required")
	}
	if e.ChannelId == "" {
		return errors.New("channel_id is required")
	}
	if e.StartTime == 0 {
		return errors.New("start_time is required")
	}
	if e.Direction == "" {
		return errors.New("direction is required")
	}
	if e.NumberFrom == "" {
		return errors.New("from_number is required")
	}
	if e.NumberTo == "" {
		return errors.New("to_number is required")
	}
	if e.UserId == "" {
		return errors.New("user_id is required")
	}
	if e.Status == "" {
		return errors.New("status is required")
	}

	return nil
}

func (e *WebhookRequestDataCall) MatchResourceEvent(resource, event string) bool {
	if resource != "call" {
		return false
	}

	switch event {
	case "ringing", "answered", "hangup", "missed", "voicemail":
		return true
	default:
		return false
	}
}

/*
type IVR struct {
	Name     string   `json:"name"`
	Number   int64    `json:"number"`
	Scenario struct{} `json:"scenario"`
}

type CallMissedEvent struct {
	CallEvent
	IVR IVR `json:"ivr"`
}
*/

type WebhookRequestDataCommentUpdated struct {
	CallId    int64    `json:"call_id"`
	ChannelId int64    `json:"channel_id"`
	Tags      []string `json:"tags"`
	Comments  string   `json:"comments"`
}

func (e *WebhookRequestDataCommentUpdated) validate() error {
	if e.CallId == 0 {
		return errors.New("call_id is required")
	}
	if e.ChannelId == 0 {
		return errors.New("channel_id is required")
	}
	if len(e.Tags) == 0 {
		return errors.New("tags cannot be empty")
	}
	if e.Comments == "" {
		return errors.New("comments cannot be empty")
	}
	return nil
}

func (e *WebhookRequestDataCommentUpdated) MatchResourceEvent(resource, event string) bool {
	return resource == "aftercall" && event == "comments_updated"
}

type WebhookRequestDataTagUpdated struct {
	CallId int64    `json:"call_id"`
	Tags   []string `json:"tags"`
}

func (e *WebhookRequestDataTagUpdated) validate() error {
	if e.CallId == 0 {
		return errors.New("call_id is required")
	}
	if len(e.Tags) == 0 {
		return errors.New("tags cannot be empty")
	}
	return nil
}

func (e *WebhookRequestDataTagUpdated) MatchResourceEvent(resource, event string) bool {
	return resource == "aftercall" && event == "tags_updated"
}

type WebhookRequestDataRecordAvailable struct {
	CallId         int64  `json:"call_id"`
	RecordLink     string `json:"record_link"`
	RecordDuration string `json:"record_duration"`
}

func (e *WebhookRequestDataRecordAvailable) validate() error {
	if e.CallId == 0 {
		return errors.New("call_id is required")
	}
	if e.RecordLink == "" {
		return errors.New("record_link is required")
	}
	if _, err := url.ParseRequestURI(e.RecordLink); err != nil {
		return errors.New("record_link must be a valid URL")
	}
	if e.RecordDuration == "" {
		return errors.New("record_duration is required")
	}
	return nil
}

func (e *WebhookRequestDataRecordAvailable) MatchResourceEvent(resource, event string) bool {
	return resource == "aftercall" && event == "record_available"
}

type WebhookRequestDataVoicemailAvailable struct {
	CallId            int64  `json:"call_id"`
	VoicemailLink     string `json:"voicemail_link"`
	VoicemailDuration string `json:"voicemail_duration"`
}

func (e *WebhookRequestDataVoicemailAvailable) validate() error {
	if e.CallId == 0 {
		return errors.New("call_id is required")
	}
	if e.VoicemailLink == "" {
		return errors.New("voicemail_link is required")
	}
	if _, err := url.ParseRequestURI(e.VoicemailLink); err != nil {
		return errors.New("voicemail_link must be a valid URL")
	}
	if e.VoicemailDuration == "" {
		return errors.New("voicemail_duration is required")
	}
	return nil
}

func (e *WebhookRequestDataVoicemailAvailable) MatchResourceEvent(resource, event string) bool {
	return resource == "aftercall" && event == "voicemail_available"
}

//
// --- Event Requests Data Sub-Entities

type IVRData struct {
	Number       string `json:"number"`
	ScenarioName string `json:"scenario_name"`
	IVRName      string `json:"ivr_name"`
}

type User struct {
	UserId    int64  `json:"user_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Photo     string `json:"photo"`
}
