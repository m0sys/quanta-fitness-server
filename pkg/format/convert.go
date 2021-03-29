package format

import (
	"strconv"
)

func ConvertToBase64(id string) (int64, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return -1, err
	}

	return n, nil
}
