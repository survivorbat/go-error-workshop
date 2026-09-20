package fridge

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFridgeClient_Configure_Success(t *testing.T) {
	t.Parallel()
	// Arrange
	var actualBody []byte

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actualBody, _ = io.ReadAll(r.Body)
		// Server returns 200
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	client := &FridgeClient{Host: server.URL}

	cfg := &Config{Name: "Fridge-y", Temperature: 5}

	// Act
	err := client.Configure(cfg)

	// Assert
	require.NoError(t, err)

	expectedBody := `{"name":"Fridge-y","temperature":5}`
	assert.JSONEq(t, expectedBody, string(actualBody))
}

func TestFridgeClient_Configure_ReturnsErrorOnInvalidBaseURL(t *testing.T) {
	t.Parallel()
	// Arrange
	client := &FridgeClient{Host: ":::/"}

	cfg := &Config{Name: "Fridge-y", Temperature: 5}

	// Act
	err := client.Configure(cfg)

	// Assert
	require.ErrorContains(t, err, "failed to build request")
	require.ErrorAs(t, err, new(url.Error))
}

func TestFridgeClient_Configure_ReturnsErrorOnConnectionIssue(t *testing.T) {
	t.Parallel()
	// Arrange
	client := &FridgeClient{Host: "http://localhost:1"}

	cfg := &Config{Name: "Fridge-y", Temperature: 5}

	// Act
	err := client.Configure(cfg)

	// Assert
	require.ErrorIs(t, err, ErrConnectionError)
	require.ErrorAs(t, err, new(net.Error))
}

func TestFridgeClient_Configure_ReturnsConfigErrorsFromAPI(t *testing.T) {
	t.Parallel()
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"message":"Config was invalid"}`))
	}))
	t.Cleanup(server.Close)

	client := &FridgeClient{Host: server.URL}

	cfg := &Config{Name: "🥶", Temperature: -20}

	// Act
	err := client.Configure(cfg)

	// Assert
	require.ErrorIs(t, err, ErrConnectionError)

	var actual *ConfigError
	require.ErrorAs(t, err, &actual)
	assert.Equal(t, "Config was invalid", actual.Message)
}

func TestFridgeClient_Configure_ReturnsResponseError(t *testing.T) {
	t.Parallel()
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("}")) // Invalid JSON
	}))
	t.Cleanup(server.Close)

	client := &FridgeClient{Host: server.URL}

	cfg := &Config{Name: "🥶", Temperature: -20}

	// Act
	err := client.Configure(cfg)

	// Assert
	require.ErrorIs(t, err, ErrAPIError)

	var actual *ResponseError
	require.ErrorAs(t, err, &actual)
	assert.Equal(t, "}", string(actual.Body))

	require.ErrorAs(t, err, new(json.InvalidUnmarshalError))
}
