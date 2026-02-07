package types

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"time"
)

type CreateSecretRequest struct {
	FilePath      string         `json:"file_path"`
	SecretWrapper Secret         `json:"secret_wrapper"`
	StoreName     string         `json:"store_name"`
	Name          string         `json:"name"`
	RerollTime    RerollDuration `json:"reroll_time"`
}

type Secret struct {
	Content          string            `json:"content,omitempty"`
	SecretReferences map[string]string `json:"secret_references,omitempty"`
}

var durationRegex = regexp.MustCompile(`(?i)(?:(\d+)D)?(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)

// Own type for Easy TypeSaftey without hacky iso standards
type RerollDuration time.Duration

// UnmarshalJSON implements the json.Unmarshaler interface
func (r *RerollDuration) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	// Check if the data starts with a quote (it's a string from the API)
	//TODO unsecure
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		duration, err := parseString(s) // Your regex parser
		if err != nil {
			return err
		}
		*r = RerollDuration(duration)
		return nil
	}

	// Otherwise, treat it as a number (from the Database)
	var nanoseconds int64
	if err := json.Unmarshal(b, &nanoseconds); err != nil {
		return err
	}
	*r = RerollDuration(nanoseconds)
	return nil
}

// Helper to keep logic clean
func parseString(s string) (time.Duration, error) {
	matches := durationRegex.FindStringSubmatch(s)
	if len(matches) == 0 {
		return 0, errors.New("invalid duration format")
	}

	var total time.Duration
	var unitFound bool // Track if at least one capture group was filled
	units := []time.Duration{24 * time.Hour, time.Hour, time.Minute, time.Second}

	for i, unit := range units {
		if matches[i+1] != "" {
			val, _ := strconv.Atoi(matches[i+1])
			total += time.Duration(val) * unit
			unitFound = true
		}
	}

	// If the regex "matched" but all capture groups were empty (e.g., an empty or weird string)
	if !unitFound {
		return 0, errors.New("invalid duration format: no valid units provided")
	}

	return total, nil
}
