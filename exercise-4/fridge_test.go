package fridge

import (
	"io"
	"net/http"
	"net/http/httptest"
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

// TODO: Finish this
// func TestFridgeClient_Configure_ReturnsExpectedErrors(t *testing.T) {
// 	t.Parallel()
//
// 	tests := map[string]struct{}{}
//
// 	for name, testData := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()
// 			// Arrange
// 			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.WriteHeader(http.StatusOK)
// 			}))
// 			t.Cleanup(server.Close)
//
// 			client := &FridgeClient{Host: server.URL}
//
// 			cfg := &Config{Name: "Fridge-y", Temperature: 5}
//
// 			// Act
// 			err := client.Configure(cfg)
//
// 			// Assert
// 			require.NoError(t, err)
// 		})
// 	}
// }
