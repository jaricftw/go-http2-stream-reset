package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"golang.org/x/net/http2"

	"github.com/carl-mastrangelo/h2c"
)

func main() {
	go startBackendServer()

	//go startProxy()
	startHTTP2ReverseProxy(7777)
	time.Sleep(500 * time.Millisecond)

	resp, err := http.Get("http://localhost:9900/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
func startHTTP2ReverseProxy(p int) int {
	fmt.Printf("starting reverse proxy with single backend port: %d\n", p)
	rpURL, err := url.Parse(fmt.Sprintf("http://localhost:%d", p))
	panicOnError(err)
	proxy := httputil.NewSingleHostReverseProxy(rpURL)
	proxy.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(netw, addr string, _ *tls.Config) (net.Conn, error) {
			ta, err := net.ResolveTCPAddr(netw, addr)
			if err != nil {
				return nil, err
			}
			return net.DialTCP(netw, nil, ta)
		},
		DisableCompression: true,
	}

	ln, err := net.Listen("tcp", ":9900")
	panicOnError(err)

	proxyServer := &http.Server{
		Handler: proxy,
	}

	go proxyServer.Serve(ln)

	return ln.Addr().(*net.TCPAddr).Port
}

func startBackendServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello, World " + r.Proto)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	s := &http.Server{
		Addr:    ":7777",
		Handler: http.HandlerFunc(handler),
	}
	h2c.AttachClearTextHandler(nil, s)
	log.Fatal(s.ListenAndServe())
}

func startProxy() {
	rp := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "localhost:7777"
		},
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				ta, err := net.ResolveTCPAddr(network, addr)
				if err != nil {
					return nil, err
				}
				return net.DialTCP(network, nil, ta)
			},
		},
	}

	log.Fatal(http.ListenAndServe(":9900", rp))
}
