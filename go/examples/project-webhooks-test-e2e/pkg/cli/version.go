package cli

import (
	"fmt"
	"strings"
)

var (
	// Naming
	SERVICE_NAME = "Ringover Webhooks Collector API"
	SERVICE_SLUG = strings.ReplaceAll(strings.ToLower(SERVICE_NAME), " ", "-")

	// Versioning
	SERVICE_VERSION_APP = "0.1.0"
	SERVICE_VERSION_API = "0.1.0"
)

func ServiceWithVersionApp() string {
	return fmt.Sprintf("%s:%s", SERVICE_SLUG, SERVICE_VERSION_APP)
}

func VersionsAppApi() string {
	return fmt.Sprintf("%s:%s", SERVICE_VERSION_APP, SERVICE_VERSION_API)
}
