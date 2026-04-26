package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler() {
	target, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	fmt.Println("Start proxy server on :3000")
	http.ListenAndServe(":3000", proxy)
}
