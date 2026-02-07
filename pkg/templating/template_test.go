package templating

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvaluateTemplate(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		data     map[string]string
		expected string
		wantErr  bool
	}{
		{
			name:    "Successful substitution",
			content: "DB_URL=postgres://{{.User}}:{{.Pass}}@localhost:5432",
			data: map[string]string{
				"User": "admin",
				"Pass": "secret123",
			},
			expected: "DB_URL=postgres://admin:secret123@localhost:5432",
			wantErr:  false,
		},
		{
			name:    "Strict missing key failure",
			content: "API_KEY={{.ApiKey}}",
			data:    map[string]string{"WrongKey": "value"},
			// This should now fail because of "missingkey=error"
			wantErr: true,
		},
		{
			name:    "Invalid template syntax",
			content: "Malformed {{.Token",
			data:    map[string]string{"Token": "abc"},
			wantErr: true,
		},
		{
			name:     "Empty data with static content",
			content:  "STATIC_VAL=123",
			data:     map[string]string{},
			expected: "STATIC_VAL=123",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf, err := EvaluateTemplate(tt.content, tt.data, tt.name)

			if tt.wantErr {
				assert.Error(t, err, "Expected an error for case: %s", tt.name)
				assert.Nil(t, buf)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, buf)
				assert.Equal(t, tt.expected, buf.String())
			}
		})
	}
}
