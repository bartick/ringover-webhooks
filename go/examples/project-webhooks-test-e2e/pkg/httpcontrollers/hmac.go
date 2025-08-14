package httpcontrollers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	callEventKey     []byte
	contactCallKey   []byte
	contactSearchKey []byte
	smartRoutingKey  []byte
)

func init() {
	callEventKey = []byte(os.Getenv("WEBHOOKS_CALL_EVENT_KEY"))
	contactCallKey = []byte(os.Getenv("WEBHOOKS_CONTACT_CALL_KEY"))
	contactSearchKey = []byte(os.Getenv("WEBHOOK_CONTACT_SEARCH_KEY"))
	smartRoutingKey = []byte(os.Getenv("WEBHOOKS_SMART_ROUTING_AND_IVR_CODE_KEY"))
}

// getSigningKey selects the signing key based on resource/event
func getSigningKey(resource, event string) []byte {
	switch {
	case event == "contact" && resource == "call":
		return contactCallKey
	case event == "contact" && resource == "search":
		return contactSearchKey
	case event == "routing" && resource == "call":
		return smartRoutingKey
	case resource == "call":
		return callEventKey
	default:
		return callEventKey // fallback
	}
}

// WebhookAuthMiddleware parses resource/event from the request body and uses the correct key
func WebhookAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log headers properly
		for key, values := range r.Header {
			fmt.Printf("[DEBUG] Header: %s = %v\n", key, values)
		}

		if reqHeadersBytes, err := json.Marshal(r.Header); err != nil {
			fmt.Printf("Could not Marshal Req Headers")
		} else {
			fmt.Printf(string(reqHeadersBytes))
		}

		// Fetch the signature header in a case-insensitive manner
		signature := r.Header.Get("X-Ringover-Webhook-Signature")
		if signature == "" {
			fmt.Println("[DEBUG] Missing X-Ringover-Webhook-Signature header")
			http.Error(w, "Missing X-Ringover-Webhook-Signature header", http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Parse resource and event from body (assume JSON)
		// We have to unmarshal the body here because the jwt secret key vary depending on the resource & events fields value

		resource, event := "", ""
		var payload struct {
			Resource string `json:"resource"`
			Event    string `json:"event"`
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			fmt.Printf("[DEBUG] Failed to parse request body: %v\n", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		resource = payload.Resource
		event = payload.Event

		secret := getSigningKey(resource, event)

		// Parse the JWT from the signature header
		token, err := jwt.Parse(signature, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != "HS512" {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		}, jwt.WithValidMethods([]string{"HS512"}))
		if err != nil || !token.Valid {
			fmt.Printf("[DEBUG] JWT signature verification failed: %v\n", err)
			fmt.Printf("[DEBUG] Secret used: %s\n", secret)
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
			return
		}

		// Restore the body for further processing
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		next.ServeHTTP(w, r)
	})
}
