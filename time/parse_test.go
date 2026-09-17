package time

import (
	"testing"
	stdtime "time"

	"github.com/stretchr/testify/require"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		layout    string
		want      stdtime.Time
		wantError bool
	}{
		{
			name:   "rfc3339 utc with Z",
			value:  "2026-08-04T00:00:00Z",
			layout: stdtime.RFC3339,
			want:   stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:   "rfc3339 with milliseconds",
			value:  "2026-08-04T00:00:00.000Z",
			layout: stdtime.RFC3339,
			want:   stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:   "rfc3339 with explicit offset",
			value:  "2026-08-04T02:00:00+02:00",
			layout: stdtime.RFC3339,
			want:   stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:   "custom date-only layout",
			value:  "2026-08-04",
			layout: "2006-01-02",
			want:   stdtime.Date(2026, stdtime.August, 4, 0, 0, 0, 0, stdtime.UTC),
		},
		{
			name:      "value does not match layout",
			value:     "2026-08-04",
			layout:    stdtime.RFC3339,
			wantError: true,
		},
		{
			name:      "empty value",
			value:     "",
			layout:    stdtime.RFC3339,
			wantError: true,
		},
		{
			name:      "garbage",
			value:     "not-a-time",
			layout:    stdtime.RFC3339,
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTime(tt.value, tt.layout)
			if tt.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.True(t, tt.want.Equal(got), "want %s, got %s", tt.want, got)
		})
	}
}
