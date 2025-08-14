package cli

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"

	"github.com/ringover/ringover-webhooks/examples/project-webhooks-test-e2e/pkg/httpcontrollers"
)

// Options for the CLI.
type Options struct {
	ListenPort int `help:"Port to listen on" short:"p" default:"16080"`
}

func Run() {

	// Create a CLI app which takes a port option.
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {

		// Logger
		// https://github.com/go-chi/httplog
		logger := httplog.NewLogger(ServiceWithVersionApp(), httplog.Options{
			// JSON:             true,
			LogLevel:         slog.LevelDebug,
			Concise:          true,
			RequestHeaders:   true,
			MessageFieldName: "msg",
			TimeFieldFormat:  time.RFC3339,
			Tags: map[string]string{
				"version": SERVICE_VERSION_APP,
				"env":     "dev",
			},
			QuietDownRoutes: []string{
				"/",
				"/ping",
			},
			QuietDownPeriod: 10 * time.Minute,
			// SourceFieldName: "source",
		})

		router := chi.NewRouter()

		router.Use(httplog.RequestLogger(logger))
		router.Use(middleware.Heartbeat("/ping"))
		router.Use(httpcontrollers.WebhookAuthMiddleware)

		router.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				httplog.LogEntrySetField(ctx, "user", slog.StringValue("user1"))

				next.ServeHTTP(w, r.WithContext(ctx))
			})
		})

		api := humachi.New(router, huma.DefaultConfig(SERVICE_NAME, SERVICE_VERSION_APP))

		// Register GET /greeting/{name}

		// To receive webhooks from Ringover API v2
		//huma.NewGroup(api, "/v2")

		// Register POST /ringover/webhooks
		huma.Register(api, huma.Operation{
			OperationID: "ringover-webhooks-push",
			Method:      http.MethodPost,
			Path:        "/ringover/webhooks",
			Summary:     "Push a Ringover Webhook",
			Description: "Push a Ringover Webhook and validate it based on its resource and event types",
			Tags:        []string{"Webhooks"},
		}, httpcontrollers.PushWebhook)

		huma.Register(api, huma.Operation{
			OperationID: "ringover-webhooks-push-routing",
			Method:      http.MethodPost,
			Path:        "/ringover/webhooks/routing",
			Summary:     "Push a Ringover Smart Routing Webhook",
			Description: "Push a Ringover Webhook for smart routing event and get a smart routing response.",
		}, httpcontrollers.PushWebhookSmartRouting)

		huma.Register(api, huma.Operation{
			OperationID: "ringover-webhooks-push-contact",
			Method:      http.MethodPost,
			Path:        "/ringover/webhooks/contact",
			Summary:     "Push a Ringover Contact Call Webhook",
			Description: "Push a Ringover Webhook for contact call event and get a contact call response.",
			// Tags:        []string{"Webhooks"},
		}, httpcontrollers.PushWebhookContactCall)

		// If you want to leave /search open, do not wrap it with the middleware
		huma.Register(api, huma.Operation{
			OperationID: "ringover-webhooks-push-search",
			Method:      http.MethodPost,
			Path:        "/ringover/webhooks/search",
			Summary:     "Push a Ringover Contact Search Webhook",
			Description: "Push a Ringover Webhook for contact search event and get a contact search response.",
			// Tags:        []string{"Webhooks"},
		}, httpcontrollers.PushWebhookContactSearch)

		// Tell the CLI how to start your server.
		hooks.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", options.ListenPort)
			http.ListenAndServe(fmt.Sprintf(":%d", options.ListenPort), router)
		})
	})

	// Run the CLI. When passed no commands, it starts the server.
	cli.Run()
}
