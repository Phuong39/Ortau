package main

import (
	"Ortau/conf"
	. "Ortau/reverseproxy"
	"fmt"
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
	//fmt.Printf("\033[1;31;40m%s\033[0m",static.Banner)
	//
	//localIpAddress := ":5056"
	//proxyHandle := &ReverseProxy{RedirectUrl: "http://www.jd.com"}
	//log.Printf("proxy addr: %v, RedirectUrl: %v\n", localIpAddress, proxyHandle)
	//
	//err := http.ListenAndServe(localIpAddress, proxyHandle)
	//if err != nil {
	//	log.Fatalln("ListenAndServe: ", err)
	//}

	a:=conf.GetCfgSectionKey("one","path")
	fmt.Println(a)

}