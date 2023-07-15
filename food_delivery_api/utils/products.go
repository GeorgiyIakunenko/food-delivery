package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseCategoryIDs(categoryIDsStr string) ([]int64, error) {
	var categoryIDs []int64
	if categoryIDsStr == "" {
		return categoryIDs, nil
	}

	idsStr := strings.Split(categoryIDsStr, ",")
	for _, idStr := range idsStr {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid category ID: %s", idStr)
		}
		categoryIDs = append(categoryIDs, id)
	}

	return categoryIDs, nil
}
