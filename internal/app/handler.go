package app

import (
	"aero/internal/balancer"
	"aero/internal/config"
	"aero/internal/logger"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func Handler(cf config.Config, verbose bool) {
	urls := make([]string, len(cf.Upstreams))

	for i, u := range cf.Upstreams {
		urls[i] = u.Url
	}

	b := balancer.NewBalancer(urls)

	proxy := &httputil.ReverseProxy{}
	proxy.Rewrite = func(pr *httputil.ProxyRequest) {
		upstream := b.Next()
		tr, err := url.Parse(upstream)
		if verbose && err != nil {
			logger.ErrorLogger(err)
		}

		logger.Logger(fmt.Sprintf("server: %s", tr), verbose)

		pr.SetURL(tr)
		pr.SetXForwarded()
	}

	b.Ping(b.Upstreams, time.Duration(cf.Proxy.HealthCheckInterval), verbose)

	fmt.Printf("Start proxy server on :%s\n", cf.Proxy.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cf.Proxy.Port), proxy)
}
