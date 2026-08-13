package time

import (
	"testing"
	stdtime "time"

	"github.com/stretchr/testify/require"
)

func TestParseRFC3339Time(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		want      stdtime.Time
		wantError bool
	}{
		{
			name:  "utc with Z",
			value: "2026-08-04T00:00:00Z",
			want:  stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:  "with milliseconds",
			value: "2026-08-04T00:00:00.000Z",
			want:  stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:  "with explicit offset",
			value: "2026-08-04T02:00:00+02:00",
			want:  stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:      "date only is not RFC3339",
			value:     "2026-08-04",
			wantError: true,
		},
		{
			name:      "empty string",
			value:     "",
			wantError: true,
		},
		{
			name:      "garbage",
			value:     "not-a-time",
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRFC3339Time(tt.value)
			if tt.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.True(t, tt.want.Equal(got), "want %s, got %s", tt.want, got)
		})
	}
}
