package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

func NewProxy(target *url.URL) *httputil.ReverseProxy{
	proxy := httputil.NewSingleHostReverseProxy(target)
	return proxy
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, url *url.URL,endpoint string) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter , r *http.Request){

		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.Header.Set("X-Forwarded-Host",r.Header.Get("Host"))
		r.Host = url.Host

		path := r.URL.Path
		r.URL.Path = strings.TrimLeft(path,endpoint)
		fmt.Printf("[TinyRP] Redirecting rewquest to %s\n",r.URL,time.Now())
		proxy.ServeHTTP(w,r)
	}
}