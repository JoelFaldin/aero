package balancer

import "sync/atomic"

type Upstream struct {
	Url    string
	Active atomic.Bool
}

type Balancer struct {
	Upstreams []Upstream
	Current   atomic.Uint32
}

func (b *Balancer) Next() string {
	n := b.Current.Add(1)

	return b.Upstreams[(int(n)-1)%len(b.Upstreams)].Url
}
