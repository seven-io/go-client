package sms77api

import (
	"strconv"
)

func toUint(id string, bitSize int) uint64 {
	n, err := strconv.ParseUint(id, 10, bitSize)

	if nil == err {
		return n
	}

	return 0
}
