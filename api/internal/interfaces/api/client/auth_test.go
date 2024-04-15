package client_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/kodmain/thetiptop/api/config"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/observability/logger"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/serializers/buffer"
	serializer "github.com/kodmain/thetiptop/api/internal/infrastructure/serializers/jwt"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/server"
	"github.com/kodmain/thetiptop/api/internal/interfaces"
	"github.com/stretchr/testify/assert"
)

type TokenStructure struct {
	JWT string `json:"jwt"`
}

const (
	GOOD_EMAIL = "user1@example.com"
	GOOD_PASS  = "ValidP@ssw0rd1"

	WRONG_EMAIL = "user2@example.com"
	WRONG_PASS  = "secret"
)

var srv *server.Server

func start(http, https string) error {
	config.Load("../../../../config.test.yml")
	logger.Info("starting application")
	config.PORT_HTTP = http
	config.PORT_HTTPS = https
	srv = server.Create()
	srv.Register(interfaces.Endpoints)
	return srv.Start()
}

func stop() error {
	logger.Info("waiting for application to shutdown")
	return srv.Stop()
}

func request(method, uri string, token string, values ...map[string][]string) ([]byte, int, error) {
	// Create a form with email and password fields
	form := url.Values{}
	if len(values) > 0 {
		for key, value := range values[0] {
			form.Set(key, value[0])
		}
	}

	// Create a new HTTP request to call /sign/up
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, 0, err
	}

	// Set the form as the request body
	if len(values) > 0 {
		req.Body = io.NopCloser(strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}

	// Perform the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	content, err := buffer.Read(resp.Body)

	return content.Bytes(), resp.StatusCode, nil
}

func TestSignUp(t *testing.T) {
	assert.Nil(t, start(":8081", ":8444"))

	users := []struct {
		email    string
		password string
		status   int
	}{
		{GOOD_EMAIL, GOOD_PASS, http.StatusCreated},
		{GOOD_EMAIL, GOOD_PASS, http.StatusConflict},
		{WRONG_EMAIL, WRONG_PASS, http.StatusBadRequest},
	}

	for _, user := range users {
		values := map[string][]string{
			"email":    {user.email},
			"password": {user.password},
		}

		_, status, err := request("POST", "http://localhost:8081/sign/up", "", values)
		assert.Nil(t, err)
		assert.Equal(t, status, user.status)
	}

	assert.Nil(t, stop())
}

func TestSignIn(t *testing.T) {
	assert.Nil(t, start(":8082", ":8445"))

	request("POST", "http://localhost:8082/sign/up", "", map[string][]string{
		"email":    {GOOD_EMAIL},
		"password": {GOOD_PASS},
	})

	users := []struct {
		email    string
		password string
		status   int
	}{
		{GOOD_EMAIL, GOOD_PASS, http.StatusOK},
		{GOOD_EMAIL, WRONG_PASS, http.StatusBadRequest},
		{GOOD_PASS, WRONG_PASS, http.StatusBadRequest},
	}

	for _, user := range users {
		values := map[string][]string{
			"email":    {user.email},
			"password": {user.password},
		}

		_, status, err := request("POST", "http://localhost:8082/sign/in", "", values)
		assert.Nil(t, err)
		assert.Equal(t, status, user.status)
	}

	assert.Nil(t, stop())
}

func TestSignRenew(t *testing.T) {
	assert.Nil(t, start(":8083", ":8446"))

	request("POST", "http://localhost:8083/sign/up", "", map[string][]string{
		"email":    {GOOD_EMAIL},
		"password": {GOOD_PASS},
	})

	content, _, _ := request("POST", "http://localhost:8083/sign/in", "", map[string][]string{
		"email":    {GOOD_EMAIL},
		"password": {GOOD_PASS},
	})

	// Déclaration de la variable qui recevra la valeur désérialisée
	var tokenData TokenStructure

	// Désérialisation du JSON dans la structure définie
	err := json.Unmarshal(content, &tokenData)
	assert.Nil(t, err)

	access, err := serializer.TokenToClaims(tokenData.JWT)
	assert.Nil(t, err)

	users := []struct {
		token  string
		status int
	}{
		{"Bearer " + *access.Refresh, http.StatusOK}, // Replace with actual valid JWT token
		{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMxMDkxMzEsImlkIjoiN2M3OTQwMGYtMDA2YS00NzVlLTk3YjYtNWRiZGUzNzA3NjAxIiwib2ZmIjo3MjAwLCJ0eXBlIjoxLCJ0eiI6IkxvY2FsIn0.5Lae56HNcQ1OHcP_FhTfcOOtHpaZVgRFy6vzzBugN7Y", http.StatusUnauthorized}, // Replace with actual expired JWT token
		{"Bearer malformed.jwt.token.here", http.StatusUnauthorized}, // Replace with actual malformed JWT token
		{"", http.StatusBadRequest},                                  // Replace with actual empty JWT token
	}

	for _, user := range users {
		_, status, err := request("GET", "http://localhost:8083/sign/renew", user.token)
		assert.Nil(t, err)
		assert.Equal(t, status, user.status)
	}

	assert.Nil(t, stop())
}