package cache

import (
	"strconv"
)

type Result struct {
	Val       string
}

func (r Result) Int() (int, error) {
	intRes, err := strconv.Atoi(r.Val)
	if err != nil {
		return 0, err
	}
	return intRes, nil
}
