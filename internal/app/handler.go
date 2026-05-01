package app

import (
	"aero/internal/balancer"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(verbose bool) {
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
		if verbose && err != nil {
			log.Fatal(err)
		}

		if verbose {
			fmt.Println("server:", tr)
		}

		pr.SetURL(tr)
		pr.SetXForwarded()
	}

	b.Ping(&b.Upstreams, 10, verbose)

	fmt.Println("Start proxy server on :3000")
	http.ListenAndServe(":3000", proxy)
}
