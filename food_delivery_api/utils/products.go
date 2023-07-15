package utils

import (
	"fmt"
	"strconv"
)

func ParseCategoryID(categoryIDStr string) (int64, error) {
	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid category ID: %s", categoryIDStr)
	}

	return categoryID, nil
}
