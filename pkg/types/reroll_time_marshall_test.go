package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRerollDuration_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		wantErr  bool
	}{
		// --- String parsing (API inputs) ---
		{
			name:     "Parse Days and Hours",
			input:    `"1D12H"`,
			expected: 24*time.Hour + 12*time.Hour,
			wantErr:  false,
		},
		{
			name:     "Parse Minutes and Seconds",
			input:    `"30M45S"`,
			expected: 30*time.Minute + 45*time.Second,
			wantErr:  false,
		},
		{
			name:     "Case Insensitivity",
			input:    `"1d5m"`,
			expected: 24*time.Hour + 5*time.Minute,
			wantErr:  false,
		},

		// --- Number parsing (Database/Int inputs) ---
		{
			name:     "Parse Raw Nanoseconds",
			input:    `3600000000000`, // 1 Hour in ns
			expected: time.Hour,
			wantErr:  false,
		},

		// --- Edge Cases ---
		{
			name:    "Invalid Format String",
			input:   `"NotADuration"`,
			wantErr: true, // Should fail regex/parsing
		},
		{
			name:    "Invalid JSON Type",
			input:   `{"key": "value"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rd RerollDuration
			err := json.Unmarshal([]byte(tt.input), &rd)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, time.Duration(rd), tt.expected)
			}
		})
	}
}

func TestCreateSecretRequest_FullUnmarshal(t *testing.T) {
	rawJSON := `{
		"name": "my-app-secret",
		"store_name": "vault-prod",
		"reroll_time": "2D",
		"secret_wrapper": {
			"content": "raw-data",
			"secret_references": {
				"DB_PASS": "kv/data/db/password"
			}
		}
	}`

	var req CreateSecretRequest
	err := json.Unmarshal([]byte(rawJSON), &req)

	assert.NoError(t, err)
	assert.Equal(t, "my-app-secret", req.Name)
	assert.Equal(t, time.Duration(48*time.Hour), time.Duration(req.RerollTime))
	assert.Equal(t, "raw-data", req.SecretWrapper.Content)
	assert.Equal(t, "kv/data/db/password", req.SecretWrapper.SecretReferences["DB_PASS"])
}
