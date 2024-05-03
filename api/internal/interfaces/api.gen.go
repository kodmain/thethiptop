// Automatically generated by api/generator/api.gen.go, DO NOT EDIT manually
// Package api implements Register method for fiber
package interfaces

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kodmain/thetiptop/api/internal/docs"
	"github.com/kodmain/thetiptop/api/internal/interfaces/api/client"
	"github.com/kodmain/thetiptop/api/internal/interfaces/status"
	"github.com/swaggo/swag"
)

func init() {
	json.Unmarshal([]byte(doc), Mapping)
}

// API represents a collection of HTTP endpoints grouped by namespace and version.
var (
	Endpoints map[string]fiber.Handler = map[string]func(*fiber.Ctx) error{
		"status.IP":             status.IP,
		"client.UpdateComplete": client.UpdateComplete,
		"client.Delete":         client.Delete,
		"client.SignUp":         client.SignUp,
		"client.SignRenew":      client.SignRenew,
		"client.Reset":          client.Reset,
		"client.SignIn":         client.SignIn,
		"status.HealthCheck":    status.HealthCheck,
		"client.FindOne":        client.FindOne,
		"client.Find":           client.Find,
		"client.UpdatePartial":  client.UpdatePartial,
		"client.Renew":          client.Renew,
	}
	Mapping = &docs.Swagger{}
	doc, _  = swag.ReadDoc()
)
