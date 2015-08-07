package cc

import (
	"math"
	"os"
	"sync/atomic"
)

var (
	MinInt32 = math.MinInt32
)

func GetRoot(s string) string {
	root := os.Getenv("TEST_ROOT")
	if root == "" {
		root = "/data/muxque"
	}
	return root + s
}

type State uint32

func (s *State) IsClosed() bool {
	return atomic.LoadUint32((*uint32)(s)) == uint32(CloseState)
}

func (s *State) ToClose() bool {
	return atomic.CompareAndSwapUint32((*uint32)(s), uint32(InitState), uint32(CloseState))
}

const (
	InitState State = iota
	CloseState
)
