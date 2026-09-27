package p062

import (
	"fmt"
	"strconv"

	"github.com/brinwiththevlin/project_euler/registry"
)

func init() {
	registry.Register(62, func(input string) (string, error) {
		target, err := strconv.Atoi(input)
		if err != nil {
			return "", fmt.Errorf("invalid permutation count: %w", err)
		}
		return strconv.FormatUint(Solve(target), 10), nil
	})
}
