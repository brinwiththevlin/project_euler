package p003

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  uint64
	}{
		{
			name:  "largest prime factor of 13195",
			limit: 13195,
			want:  29,
		},
		{
			name:  "largest prime factor of ",
			limit: 600851475143,
			want:  6857,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solve(tt.limit)
			if got != tt.want {
				t.Errorf("Solve(%d) = %d; want %d", tt.limit, got, tt.want)
			}
		})
	}
}

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		_ = Solve(600851475143)
	}
}
