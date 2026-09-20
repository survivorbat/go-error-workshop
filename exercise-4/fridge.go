package fridge

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	// ErrConnectionError is used if we failed to reach the API
	ErrConnectionError = errors.New("connection error to the fridge")

	// ErrAPIError is used if we failed to talk to the API
	ErrAPIError = errors.New("api error")
)

type FridgeClient struct {
	Host string
}

// ❗ The code below is yours to implement.

func (f *FridgeClient) Configure(cfg *Config) error {
	reqBody, _ := json.Marshal(cfg)

	req, err := http.NewRequest(http.MethodPut, f.Host+"/api/config", bytes.NewReader(reqBody))
	if err != nil {
		// Test Requirements:
		// - Must errors.AsType to url.Error (from stdlib)
		// - Must contain "failed to build request"
		return nil
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// Test requirements:
		// - Must errors.Is to ErrConnectionError
		// - Must errors.AsType to net.Error (from stdlib)
		return nil
	}

	defer func() { _ = res.Body.Close() }()

	// Success! Nothing to do here
	if res.StatusCode == http.StatusOK {
		return nil
	}

	resBody, _ := io.ReadAll(res.Body) // Suppressing this for simplicity

	var cfgErr *ConfigError
	err = json.Unmarshal(resBody, &cfgErr)
	if err != nil {
		// Test requirements:
		// - Must errors.Is to ErrAPIError
		// - Must errors.AsType to ResponseError
		// - Must errors.AsType to json.InvalidUnmarshalError
		return nil
	}

	// Test requirements:
	// - Must errors.AsType to ConfigError
	return nil
}

// ConfigError is returned from the API if the provided config does not make sense
// for the fridge to be configured as.
type ConfigError struct {
	Message string `json:"message"`
}

// Error returns the message from the API
func (c *ConfigError) Error() string {
	return c.Message
}

// ResponseError is returned if we failed to parse the response from the API
// and want to wrap it with additional information.
type ResponseError struct {
	Body []byte

	Err error
}

// Error returns a static message
func (c *ResponseError) Error() string {
	return "Failed to read API response body"
}
