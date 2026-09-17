package time

import (
	"testing"
	stdtime "time"

	"github.com/stretchr/testify/require"
)

func TestFormatTime(t *testing.T) {
	orig := stdtime.Local
	stdtime.Local = stdtime.UTC
	defer func() { stdtime.Local = orig }()

	ms := float64(stdtime.Date(2026, stdtime.August, 4, 12, 30, 45, 0, stdtime.UTC).UnixMilli())

	tests := []struct {
		name   string
		layout string
		want   string
	}{
		{
			name:   "date and time",
			layout: "2006-01-02 15:04",
			want:   "2026-08-04 12:30",
		},
		{
			name:   "time only",
			layout: "15:04:05",
			want:   "12:30:45",
		},
		{
			name:   "rfc3339",
			layout: stdtime.RFC3339,
			want:   "2026-08-04T12:30:45Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, FormatTime(ms, tt.layout))
		})
	}
}
