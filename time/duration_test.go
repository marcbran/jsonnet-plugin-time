package time

import (
	"testing"
	stdtime "time"

	"github.com/stretchr/testify/require"
)

func TestAddDurationSpec(t *testing.T) {
	ref := stdtime.Date(2026, stdtime.January, 31, 12, 0, 0, 0, stdtime.UTC)

	tests := []struct {
		name      string
		spec      string
		want      stdtime.Time
		wantError bool
	}{
		{
			name: "hours only",
			spec: "6h",
			want: ref.Add(6 * stdtime.Hour),
		},
		{
			name: "compound hours and minutes",
			spec: "2h30m",
			want: ref.Add(2*stdtime.Hour + 30*stdtime.Minute),
		},
		{
			name: "days",
			spec: "7D",
			want: ref.AddDate(0, 0, 7),
		},
		{
			name: "weeks",
			spec: "2W",
			want: ref.AddDate(0, 0, 14),
		},
		{
			name: "months land on the right day even with variable month length",
			spec: "1M",
			want: stdtime.Date(2026, stdtime.March, 3, 12, 0, 0, 0, stdtime.UTC),
		},
		{
			name: "years account for leap days",
			spec: "1Y",
			want: ref.AddDate(1, 0, 0),
		},
		{
			name: "every unit at once",
			spec: "1Y2M3W4D5h6m7s",
			want: ref.AddDate(1, 2, 3*7+4).Add(5*stdtime.Hour + 6*stdtime.Minute + 7*stdtime.Second),
		},
		{
			name: "negative sign subtracts",
			spec: "-6h",
			want: ref.Add(-6 * stdtime.Hour),
		},
		{
			name: "positive sign is explicit no-op on direction",
			spec: "+90s",
			want: ref.Add(90 * stdtime.Second),
		},
		{
			name:      "invalid unit",
			spec:      "6x",
			wantError: true,
		},
		{
			name:      "empty spec",
			spec:      "",
			wantError: true,
		},
		{
			name:      "sign with no digits",
			spec:      "-",
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddDurationSpec(ref, tt.spec)
			if tt.wantError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.True(t, tt.want.Equal(got), "want %s, got %s", tt.want, got)
		})
	}
}
