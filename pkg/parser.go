package pkg

import (
	"sort"
	"strings"
)

func ParseCategory(category string) ([]string, int) {
	parsedCategories := strings.Split(category, ",")
	sort.Strings(parsedCategories)
	return parsedCategories, len(parsedCategories)
}
