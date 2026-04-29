package app

import (
	"aero/internal/balancer"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler() {
	b := &balancer.Balancer{
		Upstreams: []balancer.Upstream{
			{Url: "http://localhost:8080"},
			{Url: "http://localhost:8081"},
			{Url: "http://localhost:8082"},
		},
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

	b.Ping(&b.Upstreams, 1)

	fmt.Println("Start proxy server on :3000")
	http.ListenAndServe(":3000", proxy)
}
