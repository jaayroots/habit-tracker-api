package utils

import "strconv"

func StrToUint(v string) (uint, error) {
	itemIDint, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(itemIDint), nil

}
