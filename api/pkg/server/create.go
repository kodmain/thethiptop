// Package server provides a simple HTTP/HTTPS server implementation for serving static and connect API.
package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/kodmain/thetiptop/api/config"
	"github.com/kodmain/thetiptop/api/internal/docs"
	"github.com/kodmain/thetiptop/api/internal/model/database"
)

var servers map[string]*Server = make(map[string]*Server)

func getConfig(cfgs ...fiber.Config) fiber.Config {
	if len(cfgs) > 0 {
		return cfgs[0]
	}

	cfg := fiber.Config{
		AppName:               config.APP_NAME,
		Prefork:               true, // Multithreading
		DisableStartupMessage: true, // Disable startup message
	}

	if os.Getppid() <= 1 {
		fmt.Println("WARNING: fiber in downgrade mode please use docker run --pid=host")
		cfg.Prefork = false // Disable to prevent bug in container
	}

	return cfg
}

// Create return a instance os Server
// Pattern Singleton
func Create(cfgs ...fiber.Config) *Server {
	cfg := getConfig(cfgs...)
	if server, exists := servers[cfg.AppName]; exists {
		return server
	}
	db, err := database.New(config.DB_DSN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server := &Server{
		app: fiber.New(cfg),
		db:  db,
	}

	server.app.Use(setGoToDoc)         // register middleware setGoToDoc
	server.app.Use(setSecurityHeaders) // register middleware setSecurityHeaders
	server.app.Get("/docs/*", swagger.New(swagger.Config{
		Title:                    config.APP_NAME,
		Layout:                   "BaseLayout",
		DocExpansion:             "list",
		DefaultModelsExpandDepth: 2,
	})) // register middleware for documentation

	servers[cfg.AppName] = server

	return server
}

// setGoToDoc is a middleware that redirect to /docs url path is like /
func setGoToDoc(c *fiber.Ctx) error {
	if c.Path() == "/index.html" || c.Path() == "/" {
		return c.Redirect("/docs", 301)
	}
	return c.Next()
}

// setSecurityHeaders is a middleware that grants best practice around security
func setSecurityHeaders(c *fiber.Ctx) error {
	// Activer HSTS (HTTP Strict Transport Security)
	c.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
	// Activer CSP (Content Security Policy)
	c.Set("Content-Security-Policy", "default-src 'unsafe-inline' 'self' fonts.gstatic.com fonts.googleapis.com;img-src data: 'self'")
	// Activer CORS (Cross-Origin Resource Sharing)
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET,POST,HEAD,PUT,DELETE,PATCH")
	c.Set("Access-Control-Allow-Headers", "*")
	c.Set("Access-Control-Allow-Credentials", "true")

	docs.SwaggerInfo.Host = c.Hostname()

	return c.Next()
}
