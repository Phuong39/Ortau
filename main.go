package main

import (
	. "Ortau/reverseproxy"
	"Ortau/static"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type ReverseProxy struct {
	RedirectUrl string
}

//修改

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(p.RedirectUrl)
	if err != nil {
		panic(err)
	}
	proxy := NewProxy(remote)
	proxy.ServeHTTP(w, r)
}

func main() {
	fmt.Printf("\033[1;31;40m%s\033[0m",static.Banner)
	fmt.Println("aaaaaa")

	localIpAddress := "127.0.0.1:5056"
	proxyHandle := &ReverseProxy{RedirectUrl: "http://www.proxygateway.com"}
	log.Printf("proxy addr: %v, RedirectUrl: %v\n", localIpAddress, proxyHandle)

	err := http.ListenAndServe(localIpAddress, proxyHandle)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}

}