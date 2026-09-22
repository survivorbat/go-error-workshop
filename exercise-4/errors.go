package fridge

import "errors"

var (
	// ErrConnectionError is used if we failed to reach the API
	ErrConnectionError = errors.New("connection error to the fridge")

	// ErrAPIError is used if the API returns something we didn't expect
	ErrAPIError = errors.New("api error")
)

// ConfigError is returned from the API if the provided config does not make sense
// for the fridge to be configured with.
//
// ❗ Test requirements:
// - Must be an error
// - Must errors.As to ConfigError
type ConfigError struct {
	Message string `json:"message"`
}

func (c *ConfigError) Error() string {
	return c.Message
}

/////////////////////////////////////
// ❗ Something is missing here... //
/////////////////////////////////////

// ResponseError is returned if we failed to parse the response from the API
// and want to wrap it with additional information.
//
// ❗ Test requirements:
// - Must be an error
// - Must errors.Is to ErrAPIError
// - Must errors.As to ActualErr
type ResponseError struct {
	Body []byte

	ActualErr error
}

// Error returns a static message
func (c *ResponseError) Error() string {
	return "Failed to read API response body"
}

////////////////////////////////////////
// ❗ Something(s) is missing here... //
////////////////////////////////////////

func (r *ResponseError) Unwrap() error {
	return r.ActualErr
}

func (r *ResponseError) Is(err error) bool {
	return errors.Is(err, ErrAPIError)
}
