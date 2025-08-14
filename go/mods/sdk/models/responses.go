package models

//
// --- Event Responses

// SmartRoutingEvent represents the smart routing event response.
type EventResponseSmartRouting struct {
	Name                 string  `json:"name"`
	Dispatch             string  `json:"dispatch"`
	MaxAttempts          int     `json:"max_attempts"`
	StartDelay           int     `json:"start_delay"`
	IsStayNotConnected   int     `json:"is_stay_not_connected"`
	IsStayInCall         int     `json:"is_stay_in_call"`
	IsStayPlannedSnoozed int     `json:"is_stay_planned_snoozed"`
	IsStaySnoozed        int     `json:"is_stay_snoozed"`
	RingOverlap          int     `json:"ring_overlap"`
	Agents               []Agent `json:"agents"`
}

// ContactCallEvent represents the contact call event response.
type EventResponseContactCall struct {
	UUID      string            `json:"uuid"`
	Firstname string            `json:"firstname"`
	Lastname  string            `json:"lastname"`
	Company   string            `json:"company"`
	URL       string            `json:"url"`
	Data      map[string]string `json:"data"`
	IsShared  bool              `json:"is_shared"`
}

type EventResponseContactSearch []Contact

//
// --- Event Responses Sub-Entities

// Agent represents an agent in the smart routing event.
type Agent struct {
	AgentType    string `json:"agent_type"`
	RingDuration int    `json:"ring_duration"`
	RingDelay    int    `json:"ring_delay"`
	Order        int    `json:"order"`
	Number       int64  `json:"number"`
	IsPreAnswer  int    `json:"is_pre_answer"`
	IsCallerID   int    `json:"is_caller_id"`
	IsHeadLine   int    `json:"is_head_line"`
}

// Contact represents an individual contact in a search response.
type Contact struct {
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Company   string   `json:"company"`
	URL       string   `json:"url"`
	Numbers   []Number `json:"numbers"`
}

// Number represents a contact's phone number.
type Number struct {
	Number int64  `json:"number"`
	Type   string `json:"type"`
}

// WebhookEventResponder allows clients to provide custom responses for webhook events.
type WebhookEventResponder interface {
	// SmartRoutingResponse returns a response for the smart routing event.
	SmartRoutingResponse(req WebhookRequestDTO) (*EventResponseSmartRouting, error)
	// ContactCallResponse returns a response for the contact call event.
	ContactCallResponse(req WebhookRequestDTO) (*EventResponseContactCall, error)

	ContactSearchResponse(req WebhookRequestDTO) (*EventResponseContactSearch, error)
}
