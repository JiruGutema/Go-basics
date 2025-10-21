package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)


func main() {

	servers := []string{
		"http://localhost:3001",
		"http://localhost:3002",
		"http://localhost:3003",
	}

	count := 0

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			targetURL, _ := url.Parse(servers[count%len(servers)])
			count++
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = targetURL.Path + req.URL.Path

			// Add or modify query params
			q := req.URL.Query()
			q.Set("proxy", "true")
			req.URL.RawQuery = q.Encode()

			// Optional: modify headers
			req.Header.Set("X-Proxy-By", "GoReverseProxy")
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Proxying:", r.URL.String())
		proxy.ServeHTTP(w, r)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Load balancer running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
