package stream

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

const http1Port = 6666

func TestHTTP1(t *testing.T) {
	startHTTPServer()
	proxyPort := startHTTP1ReverseProxy(http1Port)

	tests := []struct {
		msg  string
		port int
	}{
		{
			"direct",
			http1Port,
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

			resp, err := http.DefaultClient.Do(req)
			t.Logf("got err: %v", err)
			t.Logf("got resp: %v", resp)
			if err == nil {
				t.Logf("got reseponse body: %s", getResponseBody(t, resp))
			}
		})
	}
}

func startHTTPServer() {
	h1Handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world returned by handler")
		//panic(http.ErrAbortHandler)
	})

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", http1Port))
	panicOnError(err)

	server := &http.Server{
		Handler: h1Handler,
	}

	go server.Serve(ln)
}

func startHTTP1ReverseProxy(p int) int {
	fmt.Printf("starting reverse proxy with single backend port: %d\n", p)
	rpURL, err := url.Parse(fmt.Sprintf("http://localhost:%d", p))
	panicOnError(err)
	proxy := httputil.NewSingleHostReverseProxy(rpURL)
	proxy.Transport = &http.Transport{}

	ln, err := net.Listen("tcp", ":0")
	panicOnError(err)

	proxyServer := &http.Server{
		Handler: proxy,
	}

	go proxyServer.Serve(ln)

	return ln.Addr().(*net.TCPAddr).Port
}
