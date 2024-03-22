package sort

import (
	"fmt"
	"strings"
)

type Sort struct {
	Field     string
	Direction Direction
}

func Parse(expr string) ([]Sort, error) {
	texpr := strings.TrimSpace(expr)
	splits := strings.Split(texpr, " ")
	if len(splits) <= 1 {
		return nil, fmt.Errorf("invalid sort expression")
	}

	if len(splits)%2 != 0 {
		return nil, fmt.Errorf("a sort expr must be in pair")
	}

	var sorts []Sort
	for i := 0; i < len(splits); i += 2 {
		dir := Direction(splits[i+1])
		if !dir.Valid() {
			return nil, fmt.Errorf("invalid sort direction: %s", splits[i+1])
		}
		sorts = append(sorts, Sort{
			Field:     splits[i],
			Direction: dir,
		})
	}

	return sorts, nil
}
