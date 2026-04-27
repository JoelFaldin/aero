package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Balancer struct {
	upstreams []string
	current   atomic.Uint32
}

func (b *Balancer) Next() string {
	var wg sync.WaitGroup

	wg.Add(1)
	n := b.current.Add(1)
	wg.Done()

	return b.upstreams[(int(n)-1)%len(b.upstreams)]
}

func Handler() {
	b := &Balancer{
		upstreams: []string{"http://localhost:8080", "http://localhost:8081", "http://localhost:8082"},
	}

	proxy := &httputil.ReverseProxy{}
	proxy.Rewrite = func(pr *httputil.ProxyRequest) {
		upstream := b.Next()
		tr, err := url.Parse(upstream)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("server:", tr)

		pr.SetURL(tr)
		pr.SetXForwarded()
	}

	fmt.Println("Start proxy server on :3000")
	http.ListenAndServe(":3000", proxy)
}
