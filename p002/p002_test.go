package p002

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		limit int
		want  uint64
	}{
		{
			name:  "no even terms below the first (limit 1)",
			limit: 1,
			want:  0,
		},
		{
			name:  "limit equal to an even term includes it (limit 8)",
			limit: 8,
			want:  10,
		},
		{
			name:  "Project Euler problem 2 (under 90)",
			limit: 90,
			want:  44,
		},
		{
			name:  "Project Euler Problem 2 (under 4,000,000)",
			limit: 4000000,
			want:  4613732, // Update this with your verified answer once discovered
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

// BenchmarkSolve measures how efficient your sorting and map lookup logic is.
// Ideally, this algorithm should finish well within a few milliseconds.
func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		_ = Solve(1000)
	}
}
