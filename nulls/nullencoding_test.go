package nulls

import "testing"

func TestNullEncoding(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NullEncoding(); (err != nil) != tt.wantErr {
				t.Errorf("NullEncoding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}