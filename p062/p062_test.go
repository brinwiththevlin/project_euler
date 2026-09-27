package p062

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		name        string
		targetCount int
		want        uint64
	}{
		{
			name:        "Project Euler Example (3 Permutations)",
			targetCount: 3,
			want:        41063625, // 345^3
		},
		// {
		// 	name:        "Project Euler Problem 62 (5 Permutations)",
		// 	targetCount: 5,
		// 	want:        REPLACEME, // Update this with your verified answer once discovered
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solve(tt.targetCount)
			if got != tt.want {
				t.Errorf("Solve(%d) = %d; want %d", tt.targetCount, got, tt.want)
			}
		})
	}
}

// BenchmarkSolve measures how efficient your sorting and map lookup logic is.
// Ideally, this algorithm should finish well within a few milliseconds.
func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
		_ = Solve(5)
	}
}
