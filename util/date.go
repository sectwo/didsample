package util

import (
	"fmt"
	"time"
)

func GetCurrentDate() string {

	now := time.Now()
	nanos := now.UnixNano()

	return fmt.Sprint(nanos)
}

func GetExpireDate() string {
	now := time.Now()
	nanos := now.UnixNano()
	nanos = nanos + 1000000

	return fmt.Sprint(nanos)
}
