package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 使用 flag 包解析命令行参数
	downloadURL := flag.String("d", "", "URL to download the file from")
	username := flag.String("u", "", "Username for authentication")
	password := flag.String("p", "", "Password for authentication")
	outputPath := flag.String("o", "", "Output path and file name")
	flag.Parse()

	if *downloadURL == "" {
		fmt.Println("Please provide a URL to download the file from using the -d flag.")
		return
	}

	// 创建带有基本身份验证的HTTP请求
	req, err := http.NewRequest("GET", *downloadURL, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	req.SetBasicAuth(*username, *password)

	// 发出HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// 创建输出文件
	outFile, err := os.Create(*outputPath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outFile.Close()

	// 将响应内容写入文件
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Printf("Error writing to output file: %v\n", err)
		return
	}

	fmt.Printf("File downloaded successfully: %s\n", *outputPath)
}
