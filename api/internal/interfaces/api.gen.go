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
		"client.SignUp":         client.SignUp,
		"client.SignIn":         client.SignIn,
		"client.FindOne":        client.FindOne,
		"client.Find":           client.Find,
		"client.Delete":         client.Delete,
		"status.IP":             status.IP,
		"client.SignOut":        client.SignOut,
		"client.UpdatePartial":  client.UpdatePartial,
		"client.UpdateComplete": client.UpdateComplete,
		"client.Renew":          client.Renew,
		"client.Reset":          client.Reset,
		"status.HealthCheck":    status.HealthCheck,
	}
	Mapping = &docs.Swagger{}
	doc, _  = swag.ReadDoc()
)
