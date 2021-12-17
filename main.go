package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	reverseproxy "reverseproxy/reverseproxy"
)

type ReverseProxy struct {
	UpstreamUrl string
}

//修改


func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(p.UpstreamUrl)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func main() {
	addr := ":5050"
	proxyHandle := &ReverseProxy{UpstreamUrl: "https://www.jd.com"}
	log.Printf("proxy addr: %v, Upstream: %v\n", addr, proxyHandle)

	err := http.ListenAndServe(addr, proxyHandle)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}