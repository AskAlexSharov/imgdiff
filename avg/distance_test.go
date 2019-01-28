package avg

import (
	"context"
	"testing"
)

func Test_distance(t *testing.T) {
	tests := []struct {
		name         string
		file1        string
		file2        string
		wantDistance int
	}{
		{
			name:         "Original vs Scaled Down - same",
			file1:        "./test-png-original.png",
			file2:        "./test-png-scaled-down.png",
			wantDistance: 0,
		},
		{
			name:         "Original vs Damaged - 1%",
			file1:        "./test-png-original.png",
			file2:        "https://raw.githubusercontent.com/AskAlexSharov/imgdiff/master/test-png-damaged.png",
			wantDistance: 1,
		},
		{
			name:         "Scaled Down vs Damaged - 1%",
			file1:        "https://raw.githubusercontent.com/AskAlexSharov/imgdiff/master/test-png-scaled-down.png",
			file2:        "./test-png-damaged.png",
			wantDistance: 1,
		},
	}
	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			gotDistance, err := Distance(context.Background(), tc.file1, tc.file2)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if gotDistance != tc.wantDistance {
				t.Errorf("distance() gotDistance = %v, want %v", gotDistance, tc.wantDistance)
			}
		})
	}
}
