package client

import (
	"sync"
	"sync/atomic"
)

// changer is an interface for representing round-robin balancing API key usage.
type changer interface {
	Next() string
}

type roundRobin struct {
	apiKeys []*string
	next    uint32
	mu      sync.RWMutex
}

// newRoundRobin returns a new instance of roundRobin.
func newRoundRobin(keys ...string) changer {
	apiKeys := make([]*string, len(keys))
	for i, key := range keys {
		newKey := key
		apiKeys[i] = &newKey
	}
	return &roundRobin{
		apiKeys: apiKeys,
	}
}

func (r *roundRobin) Next() string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	n := atomic.AddUint32(&r.next, 1)
	return *r.apiKeys[(int(n)-1)%len(r.apiKeys)]
}
