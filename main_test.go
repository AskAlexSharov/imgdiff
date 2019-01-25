package main

import "testing"

func Test_distance(t *testing.T) {
	tests := []struct {
		name         string
		fname1       string
		fname2       string
		wantDistance int
	}{
		{
			name:         "Original vs Scaled Down - same",
			fname1:       "./test-png-original.png",
			fname2:       "./test-png-scaled-down.png",
			wantDistance: 0,
		},
		{
			name:         "Original vs Damaged - 3%",
			fname1:       "./test-png-original.png",
			fname2:       "./test-png-damaged.png",
			wantDistance: 3,
		},
		{
			name:         "Scaled Down vs Damaged - 3%",
			fname1:       "./test-png-scaled-down.png",
			fname2:       "./test-png-damaged.png",
			wantDistance: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDistance := distance(tt.fname1, tt.fname2)
			if gotDistance != tt.wantDistance {
				t.Errorf("distance() gotDistance = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
