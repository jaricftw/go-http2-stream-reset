package stream

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const http2Port = 7777

func TestHTTP2StreamReset(t *testing.T) {
	startHTTP2Server()
	proxyPort := startHTTP2ReverseProxy(http2Port)

	client := &http.Client{
		Timeout: time.Second,
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(netw, addr)
			},
		},
	}

	tests := []struct {
		msg  string
		port int
	}{
		{
			"direct",
			http2Port,
		},
		{
			"via proxy",
			proxyPort,
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%d", tt.port), nil)
			require.NoError(t, err)

			resp, err := client.Do(req)
			t.Logf("got err: %v", err)
			t.Logf("got resp: %v", resp)
			if err == nil {
				t.Logf("got reseponse body: %s", getResponseBody(t, resp))
			}
		})
	}
}

func startHTTP2Server() {
	h1Handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world returned by handler")
		panic(http.ErrAbortHandler)
	})
	h2Server := &http2.Server{}
	h2Handler := h2c.NewHandler(h1Handler, h2Server)

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", http2Port))
	panicOnError(err)

	server := &http.Server{
		Addr:    ln.Addr().String(),
		Handler: h2Handler,
	}

	go server.Serve(ln)
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

	ln, err := net.Listen("tcp", ":0")
	panicOnError(err)

	proxyServer := &http.Server{
		Addr:    ln.Addr().String(),
		Handler: proxy,
	}

	go proxyServer.Serve(ln)

	return ln.Addr().(*net.TCPAddr).Port
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func getResponseBody(t *testing.T, resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	return string(b)
}
