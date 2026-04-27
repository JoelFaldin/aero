package balancer

import "sync/atomic"

type Balancer struct {
	Upstreams []string
	Current   atomic.Uint32
}

func (b *Balancer) Next() string {
	n := b.Current.Add(1)

	return b.Upstreams[(int(n)-1)%len(b.Upstreams)]
}
