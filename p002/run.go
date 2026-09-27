package p002

import (
	"fmt"
	"strconv"

	"github.com/brinwiththevlin/project_euler/registry"
)

func init() {
	registry.Register(2, func(input string) (string, error) {
		target, err := strconv.Atoi(input)
		if err != nil {
			return "", fmt.Errorf("invalid limit: %w", err)
		}
		if target <= 2 {
			return "", fmt.Errorf("limit must be greater than 2, got %d", target)
		}
		return strconv.FormatUint(Solve(target), 10), nil
	})
}
