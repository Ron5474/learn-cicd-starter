package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input     http.Header
		wantKey   string
		wantError error
	}{
		"Valid ApiKey": {
			input:     http.Header{"Authorization": []string{"ApiKey my-secret-api-key"}},
			wantKey:   "my-secret-api-key",
			wantError: nil,
		},
		"Invalid Key": {
			input:     http.Header{"Authorization": []string{"ApiKey"}},
			wantKey:   "",
			wantError: errors.New("malformed authorization header"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotKey, gotError := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.wantKey, gotKey) {
				t.Fatalf("expected key: %v, got: %v", tc.wantKey, gotKey)
			}

			// Corrected error comparison logic
			if tc.wantError == nil && gotError != nil {
				t.Fatalf("expected no error, but got: %v", gotError)
			}
			if tc.wantError != nil && gotError == nil {
				t.Fatalf("expected error: %v, but got none", tc.wantError)
			}
			if tc.wantError != nil && gotError != nil && tc.wantError.Error() != gotError.Error() {
				t.Fatalf("expected error: %v, got: %v", tc.wantError, gotError)
			}
		})
	}
}
