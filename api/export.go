package api

import (
	"sync/atomic"
)

// SwapRequestCount is
func SwapRequestCount() float64 {

	return float64(atomic.SwapUint32(&requestCount, 0))
}
