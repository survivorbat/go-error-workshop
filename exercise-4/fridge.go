package fridge

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type FridgeClient struct {
	Host string
}

func (f *FridgeClient) Configure(cfg *Config) error {
	reqBody, _ := json.Marshal(cfg)

	// ❗ The code below is yours to finish.

	req, err := http.NewRequest(http.MethodPut, f.Host+"/api/config", bytes.NewReader(reqBody))
	if err != nil {
		// ❗ Test Requirements:
		// - Must errors.AsType to *url.Error returned from NewRequest
		// - Must contain "failed to build request for URL {f.Host}"
		return fmt.Errorf("failed to build request for URL %s: %w", f.Host, err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// ❗ Test requirements:
		// - Must errors.Is to ErrConnectionError
		// - Must errors.AsType to *net.Error (returned from Do)
		return errors.Join(err, ErrConnectionError)
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
		// ❗ Test requirements:
		// - Must errors.Is to ErrAPIError
		// - Must errors.AsType to *ResponseError
		// - Must errors.AsType to *json.SyntaxError (returned from Unmarshal)
		return &ResponseError{
			Body:      resBody,
			ActualErr: err,
		}
	}

	// ❗ Test requirements:
	// - Must errors.AsType to *ConfigError
	return cfgErr
}
