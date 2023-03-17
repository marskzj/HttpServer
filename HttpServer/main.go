package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// 使用 flag 包解析命令行参数
	port := flag.String("P", "8080", "Port to run the server on")
	username := flag.String("u", "your_username", "Username for authentication")
	password := flag.String("p", "your_password", "Password for authentication")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || user != *username || pass != *password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		// 获取客户端IP、请求方法和访问的地址
		clientIP := r.RemoteAddr
		requestMethod := r.Method
		requestURL := r.URL.String()

		// 打印客户端IP、操作和访问的地址
		log.Printf("Client IP: %s, Operation: %s, URL: %s", clientIP, requestMethod, requestURL)

		fs := http.FileServer(http.Dir("."))
		fs.ServeHTTP(w, r)
	})

	log.Printf("Starting server on port %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
