package p005

import (
	"fmt"
	"strconv"

	"github.com/brinwiththevlin/project_euler/registry"
)

func init() {
	registry.Register(5, func(input string) (string, error) {
		target, err := strconv.Atoi(input)
		if err != nil {
			return "", fmt.Errorf("invalid limit: %w", err)
		}
		return strconv.FormatUint(Solve(target), 10), nil
	})
}
