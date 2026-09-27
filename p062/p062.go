package p062

import (
	"slices"
	"strconv"
)

func Solve(targetCount int) uint64 {
	type cubeGroup struct {
		cubeCount int
		smallest  uint64
	}

	groups := make(map[string]*cubeGroup)
	currentDigits := 0
	candidates := make(map[*cubeGroup]struct{})

	for n := uint64(1); ; n++ {
		cube := n * n * n

		sorted := sortDigits(cube)

		if len(sorted) > currentDigits {
			if len(candidates) > 0 {
				// return the smallest candidate that still has exactly targetCount
				var smallest uint64
				var foundAny bool

				for c := range candidates {
					if !foundAny || c.smallest < smallest {
						smallest = c.smallest
						foundAny = true
					}
				}
				return smallest
			}

			currentDigits = len(sorted)
		}

		if group, ok := groups[sorted]; ok {
			group.cubeCount++
			if group.cubeCount == targetCount {
				candidates[group] = struct{}{}
			}
			if group.cubeCount == targetCount+1 {
				delete(candidates, group)
			}
		} else {
			groups[sorted] = &cubeGroup{smallest: cube, cubeCount: 1}
		}

	}
}

func sortDigits(number uint64) string {
	s := []rune(strconv.FormatUint(number, 10))
	slices.Sort(s)
	return string(s)
}
