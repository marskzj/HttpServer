package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := flag.String("p", "8080", "Port to run the server on")
	flag.Parse()

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 读取请求正文中的截图数据
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// 获取发送者IP并将其用作文件名
		ip := strings.Split(r.RemoteAddr, ":")[0]
		fileName := fmt.Sprintf("%s.jpg", ip)

		// 将截图数据保存为文件
		err = ioutil.WriteFile(fileName, data, 0644)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Screenshot saved as %s", fileName)
	})

	log.Printf("Starting server on port %s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
