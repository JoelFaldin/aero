package balancer

import (
	"net/http"
	"sync/atomic"
	"time"
)

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
	if b.Upstreams[(int(n)-1)%len(b.Upstreams)].Active.Load() == false {
		n = b.Current.Add(1)
	}

	return b.Upstreams[(int(n)-1)%len(b.Upstreams)].Url
}

func (b *Balancer) Ping(upstream *[]Upstream, interval time.Duration) {
	for i := range b.Upstreams {
		go func(up *Upstream) {
			tick := time.NewTicker(interval * time.Second)

			for range tick.C {
				client := &http.Client{Timeout: 10 * time.Second}

				res, err := client.Get(up.Url + "/health")
				if err != nil {
					up.Active.Store(false)
					// fmt.Println("server", up.Url, "isnt active")
					continue
				}

				if res.Status != "200 OK" {
					up.Active.Store(false)
					// fmt.Println("server", up.Url, "isnt active")
					continue
				}

				res.Body.Close()

				up.Active.Store(true)
				// fmt.Println("server", up.Url, "is active")
			}
		}(&b.Upstreams[i])
	}
}
